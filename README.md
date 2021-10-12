# Onboarding AWS

## Getting Started

### Requirements

- Go v1.16

## Contacts: data model

- `id`: String (Partition Key)
- `firstName`: String
- `lastName`: String
- `status`: String

This table has a **simple primary key** made up of just a partition key (`id`)

## AWS Resources

- **DynamoDB**
    - `Contacts_SS`
- **Lambdas**
    - `insert-contact-ss` (Activity 1)
    - `get-contact-ss` (Activity 2)
    - `process-contact-ss` (Activity 3)
    - `process-status-ss` (Activity 4)
- **API Gateway**
    - `api-contacts-ss`
- **SNS**
    - `new-contact-ss`