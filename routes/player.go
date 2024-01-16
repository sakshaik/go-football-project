package routes

import (
	"net/http"
	"strconv"

	"example.com/football-project/models/player"
	"github.com/gin-gonic/gin"
)

func addPlayer(context *gin.Context) {
	var player player.Player
	err := context.ShouldBindJSON(&player)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = player.AddPlayer()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding a new player."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Player added successfully."})
}

func findPlayerbyID(context *gin.Context) {
	playerId, err := strconv.ParseInt(context.Param("playerId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing player id."})
		return
	}
	p, err := player.FindPlayerByID(playerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting player info by ID"})
		return
	}
	context.JSON(http.StatusOK, p)
}
