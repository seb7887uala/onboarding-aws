package processor

import (
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/repository"
)

type Processor interface {
	UpdateStatus(contact models.Contact) error
}

type processor struct {
	r repository.ContactRepository
}

func New(r repository.ContactRepository) Processor {
	return &processor{
		r,
	}
}

func (p *processor) UpdateStatus(contact models.Contact) error {
	// Change contact status to 'PROCESSED'
	contact.SetProcessed()

	if err := p.r.UpdateContact(contact); err != nil {
		return err
	}

	return nil
}
