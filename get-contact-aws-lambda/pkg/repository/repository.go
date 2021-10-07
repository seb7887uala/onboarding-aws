package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/models"
)

const (
	TableName = "Contacts_SS"
	Region    = "us-east-1"
)

type ContactRepository interface {
	GetContact(id string) (models.Contact, error)
}

type repository struct {
	TableName string
	db        *dynamodb.DynamoDB
}

func New() ContactRepository {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(Region)}))
	db := dynamodb.New(sess)

	return &repository{
		TableName,
		db,
	}
}

func (r *repository) GetContact(id string) (models.Contact, error) {
	res, err := r.db.GetItem(
		&dynamodb.GetItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"id": {
					S: aws.String(id),
				},
			},
			TableName: aws.String(r.TableName),
		},
	)
	if err != nil {
		return models.Contact{}, err
	}

	if res.Item == nil {
		return models.Contact{}, fmt.Errorf("Contact not found")
	}

	var contact models.Contact
	if err = dynamodbattribute.UnmarshalMap(res.Item, &contact); err != nil {
		return models.Contact{}, err
	}

	return contact, nil
}
