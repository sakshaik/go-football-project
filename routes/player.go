package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/football-project/models/clubs"
	"example.com/football-project/models/player"
	"github.com/gin-gonic/gin"
)

func getAllPositions(context *gin.Context) {
	positions, err := player.GetAllPositions()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching Positions."})
		return
	}
	context.JSON(http.StatusOK, positions)
}

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

func findPlayer(context *gin.Context) {
	var player player.Player
	value := context.Query("player-id")
	player.FirstName = context.Query("first-name")
	player.LastName = context.Query("last-name")
	if value != "" {
		playerId, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing player id."})
			return
		}
		player.ID = playerId
		err = player.FindPlayerByID()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting player info by ID", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, player)
	} else if player.FirstName != "" || player.LastName != "" {
		players, err := player.FindPlayersByName()
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				gin.H{"message": "Error getting player info by name", "error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, players)
	}

}

func deletePlayer(context *gin.Context) {
	playerId, err := strconv.ParseInt(context.Param("playerId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing player id."})
		return
	}
	var player player.Player
	player.ID = playerId
	err = player.FindPlayerByID()
	if err != nil || (player.LastName == "" && player.FirstName == "") {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error finding player details."})
		return
	}
	err = clubs.RemovePlayerReference(player.ID)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error deleting player reference."})
		return
	}
	err = player.DeletePlayer()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error deleting player profile."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Player deleted successfully."})
}
