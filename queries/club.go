package queries

const (
	INSERT_CLUB = `
		INSERT INTO 
		club (club_name, league_id)
		VALUES (?,?)
	`
)
