package handler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/logger"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/processor"
	"go.uber.org/zap"
)

type Handler interface {
	ProcessStatus(ctx context.Context, evt events.SNSEvent)
}

type handler struct {
	p processor.Processor
}

func New(p processor.Processor) Handler {
	return &handler{
		p,
	}
}

func (h *handler) ProcessStatus(ctx context.Context, evt events.SNSEvent) {
	log := logger.Setup()

	for _, record := range evt.Records {
		var contact models.Contact
		if err := json.Unmarshal([]byte(record.SNS.Message), &contact); err != nil {
			log.Error("Error unmarshalling SNS record")
		}

		// Update contact status
		if err := h.p.UpdateStatus(contact); err != nil {
			log.Error("Error processing contact",
				zap.String("ContactID", contact.ID),
				zap.String("ErrorMessage", err.Error()),
			)
		} else {
			log.Info("Contact processed",
				zap.String("ContactID", contact.ID),
			)
		}
	}
}
