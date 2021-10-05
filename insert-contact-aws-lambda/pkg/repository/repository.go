package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/models"
)

const (
	TableName = "Contacts"
	Region    = "us-east-1"
)

type ContactRepository interface {
	Insert(firstName string, lastName string) (models.Contact, error)
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

func (r *repository) Insert(firstName string, lastName string) (models.Contact, error) {
	contact := models.Contact{
		ID:        uuid.New().String(),
		FirstName: firstName,
		LastName:  lastName,
		Status:    "CREATED",
	}

	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return models.Contact{}, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.TableName),
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return models.Contact{}, nil
	}

	return contact, nil
}
