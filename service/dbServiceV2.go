package services

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"

	configuration "belajar_golang/configuration"

	"github.com/jmoiron/sqlx"
)

func NewConvertToStruct(data []map[string]interface{}, result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("result argument must be a slice address")
	}

	slicev := resultv.Elem()
	elemt := slicev.Type().Elem()

	for _, item := range data {
		elemv := reflect.New(elemt)

		jsonData, err := json.Marshal(item)
		if err != nil {
			return err
		}

		err = json.Unmarshal(jsonData, elemv.Interface())
		if err != nil {
			log.Printf("Warning: %v", err)
			elemv.Elem().Set(reflect.Zero(elemt))
		}

		slicev = reflect.Append(slicev, elemv.Elem())
	}

	resultv.Elem().Set(slicev)
	return nil
}

type MainDBService struct {
	db *sqlx.DB
}

type QueryResult []map[string]interface{}

var mainDBService *MainDBService

func MainDB() *MainDBService {
	if mainDBService != nil {
		return mainDBService
	}

	config := configuration.MainDBConfig()
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.User, config.Password, config.Host, config.Port))
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(2)

	mainDBService = &MainDBService{db: db}
	return mainDBService
}

func (s *MainDBService) Query(query string) (*QueryResult, error) {

	defer func() {
		if r := recover(); r != nil {
			// error = fmt.Errorf("panic recovered: %v", r)
			// result = nil
			log.Println(r)
		}
	}()
	// use s.db to execute your query
	db := s.db

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
		rows.Close()
		return (*QueryResult)(&result), err

		/**
		 * Code above is test using Normal method, the code below
		 * is using prepare statement
		 */
		// stmt, err := db.Preparex(query)
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return nil, err
		// }
		// defer stmt.Close() // don't forget to close the statement when you're done with it

		// rows, err := stmt.Queryx()
		// if err != nil {
		// 	log.Println(err.Error())
		// 	return nil, err
		// }
		// defer rows.Close() // close the rows when you're done with them

		// var result []map[string]interface{}
		// for rows.Next() {
		// 	row := make(map[string]interface{})
		// 	err = rows.MapScan(row)
		// 	if err != nil {
		// 		log.Println(err.Error())
		// 		return nil, err
		// 	}
		// 	result = append(result, row)
		// }

		// return result, nil

	} else {
		// result, err := db.Exec(query)
		_, err := db.Exec(query)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}

		return &QueryResult{nil}, err
	}
}

func (r *QueryResult) ToStruct(result interface{}) error {
	//jika error
	if *r == nil {
		return fmt.Errorf("result is nil")
	}

	hasil_query := []map[string]interface{}(*r)

	err := NewConvertToStruct(hasil_query, result)

	if err != nil {
		return err
	}

	return nil
}

func (s *MainDBService) QueryToStruct(query string, result interface{}) (*QueryResult, error) {
	queryResult, err_query := s.Query(query)
	if err_query != nil {
		return nil, fmt.Errorf("query failed")
	}
	err := queryResult.ToStruct(result)
	if err != nil {
		return nil, err
	}
	return queryResult, nil
}
