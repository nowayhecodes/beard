package beard

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func ImprovedJsonUnmarshal[T any](target *T, data []byte) []error {
	err := json.Unmarshal(data, target)

	if err != nil {
		return []error{err}
	}

	errorList := make([]error, 0)
	fields := reflect.ValueOf(target).Elem()

	for i := 0; i < fields.NumField(); i++ {
		tags := fields.Type().Field(i).Tag.Get("json")

		if strings.Contains(tags, "required") && fields.Field(i).IsZero() {
			errorList = append(
				errorList,
				fmt.Errorf("[Beard] Missing required field '%s'", fields.Type().Field(i).Name))
		}
	}

	if len(errorList) != 0 {
		return errorList
	}

	return nil
}
