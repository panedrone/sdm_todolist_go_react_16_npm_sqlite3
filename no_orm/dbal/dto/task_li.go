package dto

// Code generated by a tool. DO NOT EDIT.
// https://sqldalmaker.sourceforge.net/

// TaskLi `task` list item: [-] p_id, [-] t_comments
type TaskLi struct {
	TId       int64  `json:"t_id"`
	TPriority int64  `json:"t_priority"`
	TDate     string `json:"t_date"`
	TSubject  string `json:"t_subject"`
}
