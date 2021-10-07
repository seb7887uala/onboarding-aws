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
    - `insert-contact-ss`
- **API Gateway**