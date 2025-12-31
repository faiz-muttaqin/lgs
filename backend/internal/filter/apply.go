package filter

import (
	"fmt"
	"strings"

	"github.com/faiz-muttaqin/lgs/backend/pkg/util"
	"gorm.io/gorm"
)

func ApplyQueryFilters(db *gorm.DB, query map[string][]string, schema map[string]FieldSchema) (*gorm.DB, error) {
	for rawKey, values := range query {

		field, op := parseKey(rawKey)

		// Skip non-filter query params (pagination, sorting, etc.)
		if util.Contains([]string{"draw", "start", "length", "sort", "fields", "schema"}, field) {
			continue
		}

		meta, ok := schema[field]
		if !ok {
			// Check if it's a typo or unknown field
			return nil, &FilterError{
				Code:    ErrInvalidField,
				Message: fmt.Sprintf("Field '%s' does not exist or is not available for filtering", field),
				Field:   field,
			}
		}

		if !meta.Filterable {
			return nil, &FilterError{
				Code:    ErrFieldNotFilterable,
				Message: fmt.Sprintf("Field '%s' is not filterable", field),
				Field:   field,
			}
		}

		// Validate operator
		if !util.Contains(meta.Operators, Operator(op)) {
			return nil, &FilterError{
				Code:     ErrOperatorNotAllowed,
				Message:  fmt.Sprintf("Operator '%s' is not allowed for field '%s'", op, field),
				Field:    field,
				Operator: op,
				Allowed:  meta.Operators,
			}
		}

		for _, val := range values {
			var err error
			db, err = applyCondition(db, meta, Operator(op), val)
			if err != nil {
				return nil, err
			}
		}
	}
	return db, nil
}

func parseKey(k string) (string, string) {
	if !strings.Contains(k, "[") {
		return k, "eq"
	}
	f := k[:strings.Index(k, "[")]
	op := strings.TrimSuffix(strings.Split(k, "[")[1], "]")
	return f, op
}

func ApplySorting(q *gorm.DB, sort string, schema map[string]FieldSchema) (*gorm.DB, error) {
	if sort == "" {
		return q, nil
	}

	for _, s := range strings.Split(sort, ",") {
		s = strings.TrimSpace(s)
		dir := "asc"
		if strings.HasPrefix(s, "-") {
			dir = "desc"
			s = s[1:]
		}

		f, ok := schema[s]
		if !ok {
			return nil, &FilterError{
				Code:    ErrInvalidField,
				Message: fmt.Sprintf("Sort field '%s' does not exist", s),
				Field:   s,
			}
		}

		if !f.Sortable && s != "id" {
			return nil, &FilterError{
				Code:    ErrFieldNotSortable,
				Message: fmt.Sprintf("Field '%s' is not sortable", s),
				Field:   s,
			}
		}

		q = q.Order(f.DBColumn + " " + dir)
	}
	return q, nil
}
