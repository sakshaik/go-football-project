package leagues

import (
	"bytes"
	"errors"

	"example.com/football-project/db"
	"example.com/football-project/models/base"
	"example.com/football-project/models/global"
	"example.com/football-project/queries"
)

type League struct {
	ID            int64                `json:"league_id"`
	Name          string               `json:"league_name"`
	Country       global.Country       `json:"country"`
	Confederation global.Confederation `json:"confederation"`
}

type Search struct {
	Country       int64 `json:"country_id"`
	Confederation int64 `json:"confederation_id"`
	League        int64 `json:"league_id"`
}

func (l *League) AddLeague() error {
	return base.InsertData(queries.ADD_LEAGUE, base.GenerateParamsInterface(l.Name, l.Country.ID, l.Confederation.ID))
}

func (search *Search) GetLeagues() ([]League, error) {
	var query bytes.Buffer
	query.WriteString(queries.FIND_LEAGUES_BASE_QUERY)
	var params []interface{}
	if *search != (Search{}) {
		query.WriteString(queries.WHERE)
		switch {
		case search.League != 0:
			query.WriteString(queries.LEAGUE_CLAUSE)
			params = append(params, search.League)
		case search.Country != 0 && search.Confederation != 0:
			query.WriteString(queries.COUNTRY_AND_CONFEDERATION_CLAUSE)
			params = append(params, search.Country, search.Confederation)
		case search.Country == 0 && search.Confederation != 0:
			query.WriteString(queries.CONFEDERATION_CLAUSE)
			params = append(params, search.Confederation)
		case search.Country != 0 && search.Confederation == 0:
			query.WriteString(queries.COUNTRY_CLAUSE)
			params = append(params, search.Country)
		default:
			return nil, errors.New("Search criteria is invalid")
		}
	} else {
		return nil, errors.New("Search criteria is invalid")
	}

	rows, err := db.DB.Query(query.String(), params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []League
	for rows.Next() {
		var league League
		err := rows.Scan(&league.ID, &league.Name,
			&league.Country.ID, &league.Country.Name, &league.Country.Continent.ID, &league.Country.Continent.Name,
			&league.Confederation.ID, &league.Confederation.Name, &league.Confederation.Description,
			&league.Confederation.Continent.ID, &league.Confederation.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, league)
	}
	if data == nil {
		return []League{}, nil
	}
	return data, nil
}
