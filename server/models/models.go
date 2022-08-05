package models

import "time"

type Task struct {
	ID        uint
	Title     string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  time.Time
	Status    StatusType
}

type Comment struct {
	ID        uint
	TaskID    uint
	Message   string
	CreatedAt time.Time
}

type StatusType int

const (
	StatusOpen       StatusType = 0
	StatusInProgress StatusType = 1
	StatusClosed     StatusType = 2
)
