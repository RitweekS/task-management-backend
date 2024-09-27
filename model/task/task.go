package task

type CreateTask struct {
    // Id          int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`   
    UserId      int        `gorm:"column:user_id" json:"user_id"`                  
    Title       string     `gorm:"column:title;not null" json:"title"`          
    Description *string    `gorm:"column:description" json:"description,omitempty"`
    DueDate     string      `gorm:"column:due_date" json:"due_date"`
    Priority    string     `gorm:"column:priority;default:normal" json:"priority"`
}

type GetTask struct {
    Id          int        `json:"id"`   
    UserId      int        `json:"user_id"`                  
    Title       string     `json:"title"`          
    Description *string    `json:"description,omitempty"`
    DueDate     string      `json:"due_date"`
    Priority    string     `json:"priority"`
    IsCompleted bool       `json:"is_completed"` 

}

type UpdateTask struct {
    // Id          int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`   
    // UserId      int        `gorm:"column:user_id" json:"user_id"`                  
    Title       string     `gorm:"column:title;not null" json:"title"`          
    Description *string    `gorm:"column:description" json:"description,omitempty"`
    DueDate     string      `gorm:"column:due_date" json:"due_date"`
    Priority    string     `gorm:"column:priority;default:normal" json:"priority"`
    IsCompleted bool       `json:"is_completed"` 
}