// package models

// import "time"

// type Task struct {
// 	ID              uint       `json:"id" gorm:"primaryKey"`
// 	Title           string     `json:"title" gorm:"not null"`
// 	Description     string     `json:"description" gorm:"not null"`
// 	Complete        bool       `json:"complete" gorm:"not null"`
// 	ReminderEnabled bool       `json:"reminder_enabled"`
// 	ReminderAt      *time.Time `json:"reminder_at"`
// 	ReminderSent    bool       `json:"reminder_sent"`
// 	UpdatedAt       time.Time  `json:"updated_at"`
// 	CreatedAt       time.Time  `json:"created_at"`
// }

package models

import "time"

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	Complete    bool   `json:"complete" gorm:"not null"`

	ReminderEnabled bool       `json:"reminder_enabled" gorm:"default:false"`
	ReminderAt      *time.Time `json:"reminder_at"`
	ReminderSent    bool       `json:"reminder_sent" gorm:"default:false"`

	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
