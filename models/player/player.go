package player

import (
	"example.com/football-project/db"
	"example.com/football-project/models/global"
	"example.com/football-project/queries"
)

type Player struct {
	ID        int64       `json:"player_id"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Age       int64       `json:"age"`
	City      global.City `json:"city"`
}

func (p *Player) AddPlayer() error {
	query := queries.INSERT_PLAYER
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(p.FirstName, p.LastName, p.Age, p.City.ID)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}

func FindPlayerByID(id int64) (*Player, error) {
	query := queries.FIND_PLAYER_BY_ID
	row := db.DB.QueryRow(query, id)
	var player Player
	err := row.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age,
		&player.City.ID, &player.City.Name,
		&player.City.Country.ID, &player.City.Country.Name,
		&player.City.Country.Continent.ID, &player.City.Country.Continent.Name)
	if err != nil {
		return nil, err
	}
	return &player, nil
}
