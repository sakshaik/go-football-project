package routes

import (
	"net/http"

	"example.com/football-project/models/clubs"
	"github.com/gin-gonic/gin"
)

func addClub(context *gin.Context) {
	var club clubs.Club
	err := context.ShouldBindJSON(&club)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = club.AddClub()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding a new club."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Club added successfully."})
}

func addPlayersToClub(context *gin.Context) {
	var ref clubs.PlayerExtRef
	err := context.ShouldBindJSON(&ref)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = ref.AddPlayerToClub()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding a new player to club"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Players added to Club successfully."})
}

func searchClub(context *gin.Context) {
	var search clubs.Search
	err := context.ShouldBindJSON(&search)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	var data any
	if search.IncludePlayers && search.Club != 0 {
		data, err = search.GetClubAndPlayerDetails()
	} else {
		data, err = search.GetClubDetails()
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, data)
}
