package routes

import (
	"net/http"
	"strconv"

	"example.com/football-project/models/clubs"
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

func rejectTransferRequest(context *gin.Context) {
	var transfer transfer.Transfer
	err := context.ShouldBindJSON(&transfer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	err = transfer.Transfers[0].RejectTransferRequest()
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error rejecting transfer request", "error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Transfer Request rejected successfully."})
}

func approveTransferRequest(context *gin.Context) {
	var request transfer.Transfer
	err := context.ShouldBindJSON(&request)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse sent request."})
		return
	}
	currentRequest, err := transfer.FindPlayerTransferRequests(request.Player.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error finding current transfer requests"})
		return
	}
	currentClub := currentRequest.From
	var newClub clubs.Club
	var found bool
	for _, transferRequest := range currentRequest.Transfers {
		if transferRequest.ID == request.Transfers[0].ID {
			err = transferRequest.ApproveTransferRequest()
			newClub = transferRequest.To
			found = true
		} else {
			err = transferRequest.RejectTransferRequest()
		}
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				gin.H{"message": "Error approving or rejecting transfer request", "error": err.Error()})
			return
		}
	}
	if !found {
		context.JSON(http.StatusBadRequest,
			gin.H{"message": "Requested transfer id not found for player"})
		return
	}
	err = clubs.RemovePlayerFromClub(request.Player.ID, currentClub.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error removing player from club"})
		return
	}
	err = clubs.AddPlayerToClub(request.Player.ID, newClub.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			gin.H{"message": "Error adding player to new club"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Player transferred successfully."})
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
