package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/repository"
)

var h handler.Handler

func init() {
	contactRepository := repository.New()
	h = handler.New(contactRepository)
}

func main() {
	lambda.Start(h.Insert)
}
