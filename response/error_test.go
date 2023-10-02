package response

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResponseError(t *testing.T) {
	errorResponse := NewError("some error")
	assert.Equal(t, errorResponse.Detail, "some error")
}

func TestPermissionDeniedError(t *testing.T) {
	assert.Equal(t, PermissionDeniedError, NewError("permission denied"))
}
