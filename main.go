package main

import (
	"net/http"
	"strconv"

	"github.com/basharatoum/weatherservice/weatherservice"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/weather/:lat/:long", getWeather)

	r.Run(":8080")
}

func getWeather(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Param("lat"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude parameter"})
		return
	}

	long, err := strconv.ParseFloat(c.Param("long"), 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude parameter"})
		return
	}

	weather, err := weatherservice.GetWeather(long, lat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, weather)
}
