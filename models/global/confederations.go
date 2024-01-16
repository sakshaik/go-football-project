package global

import (
	"example.com/football-project/db"
	"example.com/football-project/queries"
)

type Confederation struct {
	ID          int64     `json:"confederation_id"`
	Name        string    `json:"confederation_name"`
	Description string    `json:"description"`
	Continent   Continent `json:"continent"`
}

func GetAllConfederations() ([]Confederation, error) {
	rows, err := db.DB.Query(queries.GET_ALL_CONFEDERATIONS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Confederation
	for rows.Next() {
		var confederation Confederation
		err := rows.Scan(&confederation.ID, &confederation.Name, &confederation.Description, &confederation.Continent.ID, &confederation.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, confederation)
	}
	return data, nil
}
