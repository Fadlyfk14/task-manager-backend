package model

import "time"

type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    string    `json:"deadline"`
	Status      string    `json:"status" gorm:"default:'todo'"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Task) TableName() string { return "task" }

type User struct {
	ID       string `json:"id" gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()"`
	Username string `json:"username" gorm:"column:username;unique;not null"`
	Password string `json:"-" gorm:"column:password;not null"`
	Role     string `json:"role" gorm:"column:role;type:varchar(30);default:user"`
}

func (User) TableName() string { return "users" }

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
