package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/logger"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/publisher"
	"go.uber.org/zap"
)

type Handler interface {
	PublishContact(ctx context.Context, e events.DynamoDBEvent)
}

type handler struct {
	sns publisher.SNSPublisher
}

func New(sns publisher.SNSPublisher) Handler {
	return &handler{
		sns,
	}
}

func (h *handler) PublishContact(ctx context.Context, e events.DynamoDBEvent) {
	log := logger.Setup()

	for _, record := range e.Records {
		log.Info("Processing DynamoDB event",
			zap.String("EventID", record.EventID),
			zap.String("EventName", record.EventName),
			zap.String("EventSource", record.EventSource),
		)
		// Only publish when a new contact has been created
		if record.EventName == "INSERT" {
			msg := models.Contact{
				ID:        record.Change.NewImage["id"].String(),
				FirstName: record.Change.NewImage["firstName"].String(),
				LastName:  record.Change.NewImage["lastName"].String(),
				Status:    record.Change.NewImage["status"].String(),
			}
			// Publish message to SNS topic
			if err := h.sns.Publish(msg.String(), msg.ID); err != nil {
				log.Error("Error publishing contact to SNS topic", zap.String("ContactID", msg.ID))
			} else {
				log.Info("Contact published to SNS topic", zap.String("ContactID", msg.ID))
				return
			}
		}
	}
}
