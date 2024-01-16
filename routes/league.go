package routes

import (
	"fmt"
	"net/http"
	"strconv"

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
	params := context.Request.URL.Query()
	for key, value := range params {
		if key == "countryId" {
			countryId, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse country id parameter"})
				return
			}
			data, err := leagues.GetLeaguesByCountry(countryId)
			if err != nil {
				fmt.Println(err)
				context.JSON(http.StatusInternalServerError, gin.H{"message": "Error getting league data by country ID"})
				return
			}
			context.JSON(http.StatusOK, data)
			break
		}
	}
}
