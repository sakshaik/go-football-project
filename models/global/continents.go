package global

import (
	"example.com/football-project/db"
	"example.com/football-project/queries"
)

type Continent struct {
	ID   int64  `json:"continent_id"`
	Name string `json:"continent_name"`
}

func GetAllContinents() ([]Continent, error) {
	rows, err := db.DB.Query(queries.GET_ALL_CONTINENTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Continent
	for rows.Next() {
		var continent Continent
		err := rows.Scan(&continent.ID, &continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, continent)
	}
	return data, nil
}
