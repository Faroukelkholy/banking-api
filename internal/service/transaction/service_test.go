package transaction

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mocksData "github.com/faroukelkholy/bank/internal/service/transaction/mocks"
	"github.com/faroukelkholy/bank/internal/storage/mocks"
)


func TestCreateTransaction_Success(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateTransaction",mock.AnythingOfType("*storage.Transaction")).Return(nil)

	s := New(rMock)
	err := s.CreateTransaction(mocksData.Transfer)

	assert.Nil(t, err)
}

func TestCreateTransaction_Error(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateTransaction", mock.AnythingOfType("*storage.Transaction")).Return(errors.New("no transaction table"))

	s := New(rMock)
	err := s.CreateTransaction(mocksData.Transfer)

	assert.NotNil(t, err)
}

func TestCreateTransaction_BusinessRule_Error(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("CreateTransaction", mock.AnythingOfType("*storage.Transaction")).Return(errors.New("no transaction table"))

	s := New(rMock)
	mocksData.Transfer.Amount = 0
	err := s.CreateTransaction(mocksData.Transfer)

	assert.NotNil(t, err)
	assert.EqualValues(t, "amount must be positive",err.Error())
}


