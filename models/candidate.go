package models

import (
	"gorm.io/gorm"
)

type Candidate struct {
	gorm.Model
	Name    string `json:"name"`
	Votes   int    `json:"votes"`
	Country string `json:"country"`
}
