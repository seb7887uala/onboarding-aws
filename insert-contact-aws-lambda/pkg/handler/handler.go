package handler

import (
	"context"
	"encoding/json"

	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/repository"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/utils/apigw"
)

type Handler interface {
	Insert(ctx context.Context, req apigw.Request) (apigw.Response, error)
}

type handler struct {
	r repository.ContactRepository
}

func New(r repository.ContactRepository) Handler {
	return &handler{
		r,
	}
}

func (h *handler) Insert(ctx context.Context, req apigw.Request) (apigw.Response, error) {
	var insertReq models.InsertRequest
	json.Unmarshal([]byte(req.Body), &insertReq)

	if insertReq.FirstName == "" || insertReq.LastName == "" {
		return apigw.BadRequestResponse("Request validation error. Missing required fields"), nil
	}

	item, err := h.r.Insert(insertReq.FirstName, insertReq.LastName)
	if err != nil {
		return apigw.InternalErrorResponse(), nil
	}

	contact, err := json.Marshal(item)
	if err != nil {
		return apigw.InternalErrorResponse(), nil
	}

	return apigw.OkResponse(string(contact)), nil
}
