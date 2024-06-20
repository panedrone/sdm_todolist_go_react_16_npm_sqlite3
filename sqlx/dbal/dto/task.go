package dto

// Code generated by a tool. DO NOT EDIT.
// Additions may be hand-coded in a separate go-file.
// https://sqldalmaker.sourceforge.net/

type Task struct {
	TId       int64  `json:"t_id" db:"t_id"` // PK
	PId       int64  `json:"p_id" db:"p_id"` // FK ref. column -> Project
	TPriority int64  `json:"t_priority" db:"t_priority"`
	TDate     string `json:"t_date" db:"t_date"`
	TSubject  string `json:"t_subject" db:"t_subject"`
	TComments string `json:"t_comments" db:"t_comments"`
}
