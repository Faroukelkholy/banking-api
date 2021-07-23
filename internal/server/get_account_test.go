package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/faroukelkholy/bank/internal/service/account/mocks"
)

func TestGetAccount(t *testing.T) {
	// setup mocks
	account := mocks.Account

	sMock := &mocks.Service{}
	sMock.On("GetAccount", mock.AnythingOfType("string")).Return(account, nil)

	// setup handler
	h := GAHandler(sMock)

	c, rec := setupTest(http.MethodGet, "/accounts", []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			data := res.Data

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, account.Name, data.(map[string]interface{})["name"])
			assert.Equal(t, account.Balance, data.(map[string]interface{})["balance"])
		}
	}
}

func TestGetAccount_Error(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("GetAccount", mock.AnythingOfType("string")).Return(nil, errors.New("no account table"))

	// setup handler
	h := GAHandler(sMock)

	c, rec := setupTest(http.MethodGet, "/accounts", []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "internal error", HTTPErr.Title)
			assert.Equal(t, "", HTTPErr.Description)
		}
	}
}

func TestGetAccount_NotFound(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("GetAccount", mock.AnythingOfType("string")).Return(nil, nil)

	// setup handler
	h := GAHandler(sMock)

	c, rec := setupTest(http.MethodGet, "/accounts", []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "account not found", HTTPErr.Title)
			assert.Equal(t, "", HTTPErr.Description)
		}
	}
}

