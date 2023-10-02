package attrs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAttrsHeaderTag(t *testing.T) {
	attrs := ContextAttrs{
		XAuthToken:    "whatever",
		AuthToken:     "whatever",
		Locale:        "ru",
		CorrelationId: "14b5ce2b-4475-4bc8-a70b-78143518a5d9",
	}
	expected := map[string]string{
		"X-Authorization":  "whatever",
		"Authorization":    "whatever",
		"Accept-Language":  "ru",
		"X-Correlation-Id": "14b5ce2b-4475-4bc8-a70b-78143518a5d9",
	}

	assert.Equal(t, attrs.GetByTagName("header"), expected)
}

func TestAttrsJSONTag(t *testing.T) {
	attrs := ContextAttrs{
		XAuthToken:    "whatever",
		AuthToken:     "whatever",
		Locale:        "ru",
		CorrelationId: "14b5ce2b-4475-4bc8-a70b-78143518a5d9",
	}
	expected := map[string]string{
		"x_auth_token":   "whatever",
		"auth_token":     "whatever",
		"locale":         "ru",
		"correlation_id": "14b5ce2b-4475-4bc8-a70b-78143518a5d9",
	}

	assert.Equal(t, attrs.GetByTagName("json"), expected)
}
