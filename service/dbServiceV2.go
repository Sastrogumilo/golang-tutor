package services

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

func MainQuery(query string) (result interface{}, error error) {

	defer func() {
		if r := recover(); r != nil {
			error = fmt.Errorf("panic recovered: %v", r)
			result = nil
			log.Println(error)
		}
	}()

	dbMainHost := os.Getenv("DB_MAIN_HOST")
	dbMainPort := os.Getenv("DB_MAIN_PORT")
	dbMainUser := os.Getenv("DB_MAIN_USER")
	dbMainPass := os.Getenv("DB_MAIN_PASS")

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbMainUser, dbMainPass, dbMainHost, dbMainPort))
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	isSelectOrCall, _ := regexp.MatchString("(?i)^(select|call)", strings.TrimSpace(query))
	if isSelectOrCall {
		rows, err := db.Queryx(query)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		var result []map[string]interface{}
		for rows.Next() {
			row := make(map[string]interface{})
			err = rows.MapScan(row)
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			result = append(result, row)
		}

		return result, nil
	} else {
		result, err := db.Exec(query)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return result, nil
	}

}
