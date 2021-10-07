package handler

import (
	"context"
	"encoding/json"

	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/logger"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/repository"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/utils/apigw"
	"go.uber.org/zap"
)

type Handler interface {
	GetItem(ctx context.Context, req apigw.Request) (apigw.Response, error)
}

type handler struct {
	r repository.ContactRepository
}

func New(r repository.ContactRepository) Handler {
	return &handler{
		r,
	}
}

func (h *handler) GetItem(ctx context.Context, req apigw.Request) (apigw.Response, error) {
	log := logger.Setup()
	id := req.PathParameters["id"]

	// Log request
	log.Info("Request",
		zap.String("ID", id),
	)

	item, err := h.r.GetContact(id)
	if err != nil {
		return apigw.NotFoundResponse(err.Error()), nil
	}

	contact, err := json.Marshal(item)
	if err != nil {
		return apigw.InternalErrResponse(), nil
	}

	return apigw.OkResponse(string(contact)), nil
}
