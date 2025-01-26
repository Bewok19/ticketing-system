package main

import (
	"log"
	"ticketing-system/config"
	"ticketing-system/controller"
	"ticketing-system/entity"
	"ticketing-system/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
    config.ConnectDB()

	// Migrasi tabel
    err := config.DB.AutoMigrate(&entity.User{}, &entity.Event{}, &entity.Ticket{})
    if err != nil {
        log.Fatal("Failed to migrate database: ", err)
    }

	r := gin.Default()

	//Register Routes
	r.POST("/register", controller.RegisterUser)
	//Login Routes
	r.POST("/login", controller.LoginUser)

	//Event Routes
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AdminOnly()) // Hanya admin yang dapat mengakses

	adminRoutes.GET("/events", controller.GetEvents)
	adminRoutes.GET("/events/:id", controller.GetEventByID)
	adminRoutes.POST("/events", controller.CreateEvent)
	adminRoutes.PUT("/events/:id", controller.UpdateEvent)
	adminRoutes.DELETE("/events/:id", controller.DeleteEvent)

	//Routes untuk user melihat daftar event
	r.GET("/events", controller.GetEvents) 
	r.GET("/events/:id", controller.GetEventByID)

	//Ticket Routes
	ticket := r.Group("/tickets")
	ticket.Use(middleware.AuthMiddleware())
{
    ticket.POST("/", controller.CreateTicket)
    ticket.GET("/", controller.GetTickets)
    ticket.DELETE("/:id", controller.DeleteTicket)
}

	
	// Report Routes
	reportRoutes := r.Group("/reports")
	reportRoutes.Use(middleware.AdminOnly())
{
	reportRoutes.GET("/summary", controller.GetSummaryReport)
	reportRoutes.GET("/event/:id", controller.GetEventReport)
}

    log.Println("Starting the application...")
	r.Run(":8080")
}