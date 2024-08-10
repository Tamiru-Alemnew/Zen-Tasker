package models


type Task struct {
    ID          int                `json:"id" bson:"_id,omitempty"` 
    Title       string             `json:"title" bson:"title"`
    Description string             `json:"description" bson:"description"`
    DueDate     string             `json:"due_date" bson:"due_date"`
    Status      string             `json:"status" bson:"status"`
    UserID      int                `json:"user_id" bson:"user_id"` // Use ObjectID for User reference
}