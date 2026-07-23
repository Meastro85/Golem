package tool

import (
	"Golem/internal"
	"reflect"
	"strings"
)

// ParamSchema represents the schema for the parameters a tool accepts. It is used to validate the input parameters and generate documentation for the tool.
type ParamSchema struct {
	Type       string                `json:"type"`
	Properties map[string]PropSchema `json:"properties"`
	Required   []string              `json:"required,omitempty"`
}

// PropSchema represents the schema for a single property in the ParamSchema. It includes the type of the property and an optional description.
type PropSchema struct {
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

// SchemaFor generates a ParamSchema for the given struct type. It inspects the fields of the struct and creates a schema based on their types and tags.
func SchemaFor(t reflect.Type) ParamSchema {
	schema := ParamSchema{
		Type:       "object",
		Properties: make(map[string]PropSchema),
	}

	for field := range t.Fields() {
		if !field.IsExported() {
			continue
		}

		jsonName := field.Name
		omitEmpty := false
		if jsonTag, ok := field.Tag.Lookup("json"); ok {
			parts := strings.Split(jsonTag, ",")
			if len(parts) > 0 && parts[0] != "" {
				jsonName = parts[0]
			}

			for _, part := range parts[1:] {
				if part == "omitempty" {
					omitEmpty = true
				}
			}
		}

		schema.Properties[jsonName] = PropSchema{
			Type:        internal.JsonType(field.Type),
			Description: field.Tag.Get("description"),
		}
		if !omitEmpty {
			schema.Required = append(schema.Required, jsonName)
		}
	}

	return schema
}
