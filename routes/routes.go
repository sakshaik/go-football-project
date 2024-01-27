package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", defaultMessage)
	global := server.Group("/global")
	global.GET("/continents", getAllContinents)
	global.GET("/countries", getAllCountries)
	global.GET("/cities", getAllCities)
	global.GET("/confederations", getAllConfederations)
	global.POST("/city", addCity)

	server.GET("/player/positions", getAllPositions)
	server.POST("/player", addPlayer)
	server.GET("/player", findPlayer)
	server.DELETE("/player/:playerId", deletePlayer)

	server.POST("/league", addLeague)
	server.POST("/search/league", searchLeagues)
	server.POST("/club", addClub)
	server.POST("/club/player", addPlayersToClub)
	server.POST("/search/club", searchClub)

	server.POST("/transfer/submit", submitTransferRequest)
	server.PUT("/transfer/reject", rejectTransferRequest)
	server.PUT("/transfer/approve", approveTransferRequest)
	server.GET("/transfer/player/:playerId", findPlayerTransferRequest)
}

func defaultMessage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
