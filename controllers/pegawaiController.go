package controllers

import (
	"go_restaurant/config"
	"go_restaurant/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePegawai(c echo.Context) error {
    nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_telp := c.FormValue("no_telp")
	username := c.FormValue("username")
	password := c.FormValue("password")

    pegawai := &models.Pegawai{
        Nama:  nama,
		Alamat: alamat,
		No_telp: no_telp,
		Username: username,
		Password: password,
    }

    if err := c.Bind(pegawai); err != nil {
        return err
    }

    if err := config.DB.Create(&pegawai).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal membuat daftar pegawai"})
    }

    return c.JSON(http.StatusCreated, pegawai)
}


func GetAllPegawai(c echo.Context) error {
	var pegawais []models.Pegawai

	if err := config.DB.Find(&pegawais).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"pegawais not found"})
	}

	return c.JSON(http.StatusOK, pegawais)
}


//via path url, http://localhost:8080/pegawai/read/3. untuk pencarian dengan id prefer yg ini
func GetPegawaiById(c echo.Context) error {
	id := c.Param("id")
	pegawai := new(models.Pegawai)

	if err:= config.DB.First(&pegawai, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"pegawai dengan id tersebut tidak ada"})
	}

	return c.JSON(http.StatusOK, pegawai)
}


//via path params, http://localhost:8080/pegawai/read?id=3
func GetPegawaiByIdParam(c echo.Context) error {
	id := c.QueryParam("id")
	pegawai := new(models.Pegawai)

	if err:= config.DB.First(&pegawai, id).Error; err != nil{
		return c.JSON(http.StatusNotFound, map[string]string{"error":"pegawai dengan id tersebut tidak ada"})
	}

	return c.JSON(http.StatusOK, pegawai)
}


func UpdatePegawaiById(c echo.Context) error {
	id := c.Param("id")
	nama:= c.FormValue("nama")
	alamat := c.FormValue("alamat")
	no_telp := c.FormValue("no_telp")
	username := c.FormValue("username")
	password := c.FormValue("password")

	pegawai := new(models.Pegawai)
	if err := config.DB.First(&pegawai, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "pegawai dengan id tersebut tidak ditemukan"})
    }

	pegawai.Nama = nama
	pegawai.Alamat = alamat
	pegawai.No_telp = no_telp
	pegawai.Username = username
	pegawai.Password = password

	if err := config.DB.Save(&pegawai).Error; err!= nil{
		return c.JSON(http.StatusInternalServerError,map[string]string{"error":"Data tidak dapar tersimpan"})
	}

	return c.JSON(http.StatusOK, pegawai)
}


func DeletePegawaiById(c echo.Context) error {
	id := c.Param("id")

	pegawai := new(models.Pegawai)
	if err := config.DB.First(&pegawai, id).Error; err!=nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data tidak ditemukan dengan id"})
	}

	if err := config.DB.Delete(&pegawai).Error; err != nil{
		return c.JSON(http.StatusNotFound,map[string]string{"error":"data sudah terhapus"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message":"Data sudah dihapus"})
}


func LoginPegawai(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	pegawai, err := models.FindUserByUsername(config.DB, username)
	if err != nil{
		return c.JSON(http.StatusUnauthorized, map[string]string{"error":"username tidak ditemukan"})
	}

	if err := pegawai.CheckPassword(password); err != nil{
		return c.JSON(http.StatusUnauthorized, map[string]string{"error":"password salah"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message":"login berhasil"})
}