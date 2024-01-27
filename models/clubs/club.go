package clubs

import (
	"bytes"

	"example.com/football-project/db"
	"example.com/football-project/models/base"
	"example.com/football-project/models/leagues"
	"example.com/football-project/models/player"
	"example.com/football-project/queries"
)

type Club struct {
	ID      int64           `json:"club_id"`
	Name    string          `json:"club_name"`
	League  leagues.League  `json:"league"`
	Players []player.Player `json:"players,omitempty"`
}

type PlayerExtRef struct {
	Club    Club            `json:"club"`
	Players []player.Player `json:"players"`
}

type PlayerExtRefData struct {
	Club   int64 `json:"club_id"`
	Player int64 `json:"player_id"`
}

type Search struct {
	Club           int64 `json:"club_id"`
	League         int64 `json:"league_id"`
	IncludePlayers bool  `json:"include_players"`
}

func (c *Club) AddClub() error {
	return base.InsertData(queries.ADD_CLUB, base.GenerateParamsInterface(c.Name, c.League.ID))
}

func (p *PlayerExtRef) AddPlayersToClub() error {
	for _, player := range p.Players {
		err := AddPlayerToClub(player.ID, p.Club.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func AddPlayerToClub(playerId, clubId int64) error {
	err := base.InsertData(queries.ADD_PLAYER_TO_CLUB, base.GenerateParamsInterface(playerId, clubId))
	if err != nil {
		return err
	}
	return nil
}

func RemovePlayerFromClub(playerId, clubId int64) error {
	err := base.UpdateOrDeleteData(queries.REMOVE_PLAYER_FROM_CLUB, base.GenerateParamsInterface(playerId, clubId))
	if err != nil {
		return err
	}
	return nil
}

func RemovePlayerReference(playerId int64) error {
	err := base.UpdateOrDeleteData(queries.REMOVE_PLAYER_REFERENCE, base.GenerateParamsInterface(playerId))
	if err != nil {
		return err
	}
	return nil
}

func (club *Club) GetClubDetails() error {
	var query bytes.Buffer
	query.WriteString(queries.FIND_CLUB_BASE_QUERY)
	query.WriteString(queries.WHERE)
	query.WriteString(queries.CLUB_ID_CLAUSE)
	row := db.DB.QueryRow(query.String(), club.ID)
	err := row.Scan(&club.ID, &club.Name, &club.League.ID, &club.League.Name,
		&club.League.Country.ID, &club.League.Country.Name,
		&club.League.Country.Continent.ID, &club.League.Country.Continent.Name,
		&club.League.Confederation.ID, &club.League.Confederation.Name, &club.League.Confederation.Description,
		&club.League.Confederation.Continent.ID, &club.League.Confederation.Continent.Name)
	if err != nil {
		return err
	}
	return nil
}

func (s *Search) GetClubDetails() ([]Club, error) {
	var query bytes.Buffer
	query.WriteString(queries.FIND_CLUB_BASE_QUERY)
	query.WriteString(queries.WHERE)
	var args []interface{}
	if s.Club != 0 {
		args = append(args, s.Club)
		query.WriteString(queries.CLUB_ID_CLAUSE)
	} else if s.League != 0 {
		args = append(args, s.League)
		query.WriteString(queries.LEAGUE_ID_CLAUSE)
	}
	rows, err := db.DB.Query(query.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Club
	for rows.Next() {
		var club Club
		err := rows.Scan(&club.ID, &club.Name, &club.League.ID, &club.League.Name,
			&club.League.Country.ID, &club.League.Country.Name,
			&club.League.Country.Continent.ID, &club.League.Country.Continent.Name,
			&club.League.Confederation.ID, &club.League.Confederation.Name, &club.League.Confederation.Description,
			&club.League.Confederation.Continent.ID, &club.League.Confederation.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, club)
	}
	return data, nil
}

func (s *Search) GetClubAndPlayerDetails() (*Club, error) {
	clubs, err := s.GetClubDetails()
	if err != nil {
		return nil, err
	}
	club := &clubs[0]
	err = club.GetPlayerExternalRef()
	if err != nil {
		return nil, err
	}
	for index, _ := range club.Players {
		err = club.Players[index].FindPlayerByID()
		if err != nil {
			return nil, err
		}
	}
	return club, nil
}

func (c *Club) GetPlayerExternalRef() error {
	rows, err := db.DB.Query(queries.FIND_PLAYER_REFERENCE_BY_CLUB_ID, c.ID)
	if err != nil {
		return err
	}
	defer rows.Close()
	var data []player.Player
	for rows.Next() {
		var player player.Player
		err = rows.Scan(&player.ID)
		if err != nil {
			return err
		}
		data = append(data, player)
	}
	c.Players = data
	return nil
}

func GetClubDetailsByPlayerId(playerId int64) (Club, error) {
	row := db.DB.QueryRow(queries.FIND_CLUB_BY_PLAYER_ID, playerId)
	var club Club
	err := row.Scan(&club.ID, &club.Name, &club.League.ID, &club.League.Name,
		&club.League.Country.ID, &club.League.Country.Name,
		&club.League.Country.Continent.ID, &club.League.Country.Continent.Name,
		&club.League.Confederation.ID, &club.League.Confederation.Name, &club.League.Confederation.Description,
		&club.League.Confederation.Continent.ID, &club.League.Confederation.Continent.Name)
	if err != nil {
		return Club{}, err
	}
	return club, err
}
