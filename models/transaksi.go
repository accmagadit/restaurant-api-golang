package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaksi struct {
	gorm.Model
	ID            string      `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time   `json:"createdAt"`
	Pegawai       Pegawai `gorm:"foreignKey:PegawaiID"`
	Makanan       Makanan `gorm:"foreignKey:MakananID"`
	Minuman       Minuman `gorm:"foreignKey:MinumanID"`
	JumlahMakanan uint        `json:"jumlahMakanan"`
	JumlahMinuman uint        `json:"jumlahMinuman"`
	TotalHarga    uint        `json:"totalHarga"`
	PegawaiID     uint        `gorm:"not null" json:"pegawaiID"`
	MakananID     uint        `gorm:"not null" json:"makananID"`
	MinumanID     uint        `gorm:"not null" json:"minumanID"`
}
