package service

import (
	"context"
	"fmt"

	"github.com/samandar2605/medium_notification_service/config"
	pb "github.com/samandar2605/medium_notification_service/genproto/notification_service"
	emailPkg "github.com/samandar2605/medium_notification_service/pkg/email"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NotificationService struct {
	pb.UnimplementedNotificationServiceServer
	cfg *config.Config
}

func NewNotificationService(cfg *config.Config) *NotificationService {
	return &NotificationService{
		cfg: cfg,
	}
}

func (s *NotificationService) SendEmail(ctx context.Context, req *pb.SendEmailRequest) (*pb.Empty, error) {
	fmt.Println(req)
	err := emailPkg.SendEmail(s.cfg, &emailPkg.SendEmailRequest{
		To:      []string{req.To},
		Subject: req.Subject,
		Body:    req.Body,
		Type:    req.Type,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return &pb.Empty{}, nil
}
