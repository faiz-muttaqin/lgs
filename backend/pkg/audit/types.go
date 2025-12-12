package audit

import "gorm.io/datatypes"

type Entry struct {
	Action     string
	Resource   string
	ResourceID string

	BeforeData datatypes.JSON
	AfterData  datatypes.JSON

	Status  string
	Message string
}
