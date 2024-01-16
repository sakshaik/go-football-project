package leagues

import (
	"example.com/football-project/db"
	"example.com/football-project/models/global"
	"example.com/football-project/queries"
)

type League struct {
	ID            int64                `json:"league_id"`
	Name          string               `json:"league_name"`
	Country       global.Country       `json:"country"`
	Confederation global.Confederation `json:"confederation"`
}

func (l *League) AddLeague() error {
	stmt, err := db.DB.Prepare(queries.INSERT_LEAGUE)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(l.Name, l.Country.ID, l.Confederation.ID)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}

func GetLeaguesByCountry(id int64) ([]League, error) {
	rows, err := db.DB.Query(queries.FIND_LEAGUES_BY_COUNTRY_ID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []League
	for rows.Next() {
		var league League
		err := rows.Scan(&league.ID, &league.Name,
			&league.Country.ID, &league.Country.Name, &league.Country.Continent.ID, &league.Country.Continent.Name,
			&league.Confederation.ID, &league.Confederation.Name,
			&league.Confederation.Continent.ID, &league.Confederation.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, league)
	}
	return data, nil
}
