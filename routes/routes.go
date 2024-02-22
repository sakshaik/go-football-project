package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routeProxy func(*gin.Engine)

func Register(server *gin.Engine) {
	registerAllRoutes(server,
		registerGlobalRoutes, registerClubRoutes,
		registerLeagueRoutes, registerPlayerRoutes,
		registerSearchRoutes, registerTransferRoutes)
}

func registerAllRoutes(server *gin.Engine, routes ...routeProxy) {
	for _, route := range routes {
		route(server)
	}
}

func registerGlobalRoutes(server *gin.Engine) {
	server.GET("/", defaultMessage)
	global := server.Group("/global")
	global.GET("/continents", getAllContinents)
	global.GET("/countries", getAllCountries)
	global.GET("/cities", getAllCities)
	global.GET("/confederations", getAllConfederations)
	global.POST("/city", addCity)
}

func registerPlayerRoutes(server *gin.Engine) {
	player := server.Group("/player")
	player.GET("/positions", getAllPositions)
	player.POST("", addPlayer)
	player.GET("", findPlayer)
	player.DELETE("/:playerId", deletePlayer)
}

func registerClubRoutes(server *gin.Engine) {
	club := server.Group("/club")
	club.POST("", addClub)
	club.POST("/player", addPlayersToClub)
}

func registerLeagueRoutes(server *gin.Engine) {
	league := server.Group("/league")
	league.POST("", addLeague)
}

func registerSearchRoutes(server *gin.Engine) {
	search := server.Group("/search")
	search.POST("/league", searchLeagues)
	search.POST("/club", searchClub)
}

func registerTransferRoutes(server *gin.Engine) {
	transfer := server.Group("/transfer")
	transfer.POST("/submit", submitTransferRequest)
	transfer.PUT("/reject", rejectTransferRequest)
	transfer.PUT("/approve", approveTransferRequest)
	transfer.GET("/player/:playerId", findPlayerTransferRequest)
}

func defaultMessage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello!"})
}
