package account

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/faroukelkholy/bank/internal/storage/mocks"
)

var (
	id = "123e4567-e89b-12d3-a456-426614174000"
)

func TestGetAccount_Success(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("GetAccount", mock.AnythingOfType("string")).Return(mocks.Account, nil)

	s := New(rMock)
	result, err := s.GetAccount(id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, mocks.Account.Name, result.Name)
	assert.EqualValues(t, mocks.Account.Balance, result.Balance)
}

func TestGetAccount_Error(t *testing.T) {
	var cases = []struct {
		title string
	}{
		{
			"caseNoRows",
		},
		{
			"caseDBError",
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			switch test.title {
			case "caseNoRows":
				rMock := &mocks.Repository{}
				rMock.On("GetAccount", mock.AnythingOfType("string")).Return(nil, nil)

				s := New(rMock)
				result, err := s.GetAccount(id)

				assert.Nil(t, err)
				assert.Nil(t, result)
			case "caseDBError":
				rMock := &mocks.Repository{}
				rMock.On("GetAccount", mock.AnythingOfType("string")).Return(nil, errors.New("no account table"))

				s := New(rMock)
				result, err := s.GetAccount(id)

				assert.NotNil(t, err)
				assert.Nil(t, result)
			}
		})
	}
}

func TestTestGetAccountTransactions_Success(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("GetAccountTransactions", mock.AnythingOfType("string")).Return(mocks.Transactions, nil)

	s := New(rMock)
	result, err := s.GetAccountTransactions(id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, mocks.Transactions[0].Amount, result[0].Amount)
}

func TestTestGetAccountTransactions_Error(t *testing.T) {
	var cases = []struct {
		title string
	}{
		{
			"caseNoRows",
		},
		{
			"caseDBError",
		},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			switch test.title {
			case "caseNoRows":
				rMock := &mocks.Repository{}
				rMock.On("GetAccountTransactions", mock.AnythingOfType("string")).Return(nil, nil)

				s := New(rMock)
				result, err := s.GetAccountTransactions(id)

				assert.Nil(t, err)
				assert.Nil(t, result)
			case "caseDBError":
				rMock := &mocks.Repository{}
				rMock.On("GetAccountTransactions", mock.AnythingOfType("string")).Return(nil, errors.New("no account table"))

				s := New(rMock)
				result, err := s.GetAccountTransactions(id)

				assert.NotNil(t, err)
				assert.Nil(t, result)
			}
		})
	}
}