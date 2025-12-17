package filter

type FieldType string

const (
	String   FieldType = "string"
	Number   FieldType = "number"
	Boolean  FieldType = "boolean"
	DateTime FieldType = "datetime"
)

type Operator string

const (
	Eq         Operator = "eq" // equal
	Ne         Operator = "ne" // not equal
	Contains   Operator = "contains"
	StartsWith Operator = "starts_with"
	EndsWith   Operator = "ends_with"
	In         Operator = "in"      // comma-separated values
	Gt         Operator = "gt"      // greater than
	Gte        Operator = "gte"     // greater than or equal to
	Lt         Operator = "lt"      // less than
	Lte        Operator = "lte"     // less than or equal to
	Between    Operator = "between" // two comma-separated values
	IsNull     Operator = "is_null" // true or false
)

type FieldSchema struct {
	JSONKey    string     `json:"key"`
	DBColumn   string     `json:"column"`
	Type       FieldType  `json:"type"`
	Operators  []Operator `json:"operators"`
	Sortable   bool       `json:"sortable"`
	Filterable bool       `json:"filterable"`
	Editable   bool       `json:"editable"`
	Visible    bool       `json:"visible"`
	Selection  string     `json:"selection,omitempty"`
	TimeFormat string     `json:"time_format,omitempty"`
}

// FilterError represents a detailed filter error
type FilterError struct {
	Code     string     `json:"code"`
	Message  string     `json:"message"`
	Field    string     `json:"field,omitempty"`
	Operator string     `json:"operator,omitempty"`
	Allowed  []Operator `json:"allowed,omitempty"`
}

func (e *FilterError) Error() string {
	return e.Message
}

// Error codes
const (
	ErrInvalidField       = "INVALID_FIELD"
	ErrFieldNotFilterable = "FIELD_NOT_FILTERABLE"
	ErrInvalidOperator    = "INVALID_OPERATOR"
	ErrOperatorNotAllowed = "OPERATOR_NOT_ALLOWED"
	ErrFieldNotSortable   = "FIELD_NOT_SORTABLE"
	ErrInvalidFieldName   = "INVALID_FIELD_NAME"
)
