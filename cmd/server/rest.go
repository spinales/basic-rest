package server

import (
	"basic-rest/pkg/handlers"
	storage "basic-rest/pkg/storage/sqlite"
	"flag"
	"fmt"

	"github.com/labstack/echo/v4"
	driver "gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func main() {

	DB, err := gorm.Open(driver.Open("product.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var migrate string
	flag.StringVar(&migrate, "migrate", "no", "Genera la migracion de la base de datos.")
	flag.Parse()

	if migrate == "yes" {
		fmt.Println("Comenzo la Migracion...")
		// sqlite.Migrate(DB)
		fmt.Println("Termino la Migracion...")
	}

	service := &handlers.Service{
		ProductService: &storage.ProductService{DB},
	}

	e := echo.New()

	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	// 	AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	// }))

	e.GET("/product", service.GetProducts)
	e.GET("/equipo/:id", service.GetProduct)
	e.POST("/equipo", service.PostProduct)
	e.PUT("/equipo/:id", service.PutProduct)
	e.DELETE("/equipo/:id", service.DeleteProduct)

	e.Logger.Fatal(e.Start(":1323"))
}
