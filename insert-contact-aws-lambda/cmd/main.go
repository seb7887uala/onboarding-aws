package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/repository"
)

func main() {
	var (
		contactRepository = repository.New()
		h                 = handler.New(contactRepository)
	)
	lambda.Start(h.Insert)
}
