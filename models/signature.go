package models

// Signature represents some kind of signature (maybe for a petition?)
// Currently the same class is used for all cases
type Signature struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       int    `json:"age"`
	Message   string `json:"message"`
}
