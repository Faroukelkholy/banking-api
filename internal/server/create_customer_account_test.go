package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/faroukelkholy/bank/internal/storage/postgres"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/faroukelkholy/bank/internal/service/customer/mocks"
)

func TestCreateCustomerAccount(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*models.Account")).Return(nil)

	// setup handler
	h := CCAHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/customers/%s/accounts", mocks.Customer.ID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			data := res.Data

			assert.Equal(t, http.StatusCreated, rec.Code)
			assert.Equal(t, "created", data.(string))
		}
	}
}

func TestCreateCustomerAccount_Error(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*models.Account")).Return(errors.New("no customer table"))

	// setup handler
	h := CCAHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/customers/%s/accounts", mocks.Customer.ID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.Equal(t, "internal error", HTTPErr.Title)
		}
	}
}

func TestCreateCustomerAccount_NotFound(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*models.Account")).Return(errors.New(postgres.NoCustomerID))

	// setup handler
	h := CCAHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/customers/%s/accounts", mocks.Customer.ID), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, postgres.NoCustomerID, HTTPErr.Title)
		}
	}
}

func TestCreateCustomerAccount_BadRequest(t *testing.T) {
	// setup mocks
	account := map[string]interface{}{
		"balance": true,
		"name":    "current",
	}

	sMock := &mocks.Service{}
	sMock.On("CreateCustomerAccount", mock.AnythingOfType("string"), mock.AnythingOfType("*models.Account")).Return(nil)

	// setup handler
	h := CCAHandler(sMock)

	JSONBytes, _ := json.Marshal(account)
	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/customers/%s/accounts", mocks.Customer.ID), JSONBytes)

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Equal(t, "account data is not valid", HTTPErr.Title)
		}
	}
}
