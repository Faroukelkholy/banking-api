package customer

import (
	"errors"
	"testing"

	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	id = "123e4567-e89b-12d3-a456-426614174000"
)

func TestCCA_Success(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*models.Account")).Return(nil)

	srv := New(rMock)
	err := srv.CreateCustomerAccount(id, &models.Account{})

	assert.Nil(t, err)
}

func TestCCA_Error(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*models.Account")).Return(errors.New("no account table"))

	srv := New(rMock)
	err := srv.CreateCustomerAccount(id, &models.Account{})

	assert.NotNil(t, err)
}
