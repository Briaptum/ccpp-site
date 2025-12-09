package services

import (
	"fmt"
	"log"
	"strings"

	"ccpp-backend/internal/domain/models"
	"ccpp-backend/internal/domain/repositories"
	"ccpp-backend/internal/domain/services"
)

type contactRequestServiceImpl struct {
	repo         repositories.ContactRequestRepository
	emailService *EmailService
}

func NewContactRequestService(repo repositories.ContactRequestRepository, emailService *EmailService) services.ContactRequestService {
	return &contactRequestServiceImpl{
		repo:         repo,
		emailService: emailService,
	}
}

func (s *contactRequestServiceImpl) CreateContactRequest(request *models.ContactRequest) error {
	if request.Metadata == nil {
		request.Metadata = models.JSONB{}
	}

	if err := s.repo.Create(request); err != nil {
		return err
	}

	// Fire-and-forget email notification
	go func(r models.ContactRequest) {
		if s.emailService == nil || !s.emailService.IsConfigured() {
			return
		}

		if err := s.sendNotificationEmail(r); err != nil {
			log.Printf("Failed to send contact request notification email: %v", err)
		}
	}(*request)

	return nil
}

func (s *contactRequestServiceImpl) GetContactRequests() ([]models.ContactRequest, error) {
	return s.repo.FindAll()
}

func (s *contactRequestServiceImpl) GetContactRequest(id uint) (*models.ContactRequest, error) {
	return s.repo.FindByID(id)
}

func (s *contactRequestServiceImpl) DeleteContactRequest(id uint) error {
	return s.repo.Delete(id)
}

func (s *contactRequestServiceImpl) sendNotificationEmail(request models.ContactRequest) error {
	subject := fmt.Sprintf("New Contact Form Submission from %s", request.Name)

	messageHTML := strings.ReplaceAll(request.Message, "&", "&amp;")
	messageHTML = strings.ReplaceAll(messageHTML, "<", "&lt;")
	messageHTML = strings.ReplaceAll(messageHTML, ">", "&gt;")
	messageHTML = strings.ReplaceAll(messageHTML, "\n", "<br>")

	phoneHTML := ""
	if request.Phone != nil && *request.Phone != "" {
		phoneHTML = fmt.Sprintf(`<p style="margin: 10px 0;"><strong>Phone:</strong> <a href="tel:%s" style="color: #00d3f3; text-decoration: none;">%s</a></p>`, *request.Phone, *request.Phone)
	}

	body := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333; margin: 0; padding: 20px; background-color: #f9f9f9;">
	<div style="max-width: 600px; margin: 0 auto; background-color: #ffffff; padding: 30px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1);">
		<h2 style="color: #00d3f3; margin-top: 0; margin-bottom: 20px; font-size: 24px;">New Contact Form Submission</h2>
		
		<div style="background-color: #f5f5f5; padding: 20px; border-radius: 5px; margin: 20px 0;">
			<p style="margin: 10px 0;"><strong>Name:</strong> %s</p>
			<p style="margin: 10px 0;"><strong>Reason:</strong> %s</p>
			<p style="margin: 10px 0;"><strong>Email:</strong> <a href="mailto:%s" style="color: #00d3f3; text-decoration: none;">%s</a></p>
			%s
			<p style="margin: 10px 0;"><strong>Message:</strong></p>
			<div style="background-color: white; padding: 15px; border-left: 3px solid #00d3f3; margin: 10px 0; border-radius: 3px;">
				%s
			</div>
		</div>
		
		<div style="margin-top: 20px; padding-top: 20px; border-top: 1px solid #ddd; font-size: 12px; color: #666;">
			<p style="margin: 5px 0;"><strong>Submitted:</strong> %s</p>
			%s
		</div>
	</div>
</body>
</html>`,
		request.Name,
		request.Reason,
		request.Email,
		request.Email,
		phoneHTML,
		messageHTML,
		request.CreatedAt.Format("January 2, 2006 at 3:04 PM MST"),
		s.formatMetadata(request),
	)

	return s.emailService.SendNotificationEmail(subject, body)
}

func (s *contactRequestServiceImpl) formatMetadata(request models.ContactRequest) string {
	var parts []string

	if request.IPAddress != nil && *request.IPAddress != "" {
		parts = append(parts, fmt.Sprintf("<p><strong>IP Address:</strong> %s</p>", *request.IPAddress))
	}

	if request.UserAgent != nil && *request.UserAgent != "" {
		parts = append(parts, fmt.Sprintf("<p><strong>User Agent:</strong> %s</p>", *request.UserAgent))
	}

	if len(parts) > 0 {
		return strings.Join(parts, "")
	}

	return ""
}
