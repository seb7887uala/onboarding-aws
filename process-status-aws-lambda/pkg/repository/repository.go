package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/models"
)

const (
	TableName = "Contacts_SS"
	Region    = "us-east-1"
)

type ContactRepository interface {
	UpdateContact(contact models.Contact) error
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

func (r *repository) checkIfExists(id string) error {
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

	if err != nil || res.Item == nil {
		return fmt.Errorf("Contact %s not found", id)
	}

	return nil
}

func (r *repository) UpdateContact(contact models.Contact) error {
	if err := r.checkIfExists(contact.ID); err != nil {
		return err
	}

	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.TableName),
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
