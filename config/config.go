package config

import (
	"go_restaurant/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDataBase()  {
	dsn := "root:@tcp(127.0.0.1:3306)/restaurant?parseTime=true"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		panic("Tidak Dapat terhubung pada database")
	}
}

func Migration()  {
	DB.AutoMigrate(&models.Makanan{})
	DB.AutoMigrate(&models.Minuman{})	
	DB.AutoMigrate(&models.Pegawai{})	
	DB.AutoMigrate(&models.Pelanggan{})	
	DB.AutoMigrate(&models.Transaksi{})	
}