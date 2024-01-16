package global

import (
	"errors"

	"example.com/football-project/db"
	"example.com/football-project/queries"
)

type City struct {
	ID      int64   `json:"city_id"`
	Name    string  `json:"city_name"`
	Country Country `json:"country"`
}

func GetAllCities() ([]City, error) {
	rows, err := db.DB.Query(queries.GET_ALL_CITIES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []City
	for rows.Next() {
		var city City
		err := rows.Scan(&city.ID, &city.Name, &city.Country.ID, &city.Country.Name, &city.Country.Continent.ID, &city.Country.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, city)
	}
	return data, nil
}

func (c *City) AddCity() error {
	stmt, err := db.DB.Prepare(queries.ADD_CITY)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(c.Name, c.Country.ID)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}

func (c *City) ValidateCountry() error {
	switch {
	case c.Country.ID == 0 && c.Country.Name == "":
		return errors.New("Country details missing")
	case c.Country.ID != 0 && c.Country.Name != "":
		return c.Country.GetCountry()
	case c.Country.ID == 0 && c.Country.Name != "":
		country, err := GetCountryByName(c.Country.Name)
		if err != nil {
			return err
		} else if country.ID != 0 {
			c.Country.ID = country.ID
			return nil
		} else {
			return errors.New("Country details are not valid")
		}
	case c.Country.ID != 0 && c.Country.Name == "":
		country, err := GetCountryByID(c.Country.ID)
		if err != nil {
			return err
		} else if country.ID != 0 {
			c.Country.Name = country.Name
			return nil
		} else {
			return errors.New("Country details are not valid")
		}
	default:
		return nil
	}
}
