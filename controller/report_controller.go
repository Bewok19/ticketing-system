package controller

import (
	"net/http"
	"strconv"
	"ticketing-system/service"

	"github.com/gin-gonic/gin"
)

func GetSummaryReport(c *gin.Context) {
    report, err := service.GetSummaryReport()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": report})
}

func GetEventReport(c *gin.Context) {
    eventIDParam := c.Param("id")
    eventID, err := strconv.ParseUint(eventIDParam, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
        return
    }

    report, err := service.GetEventReport(uint(eventID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": report})
}
