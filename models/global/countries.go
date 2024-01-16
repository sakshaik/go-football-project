package global

import (
	"database/sql"
	"errors"

	"example.com/football-project/db"
	"example.com/football-project/queries"
)

type Country struct {
	ID        int64     `json:"country_id"`
	Name      string    `json:"country_name"`
	Continent Continent `json:"continent,omitempty"`
}

func GetAllCountries() ([]Country, error) {
	rows, err := db.DB.Query(queries.GET_ALL_COUNTRIES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var data []Country
	for rows.Next() {
		var country Country
		err := rows.Scan(&country.ID, &country.Name, &country.Continent.ID, &country.Continent.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, country)
	}
	return data, nil
}

func (country *Country) GetCountry() error {
	row := db.DB.QueryRow(queries.GET_COUNTRY, country.ID, country.Name)
	var value Country
	err := row.Scan(&value.ID, &value.Name, &value.Continent.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Country details are not valid")
		} else {
			return err
		}
	}
	return nil
}

func GetCountryByID(id int64) (*Country, error) {
	row := db.DB.QueryRow(queries.GET_COUNTRY_BY_ID, id)
	var value Country
	err := row.Scan(&value.ID, &value.Name, &value.Continent.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Country ID is not valid")
		} else {
			return nil, err
		}
	}
	return &value, nil
}

func GetCountryByName(name string) (*Country, error) {
	row := db.DB.QueryRow(queries.GET_COUNTRY_BY_NAME, name)
	var value Country
	err := row.Scan(&value.ID, &value.Name, &value.Continent.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Country name is not valid")
		} else {
			return nil, err
		}
	}
	return &value, nil
}
