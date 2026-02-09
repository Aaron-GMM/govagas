package schemas

import (
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
