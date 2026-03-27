package models

import "gorm.io/gorm"

// Model pemilih
type Voter struct {
	gorm.Model
	Name        string `gorm:"size:255;not null" json:"name"`
	Email       string `gorm:"unique;not null" json:"email"`
	Voted       bool   `gorm:"default:false" json:"voted"`
	CandidateID uint   `json:"candidate_id"`
	Candidate   Candidate `gorm:"foreignKey:CandidateID"`
}
