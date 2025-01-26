package controller

import (
	"net/http"
	"ticketing-system/config"
	"ticketing-system/entity"
	"ticketing-system/response"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
    var events []entity.Event
    if err := config.DB.Find(&events).Error; err != nil {
        c.JSON(http.StatusInternalServerError, response.GeneralResponse{
            Message: "Failed to fetch events",
        })
        return
    }

    // Konversi ke response struct
    var eventList []response.EventListResponse
    for _, event := range events {
        eventList = append(eventList, response.EventListResponse{
            ID:       event.ID,
            Name:     event.Name,
            Capacity: event.Capacity,
            Price:    event.Price,
            Status:   event.Status,
        })
    }

    c.JSON(http.StatusOK, response.GeneralResponse{
        Message: "List of events",
        Data:    eventList,
    })
}

func GetEventByID(c *gin.Context) {
    var event entity.Event
    id := c.Param("id")

    // Cari event berdasarkan ID
    if err := config.DB.First(&event, id).Error; err != nil {
        c.JSON(http.StatusNotFound, response.GeneralResponse{
            Message: "Event not found",
        })
        return
    }

    // Respons sukses
    c.JSON(http.StatusOK, response.GeneralResponse{
        Message: "Event fetched successfully",
        Data: response.EventResponse{
            ID:        event.ID,
            Name:      event.Name,
            Capacity:  event.Capacity,
            Price:     event.Price,
            Status:    event.Status,
            EventDate: event.EventDate,
            CreatedAt: event.CreatedAt,
            UpdatedAt: event.UpdatedAt,
        },
    })
}


func CreateEvent(c *gin.Context) {
    var input struct {
        Name     string  `json:"name" binding:"required"`
        Capacity int     `json:"capacity" binding:"required,gte=0"`
        Price    float64 `json:"price" binding:"required,gte=0"`
    }

    // Validasi input JSON
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Validasi apakah nama event sudah ada di database
    var existingEvent entity.Event
    if err := config.DB.Where("name = ?", input.Name).First(&existingEvent).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Event name already exists"})
        return
    }

    // Jika nama event tidak ada, buat event baru
    event := entity.Event{
        Name:     input.Name,
        Capacity: input.Capacity,
        Price:    input.Price,
    }

    if err := config.DB.Create(&event).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
        return
    }

     // Buat respons sukses
    eventResponse := response.EventResponse{
        ID:        event.ID,
        Name:      event.Name,
        Capacity:  event.Capacity,
        Price:     event.Price,
        Status:    event.Status,
        EventDate: event.EventDate,
        CreatedAt: event.CreatedAt,
        UpdatedAt: event.UpdatedAt,
    }
    c.JSON(http.StatusCreated, response.GeneralResponse{
        Message: "Event successfully created",
        Data:    eventResponse,
    })
}

// PUT /events/:id - Mengupdate event berdasarkan ID
func UpdateEvent(c *gin.Context) {
    var input struct {
        Name     string  `json:"name" binding:"required"`
        Capacity int     `json:"capacity" binding:"required,gte=0"`
        Price    float64 `json:"price" binding:"required,gte=0"`
    }
    id := c.Param("id")

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, response.GeneralResponse{
            Message: err.Error(),
        })
        return
    }

    var event entity.Event
    if err := config.DB.First(&event, id).Error; err != nil {
        c.JSON(http.StatusNotFound, response.GeneralResponse{
            Message: "Event not found",
        })
        return
    }

    config.DB.Model(&event).Updates(input)

    c.JSON(http.StatusOK, response.GeneralResponse{
        Message: "Event updated successfully",
        Data: response.EventResponse{
            ID:        event.ID,
            Name:      event.Name,
            Capacity:  event.Capacity,
            Price:     event.Price,
            Status:    event.Status,
            EventDate: event.EventDate,
            CreatedAt: event.CreatedAt,
            UpdatedAt: event.UpdatedAt,
        },
    })
}

func DeleteEvent(c *gin.Context) {
    var event entity.Event
    id := c.Param("id")

    if err := config.DB.First(&event, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
        return
    }

    if err := config.DB.Delete(&event).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
