package models

// Code generated by a tool. DO NOT EDIT.
// Additions may be hand-coded in a separate go-file.
// https://sqldalmaker.sourceforge.net/

// Task can be used in "AutoMigrate"
type Task struct {
	TId       int64   `json:"t_id" gorm:"column:t_id;primaryKey;autoIncrement"` // PK
	PId       int64   `json:"p_id" gorm:"column:p_id"`                          // FK ref. column -> Project
	TPriority int64   `json:"t_priority" gorm:"column:t_priority;not null"`
	TDate     string  `json:"t_date" gorm:"column:t_date;not null"`
	TSubject  string  `json:"t_subject" gorm:"column:t_subject;not null"`
	TComments string  `json:"t_comments" gorm:"column:t_comments;not null"`
	Project   Project `gorm:"foreignKey:PId;references:PId"` // FK -> Project
}