package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/publisher"
)

func main() {
	var (
		snsPublisher = publisher.New()
		h            = handler.New(snsPublisher)
	)

	lambda.Start(h.PublishContact)
}
