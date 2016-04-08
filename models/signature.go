package models

// Signature represents some kind of signature (maybe for a petition?)
// Currently the same class is used for all cases
type Signature struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
	Message   string
}
