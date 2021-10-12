package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/process-contact-aws-lambda/pkg/publisher"
)

var h handler.Handler

func init() {
	snsPublisher := publisher.New()
	h = handler.New(snsPublisher)
}

func main() {
	lambda.Start(h.PublishContact)
}
