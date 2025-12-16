package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"ccpp-backend/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct{}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginResponse struct {
	Token string    `json:"token"`
	User  loginUser `json:"user"`
}

type loginUser struct {
	ID    string   `json:"id"`
	Email string   `json:"email"`
	Role  string   `json:"role"`
	Sites []string `json:"sites"`
}

type auth0TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type auth0User struct {
	UserID      string                 `json:"user_id"`
	Email       string                 `json:"email"`
	AppMetadata map[string]interface{} `json:"app_metadata"`
}

type auth0Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request format"})
		return
	}

	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	auth0ClientID := os.Getenv("AUTH0_CLIENT_ID")
	auth0ClientSecret := os.Getenv("AUTH0_CLIENT_SECRET")

	if auth0Domain == "" || auth0ClientID == "" || auth0ClientSecret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth0 configuration missing"})
		return
	}

	_, err := h.authenticateWithAuth0(auth0Domain, auth0ClientID, auth0ClientSecret, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("authentication failed: %v", err)})
		return
	}

	user, err := h.getUserFromAuth0(auth0Domain, auth0ClientID, auth0ClientSecret, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get user information: %v", err)})
		return
	}

	if !h.hasRequiredPermissions(user) {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied: insufficient permissions"})
		return
	}

	token, err := h.createJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		Token: token,
		User: loginUser{
			ID:    user.UserID,
			Email: user.Email,
			Role:  h.getUserRole(user),
			Sites: h.getUserSites(user),
		},
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	auth0Domain := os.Getenv("AUTH0_DOMAIN")
	clientID := os.Getenv("AUTH0_CLIENT_ID")

	if auth0Domain == "" || clientID == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth0 configuration missing"})
		return
	}

	logoutURL := url.URL{
		Scheme: "https",
		Host:   auth0Domain,
		Path:   "/v2/logout",
	}

	siteURL := os.Getenv("SITE_URL")
	if siteURL == "" {
		siteURL = "http://localhost:3000"
	}

	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("returnTo", siteURL)
	logoutURL.RawQuery = params.Encode()

	c.Redirect(http.StatusTemporaryRedirect, logoutURL.String())
}

func (h *AuthHandler) Profile(c *gin.Context) {
	claimsValue, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	auth0Claims, ok := claimsValue.(*middleware.Auth0Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user data"})
		return
	}

	var role string
	var sites []string
	if auth0Claims.AppMetadata != nil {
		if roleVal, ok := auth0Claims.AppMetadata["role"].(string); ok {
			role = roleVal
		}
		if sitesVal, ok := auth0Claims.AppMetadata["sites"].([]interface{}); ok {
			for _, site := range sitesVal {
				if siteStr, ok := site.(string); ok {
					sites = append(sites, siteStr)
				}
			}
		}
	}

	user := loginUser{
		ID:    auth0Claims.Sub,
		Email: auth0Claims.Email,
		Role:  role,
		Sites: sites,
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *AuthHandler) authenticateWithAuth0(domain, clientID, clientSecret, email, password string) (*auth0TokenResponse, error) {
	urlStr := fmt.Sprintf("https://%s/oauth/token", domain)

	payload := map[string]interface{}{
		"grant_type":    "password",
		"username":      email,
		"password":      password,
		"client_id":     clientID,
		"client_secret": clientSecret,
		"scope":         "openid profile email",
		"realm":         "Username-Password-Authentication",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var auth0Err auth0Error
		_ = json.NewDecoder(resp.Body).Decode(&auth0Err)
		return nil, fmt.Errorf("auth0 error: %s", auth0Err.ErrorDescription)
	}

	var tokenResp auth0TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, err
	}

	return &tokenResp, nil
}

func (h *AuthHandler) getUserFromAuth0(domain, clientID, clientSecret, email string) (*auth0User, error) {
	mgmtToken, err := h.getManagementToken(domain, clientID, clientSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to get management token: %w", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/api/v2/users-by-email?email=%s", domain, url.QueryEscape(email)), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+mgmtToken)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user from Auth0: %s", string(body))
	}

	var users []auth0User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return &users[0], nil
}

func (h *AuthHandler) getManagementToken(domain, clientID, clientSecret string) (string, error) {
	urlStr := fmt.Sprintf("https://%s/oauth/token", domain)

	payload := map[string]interface{}{
		"grant_type":    "client_credentials",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"audience":      fmt.Sprintf("https://%s/api/v2/", domain),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("failed to get management token: %s", string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("no access token in response")
	}

	return accessToken, nil
}

func (h *AuthHandler) hasRequiredPermissions(user *auth0User) bool {
	role := h.getUserRole(user)
	sites := h.getUserSites(user)
	siteID := os.Getenv("SITE_ID")
	if siteID == "" {
		siteID = "ccphnompenh-com"
	}

	if role == "admin" {
		return true
	}

	if siteID != "" {
		for _, site := range sites {
			if site == siteID {
				return true
			}
		}
	}

	return false
}

func (h *AuthHandler) getUserRole(user *auth0User) string {
	if user.AppMetadata == nil {
		return ""
	}
	if role, ok := user.AppMetadata["role"].(string); ok {
		return role
	}
	return ""
}

func (h *AuthHandler) getUserSites(user *auth0User) []string {
	if user.AppMetadata == nil {
		return []string{}
	}
	if sites, ok := user.AppMetadata["sites"].([]interface{}); ok {
		var siteStrings []string
		for _, site := range sites {
			if siteStr, ok := site.(string); ok {
				siteStrings = append(siteStrings, siteStr)
			}
		}
		return siteStrings
	}
	return []string{}
}

func (h *AuthHandler) createJWTToken(user *auth0User) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "your-secret-key"
	}

	claims := jwt.MapClaims{
		"sub":   user.UserID,
		"email": user.Email,
		"https://ccphnompenh-com/app_metadata": map[string]interface{}{
			"role":  h.getUserRole(user),
			"sites": h.getUserSites(user),
		},
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
