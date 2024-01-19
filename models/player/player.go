package player

import (
	"bytes"

	"example.com/football-project/db"
	"example.com/football-project/models/base"
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
	return base.InsertData(queries.ADD_PLAYER, base.GenerateParamsInterface(p.FirstName, p.LastName, p.Age, p.City.ID))
}

func FindPlayerByID(id int64) (*Player, error) {
	var query bytes.Buffer
	query.WriteString(queries.FIND_PLAYER_BASE_QUERY)
	query.WriteString(queries.WHERE)
	query.WriteString(queries.PLAYER_ID_CLAUSE)
	row := db.DB.QueryRow(query.String(), id)
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

func FindPlayersByID(ids []int64) ([]Player, error) {
	var query bytes.Buffer
	query.WriteString(queries.FIND_PLAYER_BASE_QUERY)
	query.WriteString(queries.WHERE)
	query.WriteString(queries.PLAYER_IDS_CLAUSE)
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	rows, err := db.DB.Query(query.String(), args...)
	if err != nil {
		return nil, err
	}
	var data []Player
	for rows.Next() {
		var player Player
		err = rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age,
			&player.City.ID, &player.City.Name,
			&player.City.Country.ID, &player.City.Country.Name,
			&player.City.Country.Continent.ID, &player.City.Country.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, player)
	}
	return data, nil
}
