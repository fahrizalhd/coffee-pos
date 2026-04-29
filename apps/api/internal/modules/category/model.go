package category

import "time"

type Category struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
