package models

import (
	"database/sql"
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
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Comments  []Comment    `json:"comments"`
	CreatedAt time.Time    `json:"createdAt"`
	UpdatedAt time.Time    `json:"updatedAt"`
	ClosedAt  sql.NullTime `json:"closedAt" gorm:"null;default:null"`
	StatusID  uint         `json:"statusId" gorm:"default:1;notnull"`
}

func (Task) TableName() string {
	return "task"
}

type Comment struct {
	ID        uint      `json:"id"`
	TaskID    uint      `json:"taskId"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}

func (Comment) TableName() string {
	return "comment"
}

type Status struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
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
