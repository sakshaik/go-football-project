package routes

import (
	"fmt"
	"net/http"

	"example.com/football-project/models/global"
	"github.com/gin-gonic/gin"
)

func getAllContinents(context *gin.Context) {
	continents, err := global.GetAllContinents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching Continents."})
		return
	}
	context.JSON(http.StatusOK, continents)
}

func getAllCountries(context *gin.Context) {
	countries, err := global.GetAllCountries()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching Countries."})
		return
	}
	context.JSON(http.StatusOK, countries)
}

func getAllCities(context *gin.Context) {
	cities, err := global.GetAllCities()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching Cities."})
		return
	}
	context.JSON(http.StatusOK, cities)
}

func getAllConfederations(context *gin.Context) {
	confederations, err := global.GetAllConfederations()
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error Fetching Confederation data."})
		return
	}
	context.JSON(http.StatusOK, confederations)
}

func addCity(context *gin.Context) {
	var city global.City
	err := context.ShouldBindJSON(&city)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = city.ValidateCountry()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = city.AddCity()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error adding a new city."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "City added successfully."})
}
