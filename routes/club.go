package routes

import (
	"fmt"
	"net/http"

	"example.com/football-project/models/clubs"
	"github.com/gin-gonic/gin"
)

func addClub(context *gin.Context) {
	var club clubs.Club
	err := context.ShouldBindJSON(&club)
	if err != nil {
		fmt.Println(err)
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
