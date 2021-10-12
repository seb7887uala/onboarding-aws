package publisher

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

const (
	Region = "us-east-1"
)

var TopicArn string = os.Getenv("TOPIC_ARN")

type SNSPublisher interface {
	Publish(msg, id string) error
}

type snsPublisher struct {
	TopicArn string
	svc      *sns.SNS
}

func New() SNSPublisher {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(Region)}))
	svc := sns.New(sess)
	return &snsPublisher{
		TopicArn,
		svc,
	}
}

func (p *snsPublisher) Publish(msg, id string) error {
	_, err := p.svc.Publish(&sns.PublishInput{
		Message:  &msg,
		TopicArn: aws.String(p.TopicArn),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"id": {
				DataType:    aws.String("String"),
				StringValue: aws.String(id),
			},
		},
	})

	return err
}
