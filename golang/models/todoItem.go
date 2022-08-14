package models

type TodoItem struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"Title" gorm:"not null"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed" gorm:"not null DEFAULT false"`
}
