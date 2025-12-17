package filter

import (
	"reflect"
	"strings"
	"time"
)

func BuildSchemaFromStruct(model any) map[string]FieldSchema {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	schema := map[string]FieldSchema{}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		jsonTag := strings.Split(f.Tag.Get("json"), ",")[0]
		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		ui := f.Tag.Get("ui")
		gormTag := f.Tag.Get("gorm")

		column := jsonTag
		for _, part := range strings.Split(gormTag, ";") {
			if strings.HasPrefix(part, "column:") {
				column = strings.TrimPrefix(part, "column:")
			}
		}

		fieldType, ops := inferTypeAndOps(f.Type)

		schema[jsonTag] = FieldSchema{
			JSONKey:    jsonTag,
			DBColumn:   column,
			Type:       fieldType,
			Operators:  ops,
			Sortable:   strings.Contains(ui, "sortable"),
			Filterable: strings.Contains(ui, "filterable"),
			Editable:   strings.Contains(ui, "editable"),
			Visible:    strings.Contains(ui, "visible"),
			Selection:  extractSelection(ui),
			TimeFormat: f.Tag.Get("time_format"),
		}
	}

	return schema
}

func inferTypeAndOps(t reflect.Type) (FieldType, []Operator) {
	switch t {
	case reflect.TypeOf(time.Time{}):
		return DateTime, []Operator{Gt, Gte, Lt, Lte, Between, IsNull}
	}

	switch t.Kind() {
	case reflect.String:
		return String, []Operator{Eq, Ne, Contains, StartsWith, EndsWith, In, IsNull}
	case reflect.Int, reflect.Int64, reflect.Float64, reflect.Float32:
		return Number, []Operator{Eq, Ne, Gt, Gte, Lt, Lte, Between, In, IsNull}
	case reflect.Bool:
		return Boolean, []Operator{Eq, IsNull}
	}

	return String, []Operator{Eq}
}

func extractSelection(ui string) string {
	for _, part := range strings.Split(ui, ";") {
		if strings.HasPrefix(part, "selection:") {
			return strings.TrimPrefix(part, "selection:")
		}
	}
	return ""
}
