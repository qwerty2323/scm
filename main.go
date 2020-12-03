package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/qwerty2323/scm/order"
	"github.com/qwerty2323/scm/product"
	"github.com/qwerty2323/scm/supplier"
	"gorm.io/gorm"
	_ "gorm.io/driver/postgres"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api/")

	v1 := api.Group("/v1")

	v1.Get("/orders", order.GetAll)
	v1.Post("/order", order.Create)
	v1.Delete("/order/:id", order.Delete)
	v1.Put("/order/:id", order.Update)

	v1.Get("/products", product.GetAll)
	v1.Post("/product", product.Create)
	v1.Delete("/product/:id", product.Delete)
	v1.Put("/product/:id", product.Update)

	v1.Get("/suppliers", supplier.GetAll)
	v1.Post("/supplier", supplier.Create)
	v1.Delete("/supplier/:id", supplier.Delete)
	v1.Put("/supplier/:id", supplier.Update)
}

func initializeDatabase() {
	var err error

	database.DBConn, err = gorm.Open("postgres", "scm.db")

	if err != nil {
		fmt.Printf(err.Error())
	}
}

func main() {
	app := fiber.New()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(":3000")
}
