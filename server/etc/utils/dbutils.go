package utils

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ResetAndMigrateDB(db *sql.DB) error {
	// get files as *.sql
	path := GetMainPath() + "/infrastructure/store/mysql/schema/"
	files, _ := ioutil.ReadDir(path)
	if len(files) == 0 {
		log.Fatal("failed to read schema files")
	}
	for _, file := range files {
		// drop db i exist
		table := strings.Replace(file.Name(), ".sql", "", 1)
		_, err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", table))
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("dropped table: %s", table)
		}

		// create db
		b, _ := ioutil.ReadFile(path + file.Name())
		_sql := string(b)
		_, err = db.Exec(_sql)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("created table: %s", table)
		}
	}
	return nil
}
