package domain

import (
	"context"
	"errors"
	"time"
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

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type JWTService interface {
	GenerateToken(userID int, username, role string) (string, error)
    ParseToken(tokenString string) (*TokenClaims, error)
}

// TaskRepository defines the interface for task-related data operations.
type TaskRepository interface {
    Create(ctx context.Context, task *Task) error
    Update(ctx context.Context, id int, task *Task) error
    Delete(ctx context.Context, id int) error
    GetAll(ctx context.Context) ([]Task, error)
    GetByID(ctx context.Context, id int) (*Task, error)
}

type TaskUsecase interface {
	Create(ctx context.Context, task *Task) (*Task, error)
	Update(ctx context.Context, id int, task *Task) (*Task, error)
	GetAll(ctx context.Context) ([]Task, error)
	GetByID(ctx context.Context, id int) (*Task, error)
	Delete(ctx context.Context, id int) error
}

type UserRepository interface {
    Create(ctx context.Context, user *User) error
    FindByUsername(ctx context.Context, username string) (*User, error)
    GetAll(ctx context.Context) ([]User, error)
    Promote(ctx context.Context, id int) error
}

type UserUsecase interface {
    SignUp(ctx context.Context, user *User) (*User , error)
    Login(ctx context.Context , username , pasword string) (*User, string ,error)
    Promote(ctx context.Context, id int) error
}


type TokenClaims struct {
    UserID   string `json:"user_id"`
    Username string `json:"username"`
    Role     string `json:"role"`
    Exp      int64  `json:"exp"` 
}

// Valid method to satisfy the jwt.Claims interface
func (claims *TokenClaims) Valid() error {
    if time.Unix(claims.Exp, 0).Before(time.Now()) {
        return errors.New("token has expired")
    }
    return nil
}
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

