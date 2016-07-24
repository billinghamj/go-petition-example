package log

// Error implements the official Cuvva error structure
type Error struct {
	Code string                 `json:"code"`
	Meta map[string]interface{} `json:"meta"`
}

// CreateError returns a new Error
func CreateError(code string, meta map[string]interface{}) *Error {
	return &Error{code, meta}
}

func (e *Error) Error() string {
	return e.Code
}
