package routes

import (
	"net/http"
	"strconv"

	"example.com/football-project/models/transfer"
	"github.com/gin-gonic/gin"
)

func submitTransferRequest(context *gin.Context) {
	var transfer transfer.Transfer
	err := context.ShouldBindJSON(&transfer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = transfer.From.GetClubDetails()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to search current club"})
		return
	}
	err = transfer.Transfers[0].To.GetClubDetails()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to search destination club"})
		return
	}
	err = transfer.SubmitTransferRequest()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error submitting a transfer request"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Transfer Request submitted successfully."})
}

func findPlayerTransferRequest(context *gin.Context) {
	playerId, err := strconv.ParseInt(context.Param("playerId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing player id."})
		return
	}
	data, err := transfer.FindPlayerTransferRequests(playerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error retrieving transfer requests for player", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, data)
}

func findAllTransfers(context *gin.Context) {
	data, err := transfer.FindAllTransferRequests()
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error retrieving transfer requests", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, data)
}
