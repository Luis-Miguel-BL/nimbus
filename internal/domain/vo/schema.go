package vo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Schema struct {
	schema map[SchemaKey]SchemaKind
}

type SchemaKey string
type SchemaKind string

const (
	SchemaKindString SchemaKind = "string"
	SchemaKindNumber SchemaKind = "number"
	SchemaKindBool   SchemaKind = "boolean"
)

func (v *Schema) ValidateItem(fileItem *FileItem) (err error) {
	for dataKey, dataValue := range fileItem.data {
		schemaKind, findSchema := v.schema[SchemaKey(dataKey)]
		if !findSchema {
			return fmt.Errorf("unexpected data %s", dataKey)
		}

		valueKind := reflect.TypeOf(dataValue).Kind()
		if isValidKind(valueKind, schemaKind) {
			continue
		}
		if valueKind == reflect.String {
			newValue, err := convertToKind(dataValue.(string), schemaKind)
			if err != nil {
				return fmt.Errorf("cannot convert %s to %s", dataKey, schemaKind)
			}
			fileItem.data[dataKey] = newValue
		}
	}
	return nil
}

var mapAvailableKinds = map[SchemaKind][]reflect.Kind{
	SchemaKindString: {reflect.String},
	SchemaKindBool:   {reflect.Bool},
	SchemaKindNumber: {reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64},
}

func isValidKind(valueKind reflect.Kind, schemaKind SchemaKind) (isValid bool) {
	availableKinds := mapAvailableKinds[schemaKind]

	for _, availableKind := range availableKinds {
		if availableKind == valueKind {
			return true
		}
	}
	return false
}

const floatSeparator = "."

func convertToKind(value string, schemaKind SchemaKind) (newValue any, err error) {
	switch schemaKind {
	case SchemaKindBool:
		return strconv.ParseBool(value)
	case SchemaKindNumber:
		if strings.Contains(value, floatSeparator) {
			return strconv.ParseFloat(value, 64)
		}
		return strconv.Atoi(value)
	}

	return value, fmt.Errorf("invalid kind")
}
