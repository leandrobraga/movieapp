package model

// RecordID defines a record id. Together with RecordType
// indentifies unique records across all types.
type RecordID string

// RecordType defines a record type. Together with RecordID
// indetifies unique records across all types.
type RecordType string

// Existing record types
const (
	RecordTypeMovies = RecordType("movie")
)

// UserID defines a user id.
type UserID string

type RatingValue int

type Rating struct {
	RecordID   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}
