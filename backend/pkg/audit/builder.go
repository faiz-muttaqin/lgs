package audit

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
)

const (
	StatusSuccess = "SUCCESS"
	StatusFailed  = "FAILED"
)

func Create(resource string, id any) *Entry {
	return &Entry{
		Action:     "CREATE",
		Resource:   resource,
		ResourceID: fmt.Sprint(id),
	}
}

func Update(resource string, id any) *Entry {
	return &Entry{
		Action:     "UPDATE",
		Resource:   resource,
		ResourceID: fmt.Sprint(id),
	}
}

func Delete(resource string, id any) *Entry {
	return &Entry{
		Action:     "DELETE",
		Resource:   resource,
		ResourceID: fmt.Sprint(id),
	}
}

func (e *Entry) Before(v any) *Entry {
	if v != nil {
		b, _ := json.Marshal(v)
		e.BeforeData = datatypes.JSON(b)
	}
	return e
}

func (e *Entry) After(v any) *Entry {
	if v != nil {
		b, _ := json.Marshal(v)
		e.AfterData = datatypes.JSON(b)
	}
	return e
}

func (e *Entry) Success(msg ...string) *Entry {
	e.Status = StatusSuccess
	if len(msg) > 0 {
		e.Message = msg[0]
	}
	return e
}

func (e *Entry) Failed(err error) *Entry {
	e.Status = StatusFailed
	if err != nil {
		e.Message = err.Error()
	}
	return e
}
