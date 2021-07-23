package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/faroukelkholy/bank/internal/service/account/mocks"
)

func TestGetAccountTransactions(t *testing.T) {
	// setup mocks
	transactions := mocks.Transactions

	sMock := &mocks.Service{}
	sMock.On("GetAccountTransactions", mock.AnythingOfType("string")).Return(transactions, nil)

	// setup handler
	h := GATsHandler(sMock)

	c, rec := setupTest(http.MethodGet, fmt.Sprintf("/accounts/%s/transactions",mocks.Account.ID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			data := res.Data

			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, transactions[0].Amount, data.([]interface{})[0].(map[string]interface{})["amount"])
		}
	}
}

func TestGetAccountTransactions_Error(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("GetAccountTransactions", mock.AnythingOfType("string")).Return(nil, errors.New("no account table"))

	// setup handler
	h := GATsHandler(sMock)

	c, rec := setupTest(http.MethodGet, fmt.Sprintf("/accounts/%s/transactions",mocks.Account.ID), []byte{})

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

func TestGetAccountTransactions_NotFound(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("GetAccountTransactions", mock.AnythingOfType("string")).Return(nil, nil)

	// setup handler
	h := GATsHandler(sMock)

	c, rec := setupTest(http.MethodGet, fmt.Sprintf("/accounts/%s/transactions",mocks.Account.ID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, "transactions not found", HTTPErr.Title)
		}
	}
}
