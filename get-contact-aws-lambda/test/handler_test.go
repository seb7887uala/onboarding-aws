package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/pkg/utils/apigw"
	"github.com/sebsegura/onboarding-aws/get-contact-aws-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetContact(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		mockRepository = mocks.NewMockContactRepository(ctrl)
		h              = handler.New(mockRepository)
		ctx            = context.Background()
	)
	defer ctrl.Finish()

	var (
		id  = "8be884d9-e2d1-4a6d-a8eb-b635997f92b6"
		req = apigw.Request{
			PathParameters: map[string]string{
				"id": "8be884d9-e2d1-4a6d-a8eb-b635997f92b6",
			},
		}
		res = models.Contact{
			ID:        "8be884d9-e2d1-4a6d-a8eb-b635997f92b6",
			FirstName: "John",
			LastName:  "Doe",
			Status:    "CREATED",
		}
	)

	t.Run("gets contact information", func(t *testing.T) {
		mockRepository.EXPECT().GetContact(id).Return(res, nil)

		response, err := h.GetItem(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.StatusCode)

		var body models.Contact
		json.Unmarshal([]byte(response.Body), &body)

		assert.Equal(t, res, body)
	})

	t.Run("handles not found error", func(t *testing.T) {
		mockRepository.EXPECT().GetContact(id).Return(models.Contact{}, fmt.Errorf("Contact not found"))

		response, err := h.GetItem(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		assert.Equal(t, "Contact not found", response.Body)
	})
}
