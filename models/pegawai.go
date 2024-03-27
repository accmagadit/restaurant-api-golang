package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Pegawai struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey" json:"id"`
	Nama    string `gorm:"not null" json:"nama"`
	Alamat  string `gorm:"not null" json:"alamat"`
	No_telp string `gorm:"not null" json:"no_telp"`
	Username string `gorm:"not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}


func (p *Pegawai) HashPassword() error{
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}	

	p.Password = string(hashedPassword)
	return nil
}

func (p *Pegawai) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
}

func CreateUser(db *gorm.DB, pegawai *Pegawai) error {
	return db.Create(pegawai).Error
}

func FindUserByUsername(db *gorm.DB, username string) (Pegawai, error) {
	var pegawai Pegawai
	err := db.Where("username = ?", username).First(&pegawai).Error
	return pegawai, err

}