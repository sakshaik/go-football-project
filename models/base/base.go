package base

import "example.com/football-project/db"

func InsertData(query string, args []interface{}) error {
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}

func UpdateOrDeleteData(query string, args []interface{}) error {
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil
}

func GenerateParamsInterface(args ...any) []interface{} {
	if args == nil {
		return nil
	}
	params := make([]interface{}, len(args))
	copy(params, args)
	return params
}
