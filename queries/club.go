package queries

const (
	ADD_CLUB = `
		INSERT INTO 
		club (club_name, league_id)
		VALUES (?,?)
	`
)
