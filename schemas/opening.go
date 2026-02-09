package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Name             string
	Publication_data string
	Role             string
	Company          string
	Locate           string
	Remote           bool
	Link             string
	Salary           float32
	Description      string
}

type OpeningResponse struct {
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	DeletedAt        time.Time `json:"deletedAt,omitempty"`
	Name             string    `json:"name"`
	Publication_data string    `json:"publication_data"`
	Role             string    `json:"role"`
	Company          string    `json:"company"`
	Locate           string    `json:"locate"`
	Remote           bool      `json:"remote"`
	Link             string    `json:"link"`
	Salary           float32   `json:"salary"`
	Description      string    `json:"description"`
}
