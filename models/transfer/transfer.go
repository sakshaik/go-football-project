package transfer

import (
	"bytes"
	"errors"

	"example.com/football-project/constants"
	"example.com/football-project/db"
	"example.com/football-project/models/base"
	"example.com/football-project/models/clubs"
	"example.com/football-project/models/player"
	"example.com/football-project/queries"
)

type TransferRequest struct {
	ID       int64      `json:"transfer_id"`
	Status   string     `json:"status"`
	To       clubs.Club `json:"to_club"`
	Fee      float64    `json:"fee"`
	Currency string     `json:"currency"`
}

type Transfer struct {
	Player    player.Player     `json:"player"`
	From      clubs.Club        `json:"from_club"`
	Transfers []TransferRequest `json:"tranfer_requests"`
}

func (t *Transfer) SubmitTransferRequest() error {
	return base.InsertData(queries.SUBMIT_TRANSFER_REQUEST,
		base.GenerateParamsInterface(t.Player.ID, t.From.ID, t.Transfers[0].To.ID,
			t.Transfers[0].Fee, t.Transfers[0].Currency, constants.SUBMITTED))
}

func (t *TransferRequest) RejectTransferRequest() error {
	return base.UpdateOrDeleteData(queries.UPDATE_TRANSFER_REQUEST_STATUS,
		base.GenerateParamsInterface(constants.REJECTED, t.ID))
}

func (t *TransferRequest) ApproveTransferRequest() error {
	return base.UpdateOrDeleteData(queries.UPDATE_TRANSFER_REQUEST_STATUS,
		base.GenerateParamsInterface(constants.APPROVED, t.ID))
}

func FindAllTransferRequests() ([]TransferRequest, error) {
	query := `select transfer_id,to_club_id, fee, currency_code, status from transfer_player`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []TransferRequest
	for rows.Next() {
		var t TransferRequest
		err := rows.Scan(&t.ID, &t.To.ID, &t.Fee, &t.Currency, &t.Status)
		if err != nil {
			return nil, err
		}
		data = append(data, t)
	}
	return data, nil
}

func FindPlayerTransferRequests(playerId int64) (*Transfer, error) {
	var query bytes.Buffer
	query.WriteString(queries.PLAYER_TRANSFER_REQUEST_BASE_QUERY)
	query.WriteString(queries.WHERE)
	query.WriteString(queries.PLAYER_TRANSFER_PLAYER_ID_CLAUSE)
	var transfer Transfer
	var player player.Player
	player.ID = playerId
	err := player.FindPlayerByID()
	if err != nil {
		return nil, err
	}
	if player.ID == 0 || player.FirstName == "" {
		return nil, errors.New("player not found")
	}
	transfer.Player = player
	currentCLub, err := clubs.GetClubDetailsByPlayerId(player.ID)
	if err != nil {
		return nil, err
	}
	transfer.From = currentCLub
	rows, err := db.DB.Query(query.String(), playerId, transfer.From.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var transferRequests []TransferRequest
	for rows.Next() {
		var t TransferRequest
		err = rows.Scan(&t.ID, &t.To.ID,
			&t.Fee, &t.Currency, &t.Status)
		if err != nil {
			return nil, err
		}
		err = t.To.GetClubDetails()
		if err != nil {
			return nil, err
		}
		transferRequests = append(transferRequests, t)
	}
	transfer.Transfers = transferRequests
	return &transfer, nil
}
