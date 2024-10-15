package model

import "time"

type Deposit struct {
	NIK       string    `gorm:"primaryKey;size:20" json:"nik"`
	FullName  string    `gorm:"size:100;not null" json:"full_name"`
	Address   string    `gorm:"type:text;not null" json:"address"`
	Goals     string    `gorm:"type:text;not null" json:"goals"`
	Target    float64   `gorm:"type:decimal(18,2);not null" json:"target"`
	Unit      string    `gorm:"type:enum('daily','weekly','monthly');not null" json:"unit"`
	Due       int       `gorm:"not null" json:"due"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
