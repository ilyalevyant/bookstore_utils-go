package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("message", errors.New("db error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "db error", err.Causes[0])
}

func TestNewBadRequestError(t *testing.T) {

}

func TestNewNotFoundError(t *testing.T) {

}

func TestNewError(t *testing.T) {

}