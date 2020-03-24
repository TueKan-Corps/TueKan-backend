package model

//subject model
type Subject struct {
	TagID       int    `json:"tag_id" db:"tag_id"`
	SubjectName string `json:"subject_name" db:"subject_name"`
}
