package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/processor"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/repository"
)

var (
	contactProcessor processor.Processor
	h                handler.Handler
)

func init() {
	r := repository.New()
	contactProcessor = processor.New(r)
	h = handler.New(contactProcessor)
}

func main() {
	lambda.Start(h.ProcessStatus)
}
