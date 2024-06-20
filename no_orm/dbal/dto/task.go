package dto

// Code generated by a tool. DO NOT EDIT.
// Additions may be hand-coded in a separate go-file.
// https://sqldalmaker.sourceforge.net/

type Task struct {
	TId       int64  `json:"t_id"` // PK
	PId       int64  `json:"p_id"` // FK ref. column -> Project
	TPriority int64  `json:"t_priority"`
	TDate     string `json:"t_date"`
	TSubject  string `json:"t_subject"`
	TComments string `json:"t_comments"`
}