package account

import (
	"errors"
	"testing"

	"github.com/faroukelkholy/bank/internal/storage/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	id = "123e4567-e89b-12d3-a456-426614174000"
)

func TestGA_Success(t *testing.T) {
	rMock := &mocks.Repository{}
	rMock.On("GetAccount", mock.AnythingOfType("string")).Return(mocks.AccEntity, nil)

	srv := New(rMock)
	result, err := srv.GetAccount(id)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.EqualValues(t, mocks.AccEntity.Name, result.Name)
	assert.EqualValues(t, mocks.AccEntity.Balance, result.Balance)
}

func TestGetChemical_Error(t *testing.T) {
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

				srv := New(rMock)
				result, err := srv.GetAccount(id)

				assert.Nil(t, err)
				assert.Nil(t, result)
			case "caseDBError":
				rMock := &mocks.Repository{}
				rMock.On("GetAccount", mock.AnythingOfType("string")).Return(nil, errors.New("no account table"))

				srv := New(rMock)
				result, err := srv.GetAccount(id)

				assert.NotNil(t, err)
				assert.Nil(t, result)
			}
		})
	}
}
