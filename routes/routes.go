package routes

import (
	"go_restaurant/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.POST("/makanan/add",controllers.CreateMakanan)
	e.GET("/makanan/read",controllers.GetAllMakanan)
	e.GET("/makanan/read/:id",controllers.GetMakananById)
	e.PUT("/makanan/update/:id",controllers.UpdateMakananById)
	e.DELETE("/makanan/delete/:id", controllers.DeleteMakananById)
	//e.GET("/makanan/read",controllers.GetMakananByIdParam)


	e.POST("/minuman/add",controllers.CreateMinuman)
	e.GET("/minuman/read",controllers.GetAllMinuman)
	e.GET("/minuman/read/:id",controllers.GetMinumanById)
	e.PUT("/minuman/update/:id",controllers.UpdateMinumanById)
	e.DELETE("/minuman/delete/:id", controllers.DeleteMinumanById)
	//e.GET("/minuman/read",controllers.GetMinumanByIdParam)


	e.POST("/pelanggan/add",controllers.CreatePelanggan)
	e.GET("/pelanggan/read",controllers.GetAllPelanggan)
	e.GET("/pelanggan/read/:id",controllers.GetPelangganById)
	e.PUT("/pelanggan/update/:id",controllers.UpdatePelangganById)
	e.DELETE("/pelanggan/delete/:id", controllers.DeletePelangganById)
	//e.GET("/pelanggan/read",controllers.GetpelangganByIdParam)


	e.POST("/pegawai/add",controllers.CreatePegawai)
	e.GET("/pegawai/read",controllers.GetAllPegawai)
	e.GET("/pegawai/read/:id",controllers.GetPegawaiById)
	e.PUT("/pegawai/update/:id",controllers.UpdatePegawaiById)
	e.DELETE("/pegawai/delete/:id", controllers.DeletePegawaiById)
	//e.GET("/pegawai/read",controllers.GetpegawaiByIdParam)


	e.POST("/transaksi/add",controllers.CreateTransaksi)

	return e
}