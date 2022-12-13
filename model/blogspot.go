package model

type BlogSpot struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (BlogSpot) TableName() string {
	return "blog_spot"
}
