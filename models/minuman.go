package models

import "gorm.io/gorm"

type Minuman struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey" json:"id"`
	Nama  string `json:"nama"`
	Harga uint   `json:"harga"`
}
