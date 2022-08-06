package models

import (
	"time"
)

/*
// v0.0.1 won't have users
type User struct {
	ID       uint
	Username string
}
*/

type Task struct {
	// UserID    uint
	ID        uint
	Title     string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
	ClosedAt  time.Time
	StatusID  uint `gorm:"default:1;notnull"`
}

func (Task) TableName() string {
	return "task"
}

type Comment struct {
	ID        uint
	TaskID    uint
	Message   string
	CreatedAt time.Time
}

func (Comment) TableName() string {
	return "comment"
}

type Status struct {
	ID   uint
	Name string
}

func (Status) TableName() string {
	return "status"
}

var (
	StatusOpen       = Status{1, "open"}
	StatusInProgress = Status{2, "inProgress"}
	StatusClosed     = Status{3, "closed"}
)

var StatusTypes = []Status{
	StatusOpen,
	StatusInProgress,
	StatusClosed,
}

func GetStatusMap() map[string]uint {
	sMap := make(map[string]uint)
	for _, status := range StatusTypes {
		sMap[status.Name] = status.ID
	}

	return sMap
}
