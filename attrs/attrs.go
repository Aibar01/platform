package attrs

import (
	"fmt"
	"reflect"
	"strings"
)

type ContextAttrs struct {
	XAuthToken    string `header:"X-Authorization" json:"X-Authorization,omitempty"`
	AuthToken     string `header:"Authorization" json:"Authorization"`
	Locale        string `header:"Accept-Language" json:"Accept-Language"`
	CorrelationId string `header:"X-Correlation-Id" json:"X-Correlation-Id"`
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
