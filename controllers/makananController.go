package controllers

import (
	"go_restaurant/config"
	"go_restaurant/models"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMakanan(c echo.Context) error {
    nama := c.FormValue("nama")
    hargaString := c.FormValue("harga")

    harga, err := strconv.ParseUint(hargaString, 10, 0)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Format harga tidak valid"})
    }

    makanan := &models.Makanan{
        Nama:  nama,
        Harga: uint(harga), 
    }

    if err := c.Bind(makanan); err != nil {
        return err
    }

    if err := config.DB.Create(&makanan).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat daftar makanan"})
    }

    return c.JSON(http.StatusCreated, makanan)
}


func GetAllMakanan(c echo.Context) error {
	var makanans []models.Makanan

	if err := config.DB.Find(&makanans).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"makanans not found"})
	}

	return c.JSON(http.StatusOK, makanans)
}


//via path url, http://localhost:8080/makanan/read/3. untuk pencarian dengan id prefer yg ini
func GetMakananById(c echo.Context) error {
	id := c.Param("id")
	makanan := new(models.Makanan)

	if err:= config.DB.First(&makanan, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"makanan dengan id tersebut tidak ada"})
	}

	return c.JSON(http.StatusOK, makanan)
}


//via path params, http://localhost:8080/makanan/read?id=3
func GetMakananByIdParam(c echo.Context) error {
	id := c.QueryParam("id")
	makanan := new(models.Makanan)

	if err:= config.DB.First(&makanan, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"makanan dengan id tersebut tidak ada"})
	}

	return c.JSON(http.StatusOK, makanan)
}


func UpdateMakananById(c echo.Context) error {
	id := c.Param("id")
	nama:= c.FormValue("nama")
	hargaString:= c.FormValue("harga")

	harga, err := strconv.ParseUint(hargaString, 10, 0)
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]string{"error":"harga tidak dapat di konvesikan"})
	}

	makanan := new(models.Makanan)
	if err := config.DB.First(&makanan, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Makanan dengan id tersebut tidak ditemukan"})
    }

	makanan.Nama = nama
	makanan.Harga = uint(harga)

	if err := config.DB.Save(&makanan).Error; err!= nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"error":"Data tidak dapar tersimpan"})
	}

	return c.JSON(http.StatusOK, makanan)
}


func DeleteMakananById(c echo.Context) error {
	id := c.Param("id")

	makanan := new(models.Makanan)
	if err := config.DB.First(&makanan, id).Error; err!=nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data tidak ditemukan dengan id"})
	}

	if err := config.DB.Delete(&makanan).Error; err != nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data sudah terhapus"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message":"Data sudah dihapus"})
}