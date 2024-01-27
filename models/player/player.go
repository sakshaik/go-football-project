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
	Position  Position    `json:"position"`
	City      global.City `json:"city"`
}

type Position struct {
	ID        int64  `json:"position_id"`
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name"`
}

func (p *Player) AddPlayer() error {
	return base.InsertData(queries.ADD_PLAYER,
		base.GenerateParamsInterface(p.FirstName, p.LastName, p.Age, p.City.ID, p.Position.ID))
}

func (p *Player) DeletePlayer() error {
	return base.UpdateOrDeleteData(queries.REMOVE_PLAYER, base.GenerateParamsInterface(p.ID))
}

func GetAllPositions() ([]Position, error) {
	rows, err := db.DB.Query(queries.FIND_ALL_POSITIONS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Position
	for rows.Next() {
		var p Position
		err := rows.Scan(&p.ID, &p.ShortName, &p.FullName)
		if err != nil {
			return nil, err
		}
		data = append(data, p)
	}
	return data, nil
}

func (p *Player) FindPlayersByName() ([]Player, error) {
	var query bytes.Buffer
	query.WriteString(queries.FIND_PLAYER_BASE_QUERY)
	query.WriteString(queries.WHERE)
	var args []interface{}
	switch {
	case p.FirstName != "" && p.LastName != "":
		query.WriteString(queries.PLAYER_FIRST_NAME_CLAUSE)
		firstNameQuery := "%" + p.FirstName + "%"
		query.WriteString(queries.AND)
		query.WriteString(queries.PLAYER_LAST_NAME_CLAUSE)
		lastNameQuery := "%" + p.LastName + "%"
		args = base.GenerateParamsInterface(firstNameQuery, lastNameQuery)
	case p.FirstName != "" && p.LastName == "":
		firstNameQuery := "%" + p.FirstName + "%"
		query.WriteString(queries.PLAYER_FIRST_NAME_CLAUSE)
		args = base.GenerateParamsInterface(firstNameQuery)
	case p.FirstName == "" && p.LastName != "":
		query.WriteString(queries.PLAYER_LAST_NAME_CLAUSE)
		lastNameQuery := "%" + p.LastName + "%"
		args = base.GenerateParamsInterface(lastNameQuery)
	}
	rows, err := db.DB.Query(query.String(), args...)
	if err != nil {
		return nil, err
	}
	var data []Player
	for rows.Next() {
		var player Player
		err = rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age,
			&player.Position.ID, &player.Position.ShortName, &player.Position.FullName,
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

func (p *Player) FindPlayerByID() error {
	var query bytes.Buffer
	query.WriteString(queries.FIND_PLAYER_BASE_QUERY)
	query.WriteString(queries.WHERE)
	query.WriteString(queries.PLAYER_ID_CLAUSE)
	row := db.DB.QueryRow(query.String(), p.ID)
	err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Age,
		&p.Position.ID, &p.Position.ShortName, &p.Position.FullName,
		&p.City.ID, &p.City.Name,
		&p.City.Country.ID, &p.City.Country.Name,
		&p.City.Country.Continent.ID, &p.City.Country.Continent.Name)
	if err != nil {
		return err
	}
	return nil
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
	defer rows.Close()
	var data []Player
	for rows.Next() {
		var player Player
		err = rows.Scan(&player.ID, &player.FirstName, &player.LastName, &player.Age,
			&player.Position.ID, &player.Position.ShortName, &player.Position.FullName,
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
