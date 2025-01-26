package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"ticketing-system/response"
	"ticketing-system/service"

	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	// Log to verify userID type
	fmt.Printf("User ID from context: %v, Type: %T\n", userID, userID)

	// Ensure userID is uint
	uid, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID is not valid"})
		return
	}

	// Parse request body
	var req struct {
		EventID  uint `json:"event_id"`
		Quantity int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Call service to create ticket
	result, err := service.CreateTicket(uid, req.EventID, req.Quantity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket successfully created",
		"data":    result,
	})
}

func GetTickets(c *gin.Context) {
    userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userRole, exists := c.Get("userRole")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Fetch tickets based on role
	tickets, err := service.GetTickets(userID.(uint), userRole.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Transform tickets to response format
	ticketResponses := response.ToTicketResponses(tickets)

	c.JSON(http.StatusOK, gin.H{"data": ticketResponses})
}

func DeleteTicket(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	isAdmin := c.MustGet("userRole").(string) == "Admin"
	ticketID := c.Param("id")

	// Convert ticketID from string to uint
	tid, err := strconv.ParseUint(ticketID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	if err := service.DeleteTicket(uint(tid), userID, isAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ticket successfully canceled"})
}