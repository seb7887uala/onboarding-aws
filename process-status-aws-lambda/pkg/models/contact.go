package models

const (
	PROCESSED = "PROCESSED"
)

type Contact struct {
	ID        string `dynamodbav:"id" json:"id"`
	FirstName string `dynamodbav:"firstName" json:"first_name"`
	LastName  string `dynamodbav:"lastName" json:"last_name"`
	Status    string `dynamodbav:"status" json:"status"`
}

func (c *Contact) SetProcessed() {
	c.Status = PROCESSED
}
