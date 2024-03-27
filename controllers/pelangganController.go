package controllers

import (
	"github.com/labstack/echo/v4"
	"go_restaurant/config"
	"go_restaurant/models"
	"net/http"
)

func CreatePelanggan(c echo.Context) error {
	nama := c.FormValue("nama")

	pelanggan := &models.Pelanggan{
		Nama: nama,
	}

	if err := c.Bind(pelanggan); err != nil{
		return err
	}

	if err := config.DB.Create(&pelanggan).Error; err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"error":"Gagal nyambung ke database"})
	}

	return c.JSON(http.StatusCreated, pelanggan)
}


func GetAllPelanggan(c echo.Context) error {
	var pelanggans []models.Pelanggan

	if err:= config.DB.Find(&pelanggans).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"pelanggans not found"})
	}

	return c.JSON(http.StatusOK, pelanggans)
}


func GetPelangganById(c echo.Context) error{
	id := c.Param("id")
	pelanggan := new(models.Pegawai)

	if err := config.DB.First(&pelanggan, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"pelanggan dengan id itu not found"})
	}

	return c.JSON(http.StatusOK, pelanggan)
}



func UpdatePelangganById(c echo.Context) error {
	id := c.Param("id")
	nama:= c.FormValue("nama")

	pelanggan := new(models.Pelanggan)
	if err := config.DB.First(&pelanggan, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "pelanggan dengan id tersebut tidak ditemukan"})
    }

	pelanggan.Nama = nama

	if err := config.DB.Save(&pelanggan).Error; err!= nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"error":"Data tidak dapar tersimpan"})
	}

	return c.JSON(http.StatusOK, pelanggan)
}


func DeletePelangganById(c echo.Context) error {
	id := c.Param("id")

	pelanggan := new(models.Pelanggan)
	if err := config.DB.First(&pelanggan, id).Error; err!=nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data tidak ditemukan dengan id"})
	}

	if err := config.DB.Delete(&pelanggan).Error; err != nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data sudah terhapus"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message":"Data sudah dihapus"})
}