package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"size:256;NOT NULL" json:"title"`
	Description string `gorm:"size:256" json:"description"`
	Status      string `gorm:"type:enum('created', 'updated', 'deleted');default:'created'" json:"status"`
}

func (Task) TableName() string {
	return "task"
}

type CreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}
