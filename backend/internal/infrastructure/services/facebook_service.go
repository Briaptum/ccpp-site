package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"ccpp-backend/internal/domain/models"
)

type FacebookService struct {
	accessToken string
	baseURL     string
}

func NewFacebookService(accessToken string) *FacebookService {
	return &FacebookService{
		accessToken: accessToken,
		baseURL:     "https://graph.facebook.com/v18.0",
	}
}

// GetPagePhotos fetches photos from a Facebook page using Feed endpoint (requires fewer permissions)
func (s *FacebookService) GetPagePhotos(pageID string, limit int) ([]models.GalleryImage, error) {
	if limit <= 0 {
		limit = 10
	}

	// Try to fetch photos using the feed endpoint with attachments
	// This requires fewer permissions than the photos endpoint
	url := fmt.Sprintf("%s/%s/feed?fields=attachments{media,title,description},created_time&limit=%d&access_token=%s",
		s.baseURL, pageID, limit*2, s.accessToken) // Request more to filter for photos only

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch photos: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("facebook API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse the feed response
	var feedResponse struct {
		Data []struct {
			Attachments struct {
				Data []struct {
					Media struct {
						Image struct {
							Height int    `json:"height"`
							Width  int    `json:"width"`
							Src    string `json:"src"`
						} `json:"image"`
					} `json:"media"`
					Title       string `json:"title"`
					Description string `json:"description"`
				} `json:"data"`
			} `json:"attachments"`
			CreatedTime string `json:"created_time"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&feedResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert feed posts with images to gallery images
	galleryImages := make([]models.GalleryImage, 0, limit)
	for _, post := range feedResponse.Data {
		if len(post.Attachments.Data) == 0 {
			continue
		}

		for _, attachment := range post.Attachments.Data {
			if attachment.Media.Image.Src == "" {
				continue
			}

			title := attachment.Title
			if title == "" {
				title = attachment.Description
			}
			if title == "" {
				title = "Church Photo"
			}

			galleryImages = append(galleryImages, models.GalleryImage{
				ID:        "",
				Source:    attachment.Media.Image.Src,
				Thumbnail: attachment.Media.Image.Src,
				Title:     title,
				Alt:       title,
			})

			if len(galleryImages) >= limit {
				break
			}
		}

		if len(galleryImages) >= limit {
			break
		}
	}

	return galleryImages, nil
}

// GetMyPhotos fetches photos from "me" (the page/user the token belongs to)
func (s *FacebookService) GetMyPhotos(limit int) ([]models.GalleryImage, error) {
	if limit <= 0 {
		limit = 10
	}

	// Fetch photos using "me" endpoint
	url := fmt.Sprintf("%s/me/photos?type=uploaded&fields=id,images,name,created_time&limit=%d&access_token=%s",
		s.baseURL, limit, s.accessToken)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch photos: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("facebook API error (status %d): %s", resp.StatusCode, string(body))
	}

	var fbResponse models.FacebookPhotosResponse
	if err := json.NewDecoder(resp.Body).Decode(&fbResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert Facebook photos to gallery images
	galleryImages := make([]models.GalleryImage, 0, len(fbResponse.Data))
	for _, photo := range fbResponse.Data {
		if len(photo.Images) == 0 {
			continue
		}

		// Get the highest quality image
		source := photo.Images[0].Source
		thumbnail := source

		// Try to find a good thumbnail size
		for _, img := range photo.Images {
			if img.Width >= 600 && img.Width <= 1000 {
				thumbnail = img.Source
				break
			}
		}

		// Use smaller image for thumbnail if available
		if len(photo.Images) > 2 {
			thumbnail = photo.Images[len(photo.Images)-2].Source
		}

		title := photo.Name
		if title == "" {
			title = fmt.Sprintf("Photo from %s", photo.CreatedAt.Format("Jan 2, 2006"))
		}

		galleryImages = append(galleryImages, models.GalleryImage{
			ID:        photo.ID,
			Source:    source,
			Thumbnail: thumbnail,
			Title:     title,
			Alt:       title,
			CreatedAt: photo.CreatedAt,
		})
	}

	return galleryImages, nil
}
