package controllers

import (
	"go_restaurant/config"
	"go_restaurant/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateTransaksi(c echo.Context) error {
    // Ambil data dari form
    jumlahMakananString := c.FormValue("jumlahMakanan")
    jumlahMinumanString := c.FormValue("jumlahMinuman")
    pegawaiID := c.FormValue("pegawaiID")
    makananID := c.FormValue("makananID")
    minumanID := c.FormValue("minumanID")

    // Konversi jumlahMakanan dan jumlahMinuman
    jumlahMakanan, err := strconv.ParseUint(jumlahMakananString, 10, 0)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tidak dapat konversi jumlahMakanan"})
    }

    jumlahMinuman, err := strconv.ParseUint(jumlahMinumanString, 10, 0)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"error": "Tidak dapat konversi jumlahMinuman"})
    }

    // Membuat variabel untuk pegawai, makanan, dan minuman
    var pegawai models.Pegawai
    var makanan models.Makanan
    var minuman models.Minuman

    // Mengecek apakah pegawai dengan ID tersebut ada
    if err := config.DB.First(&pegawai, pegawaiID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Pegawai dengan ID tersebut tidak ditemukan"})
    }

    // Mengecek apakah makanan dengan ID tersebut ada
    if err := config.DB.First(&makanan, makananID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Makanan dengan ID tersebut tidak ditemukan"})
    }

    // Mengecek apakah minuman dengan ID tersebut ada
    if err := config.DB.First(&minuman, minumanID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"error": "Minuman dengan ID tersebut tidak ditemukan"})
    }

    // Menghitung total harga
    totalHarga := calculateTotalHarga(makanan.Harga, minuman.Harga, uint(jumlahMakanan), uint(jumlahMinuman))

    // Membuat transaksi
    transaksi := &models.Transaksi{
        Pegawai:     pegawai,
        Makanan:     makanan,
        Minuman:     minuman,
        JumlahMakanan: uint(jumlahMakanan),
        JumlahMinuman: uint(jumlahMinuman),
        TotalHarga:    totalHarga,
    }

    // Menyimpan transaksi ke database
    if err := config.DB.Create(&transaksi).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal menyimpan transaksi ke database"})
    }
    // Mengembalikan respons berhasil
    return c.JSON(http.StatusCreated, transaksi)
}



func calculateTotalHarga(hargaMakanan, hargaMinuman, jumlahMakanan uint, jumlahMinuman uint) uint {
	totalHargaMakanan := hargaMakanan * jumlahMakanan
	totalHargaMinuman := hargaMinuman * jumlahMinuman

	return totalHargaMakanan + totalHargaMinuman
}


func GetAllTransaksi(c echo.Context) error {
	var transaksis []models.Transaksi

	if err := config.DB.Find(&transaksis).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transaksis not found"})
	}

	return c.JSON(http.StatusOK, transaksis)
}


func GetTransaksiById(c echo.Context) error {
	id := c.Param("id")
	transaksi := new(models.Pegawai)

	if err := config.DB.First(&transaksi, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transaksi dengan id itu not found"})
	}

	return c.JSON(http.StatusOK, transaksi)
}


func UpdateTransaksiById(c echo.Context) error {
	id := c.Param("id")
	

	transaksi := new(models.Transaksi)
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "transaksi dengan id tersebut tidak ditemukan"})
	}

	

	if err := config.DB.Save(&transaksi).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Data tidak dapar tersimpan"})
	}

	return c.JSON(http.StatusOK, transaksi)
}


func DeleteTransaksiById(c echo.Context) error {
	id := c.Param("id")

	transaksi := new(models.Transaksi)
	if err := config.DB.First(&transaksi, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "data tidak ditemukan dengan id"})
	}

	if err := config.DB.Delete(&transaksi).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "data sudah terhapus"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Data sudah dihapus"})
}