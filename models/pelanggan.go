package models

type Pelanggan struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Nama string `json:"nama"`
}
