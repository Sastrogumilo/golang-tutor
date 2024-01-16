package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var MainDBServiceInstance *MainDBService
var MainDBServiceOnce sync.Once

func SetLocal[T any](c *fiber.Ctx, key string, value T) {
	c.Locals(key, value)
}
func GetLocal[T any](c *fiber.Ctx, key string) T {
	return c.Locals(key).(T)
}

type MainDBService struct {
	db *sql.DB
}

func (d *MainDBService) Query(queries ...string) ([]map[string]interface{}, error) {
	if d.db == nil {
		return nil, errors.New("Database connection is not initialized")
	}

	return d.query(queries...)
}

func (d *MainDBService) query(queries ...string) ([]map[string]interface{}, error) {
	if len(queries) == 0 {
		return nil, nil
	}

	var results []map[string]interface{}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for _, query := range queries {
		rows, err := d.db.QueryContext(ctx, query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			values := make([]interface{}, len(columns))
			columnPointers := make([]interface{}, len(columns))

			for i := range columns {
				columnPointers[i] = &values[i]
			}

			if err := rows.Scan(columnPointers...); err != nil {
				return nil, err
			}

			row := make(map[string]interface{})

			for i, colName := range columns {
				val := columnPointers[i].(*interface{})
				row[colName] = *val
			}

			results = append(results, row)
		}
	}

	return results, nil
}

func MainDBServiceConnection() (*MainDBService, error) {
	dbMainHost := os.Getenv("DB_MAIN_HOST")
	dbMainPort := os.Getenv("DB_MAIN_PORT")
	dbMainUser := os.Getenv("DB_MAIN_USER")
	dbMainPass := os.Getenv("DB_MAIN_PASS")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbMainUser, dbMainPass, dbMainHost, dbMainPort)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Minute * 5)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MainDBService{db: db}, nil
}

func (d *MainDBService) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
