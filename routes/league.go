package routes

import (
	"net/http"

	"example.com/football-project/models/leagues"
	"github.com/gin-gonic/gin"
)

func addLeague(context *gin.Context) {
	var league leagues.League
	err := context.ShouldBindJSON(&league)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = league.AddLeague()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding a new league."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "League added successfully."})
}

func searchLeagues(context *gin.Context) {
	var search leagues.Search
	err := context.ShouldBindJSON(&search)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	data, err := search.GetLeagues()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, data)
}
