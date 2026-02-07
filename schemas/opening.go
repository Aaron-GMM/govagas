package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	id               int
	name             string
	publication_data string
	role             string
	company          string
	Locate           string
	remote           bool
	link             string
	salary           float32
	description      string
}
