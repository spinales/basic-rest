package main

import (
	"basic-rest/pkg/handlers"
	storage "basic-rest/pkg/storage/sqlite"
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	driver "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func main() {

	DB, err := gorm.Open(driver.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	migrate := flag.Bool("migrate", false, "Genera la migracion de la base de datos.")
	flag.Parse()

	if *migrate {
		fmt.Println("Comenzo la Migracion...")
		storage.Migrate(DB)
		fmt.Println("Termino la Migracion...")
	}

	service := &handlers.Service{
		ProductService: &storage.ProductService{DB},
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/product", service.GetProducts)
	e.GET("/product/:id", service.GetProduct)
	e.POST("/product", service.PostProduct)
	e.PUT("/product/:id", service.PutProduct)
	e.DELETE("/product/:id", service.DeleteProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
