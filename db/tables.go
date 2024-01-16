package db

import (
	"example.com/football-project/filemanager"
)

func CreateTables() {
	tables, err := filemanager.ReadFileLines("db/scripts/table_list.txt")
	if err != nil {
		panic("Unable to read table file : " + err.Error())
	}
	for _, table := range tables {
		scriptFile := "db/scripts/tables/" + table + ".sql"
		script, err := filemanager.ReadFileAsString(scriptFile)
		if err != nil {
			panic("Unable to read table file : " + scriptFile + " : " + err.Error())
		}
		_, err = DB.Exec(script)
		if err != nil {
			panic("Unable to create table : " + scriptFile + " : " + err.Error())
		}
	}
}

func InsertMasterData() {
	tables, err := filemanager.ReadFileLines("db/scripts/insert_sequence.txt")
	if err != nil {
		panic("Unable to read table file : " + err.Error())
	}
	for _, table := range tables {
		scriptFile := "db/scripts/tables/insert/" + table + ".sql"
		scripts, err := filemanager.ReadFileLines(scriptFile)
		if err != nil {
			panic("Unable to read insert file : " + scriptFile + " : " + err.Error())
		}
		for _, script := range scripts {
			_, err = DB.Exec(script)
			if err != nil {
				panic("Unable to insert data in table : " + script + " : " + err.Error())
			}
		}
	}
}
