package controllers

import (
	"go_restaurant/config"
	"go_restaurant/models"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMinuman(c echo.Context) error {
    nama := c.FormValue("nama")
    hargaString := c.FormValue("harga")

    harga, err := strconv.ParseUint(hargaString, 10, 0)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Format harga tidak valid"})
    }

    minuman := &models.Minuman{
        Nama:  nama,
        Harga: uint(harga), 
    }

    if err := c.Bind(minuman); err != nil {
        return err
    }

    if err := config.DB.Create(&minuman).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat daftar minuman"})
    }

    return c.JSON(http.StatusCreated, minuman)
}


func GetAllMinuman(c echo.Context) error {
	var minumans []models.Minuman

	if err := config.DB.Find(&minumans).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"minumans not found"})
	}

	return c.JSON(http.StatusOK, minumans)
}


//via path url, http://localhost:8080/minuman/read/3. untuk pencarian dengan id prefer yg ini
func GetMinumanById(c echo.Context) error {
	id := c.Param("id")
	minuman := new(models.Minuman)

	if err:= config.DB.First(&minuman, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"minuman dengan id tersebut tidak ada"})
	}

	return c.JSON(http.StatusOK, minuman)
}


//via path params, http://localhost:8080/minuman/read?id=3
func GetMinumanByIdParam(c echo.Context) error {
	id := c.QueryParam("id")
	minuman := new(models.Minuman)

	if err:= config.DB.First(&minuman, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"minuman dengan id tersebut tidak ada"})
	}

	return c.JSON(http.StatusOK, minuman)
}


func UpdateMinumanById(c echo.Context) error {
	id := c.Param("id")
	nama:= c.FormValue("nama")
	hargaString:= c.FormValue("harga")

	harga, err := strconv.ParseUint(hargaString, 10, 0)
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error":"harga tidak dapat di konvesikan"})
	}

	minuman := new(models.Minuman)
	if err := config.DB.First(&minuman, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "minuman dengan id tersebut tidak ditemukan"})
    }

	minuman.Nama = nama
	minuman.Harga = uint(harga)

	if err := config.DB.Save(&minuman).Error; err!= nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"error":"Data tidak dapar tersimpan"})
	}

	return c.JSON(http.StatusOK, minuman)
}


func DeleteMinumanById(c echo.Context) error {
	id := c.Param("id")

	minuman := new(models.Minuman)
	if err := config.DB.First(&minuman, id).Error; err!=nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data tidak ditemukan dengan id"})
	}

	if err := config.DB.Delete(&minuman).Error; err != nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data sudah terhapus"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message":"Data sudah dihapus"})
}