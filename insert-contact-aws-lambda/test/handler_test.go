package test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/handler"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/models"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/pkg/utils/apigw"
	"github.com/sebsegura/onboarding-aws/insert-contact-aws-lambda/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsertContact(t *testing.T) {
	var (
		ctrl           = gomock.NewController(t)
		mockRepository = mocks.NewMockContactRepository(ctrl)
		h              = handler.New(mockRepository)
		ctx            = context.Background()
	)
	defer ctrl.Finish()

	var (
		req = apigw.Request{
			Body: `{
				"first_name": "John",
				"last_name": "Doe"
			}`,
		}
		badReq = apigw.Request{
			Body: `{
				"first_name": "",
				"last_name": "Doe"
			}`,
		}
		res = models.Contact{
			ID:        "8be884d9-e2d1-4a6d-a8eb-b635997f92b6",
			FirstName: "John",
			LastName:  "Doe",
			Status:    "CREATED",
		}
	)

	t.Run("inserts a new contact", func(t *testing.T) {
		mockRepository.EXPECT().Insert("John", "Doe").Return(res, nil)

		response, err := h.Insert(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusOK, response.StatusCode)

		var body models.Contact
		json.Unmarshal([]byte(response.Body), &body)

		assert.Equal(t, res, body)
	})

	t.Run("handles bad request error", func(t *testing.T) {
		response, err := h.Insert(ctx, badReq)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		assert.Equal(t, "Request validation error. Missing required fields", response.Body)
	})

	t.Run("handles internal server error", func(t *testing.T) {
		mockRepository.EXPECT().Insert("John", "Doe").Return(models.Contact{}, fmt.Errorf("An error happened"))

		response, err := h.Insert(ctx, req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	})
}
