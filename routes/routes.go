package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", defaultMessage)
	server.GET("/continents", getAllContinents)
	server.GET("/countries", getAllCountries)
	server.GET("/cities", getAllCities)
	server.GET("/confederations", getAllConfederations)
	server.POST("/city", addCity)
	server.POST("/player", addPlayer)
	server.GET("/player/:playerId", findPlayerbyID)
	server.POST("/league", addLeague)
	server.POST("/search/league", searchLeagues)
}

func defaultMessage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
