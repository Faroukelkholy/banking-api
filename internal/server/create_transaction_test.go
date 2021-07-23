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

	"github.com/faroukelkholy/bank/internal/service/transaction/mocks"
)

func TestCreateTransaction(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("CreateTransaction", mock.AnythingOfType("*models.Transfer")).Return( nil)

	// setup handler
	h := CTHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/transactions"), []byte{})

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

func TestCreateTransaction_Error(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("CreateTransaction", mock.AnythingOfType("*models.Transfer")).Return( errors.New("no transaction table"))

	// setup handler
	h := CTHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/transactions"), []byte{})

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

func TestCreateTransaction_NotFound(t *testing.T) {
	// setup mocks
	sMock := &mocks.Service{}
	sMock.On("CreateTransaction", mock.AnythingOfType("*models.Transfer")).Return( errors.New(postgres.NoAccountID))

	// setup handler
	h := CTHandler(sMock)

	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/transactions"), []byte{})

	// assertions
	var res HTTPResponse
	if assert.NoError(t, h(c)) {
		if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res)) {
			HTTPErr := res.Err

			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.Equal(t, postgres.NoAccountID, HTTPErr.Title)
		}
	}
}

func TestCreateTransaction_BadRequest(t *testing.T) {
	// setup mocks
	transfer := map[string]interface{}{
		"sender": "",
		"receiver":    "",
		"amount": true,
	}

	sMock := &mocks.Service{}
	sMock.On("CreateTransaction", mock.AnythingOfType("*models.Transfer")).Return( nil)

	// setup handler
	h := CTHandler(sMock)

	JSONBytes, _ := json.Marshal(transfer)
	c, rec := setupTest(http.MethodPost, fmt.Sprintf("/transactions"), JSONBytes)

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



