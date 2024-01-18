package clubs

import (
	"example.com/football-project/db"
	"example.com/football-project/models/leagues"
	"example.com/football-project/models/player"
	"example.com/football-project/queries"
)

type Club struct {
	ID      int64           `json:"club_id"`
	Name    string          `json:"club_name"`
	League  leagues.League  `json:"league"`
	Players []player.Player `json:"players"`
}

func (c *Club) AddClub() error {
	query := queries.INSERT_CLUB
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(c.Name, c.League.ID)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}
