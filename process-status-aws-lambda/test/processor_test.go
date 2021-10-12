package processor_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/pkg/processor"
	"github.com/sebsegura/onboarding-aws/process-status-aws-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContactProcessor(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		mockRepository = mocks.NewMockContactRepository(ctrl)
		proc           = processor.New(mockRepository)
	)
	defer ctrl.Finish()

	var (
		contact = models.Contact{
			ID:        "8be884d9-e2d1-4a6d-a8eb-b635997f92b6",
			FirstName: "John",
			LastName:  "Doe",
			Status:    "PROCESSED",
		}
		mockError = fmt.Errorf("Contact does not exist")
	)

	t.Run("updates contact status", func(t *testing.T) {
		mockRepository.EXPECT().UpdateContact(contact).Return(nil)

		err := proc.UpdateStatus(contact)

		require.NoError(t, err)
	})

	t.Run("returns error if contact does not exist", func(t *testing.T) {
		mockRepository.EXPECT().UpdateContact(contact).Return(mockError)

		err := proc.UpdateStatus(contact)

		assert.Equal(t, mockError, err)
	})
}
