package attrs

import (
	"fmt"
	"reflect"
	"strings"
)

type ContextAttrs struct {
	XAuthToken    string `header:"X-Authorization" json:"x_auth_token,omitempty"`
	AuthToken     string `header:"Authorization" json:"auth_token"`
	Locale        string `header:"Accept-Language" json:"locale"`
	CorrelationId string `header:"X-Correlation-Id" json:"correlation_id"`
}

func (a *ContextAttrs) GetByTagName(tagName string) map[string]string {
	result := make(map[string]string)
	val := reflect.ValueOf(a).Elem()
	typ := reflect.TypeOf(*a)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get(tagName)

		if parts := strings.Split(tag, ","); len(parts) > 1 {
			tag = parts[0]
		}

		if tag != "" {
			result[tag] = fmt.Sprintf("%v", field.Interface())
		}
	}

	return result
}
