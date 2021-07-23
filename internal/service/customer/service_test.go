package customer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/faroukelkholy/bank/internal/service/models"
	"github.com/faroukelkholy/bank/internal/storage/mocks"
)

var (
	id = "123e4567-e89b-12d3-a456-426614174000"
)

func TestCreateCustomerAccount_Success(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*storage.Account")).Return(nil)

	s := New(rMock)
	err := s.CreateCustomerAccount(id, &models.Account{})

	assert.Nil(t, err)
}

func TestCreateCustomerAccount_Error(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*storage.Account")).Return(errors.New("no account table"))

	s := New(rMock)
	err := s.CreateCustomerAccount(id, &models.Account{})

	assert.NotNil(t, err)
}
