package domain

import (
	"context"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
)

type Task struct {
    ID          int                `json:"id" bson:"_id,omitempty"` 
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    DueDate     string             `json:"due_date" bson:"due_date"`
    Status      string             `json:"status" bson:"status"`
    UserID      int                `json:"user_id" bson:"user_id"` // Use ObjectID for User reference
}

type User struct {
    ID       int                `bson:"_id,omitempty" json:"id,omitempty"`
    Username string             `bson:"username" json:"username"`
    Password string             `bson:"password" json:"-"`
    Role     string             `bson:"role" json:"role"`
}


// TaskRepository defines the interface for task-related data operations.
type TaskRepository interface {
    Create(ctx context.Context, task *Task) error
    Update(ctx context.Context, id int, task *Task) error
    Delete(ctx context.Context, id int) error
    GetAll(ctx context.Context) ([]Task, error)
    GetByID(ctx context.Context, id int) (*Task, error)
}

type UserRepository interface {
    Create(ctx context.Context, user *User) error
    FindByUsername(ctx context.Context, username string) (*User, error)
    Promote(ctx context.Context, id int) error
}