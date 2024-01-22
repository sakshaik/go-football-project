package queries

const (
	SUBMIT_TRANSFER_REQUEST = `
		INSERT INTO 
		transfer_player (player_id, from_club_id, to_club_id, fee, currency_code, status)
		VALUES (?,?, ?, ?, ?, ?)
	`
	UPDATE_TRANSFER_REQUEST_STATUS = `
		UPDATE transfer_player 
		SET status = ?
		where transfer_id = ?
	`
	PLAYER_TRANSFER_REQUEST_BASE_QUERY = `
		select tp.transfer_id, tp.to_club_id, tp.fee, tp.currency_code, tp.status
			from transfer_player tp
	`
	PLAYER_TRANSFER_PLAYER_ID_CLAUSE = `
		tp.player_id = ? and tp.from_club_id = ?
	`
)
