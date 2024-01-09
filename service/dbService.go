package services

import (
	"database/sql"
	"errors"
	"os"
	"sync"

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

// MainDBService handles database connections and queries
type MainDBService struct {
	db *sql.DB
}

// Query executes a query or a list of SQL statements in a transaction
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

	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	for _, query := range queries {
		rows, err := tx.Query(query)
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

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return results, nil
}

// MainDBServiceConnection creates and initializes the database connection
func MainDBServiceConnection() (*MainDBService, error) {
	dbMainHost := os.Getenv("DB_MAIN_HOST")
	dbMainPort := os.Getenv("DB_MAIN_PORT")
	dbMainUser := os.Getenv("DB_MAIN_USER")
	dbMainPass := os.Getenv("DB_MAIN_PASS")

	// fmt.Println("Environment variables:")
	// fmt.Println("DB_MAIN_HOST:", dbMainHost)
	// fmt.Println("DB_MAIN_PORT:", dbMainPort)
	// fmt.Println("DB_MAIN_USER:", dbMainUser)
	// fmt.Println("DB_MAIN_PASS:", dbMainPass) // Use for debugging, remove in production

	dataSourceName := dbMainUser + ":" + dbMainPass + "@tcp(" + dbMainHost + ":" + dbMainPort + ")/"

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Set max open connections and ping the database
	db.SetMaxOpenConns(10)
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MainDBService{db: db}, nil
}

// Close closes the database connection pool
func (d *MainDBService) Close() {
	if d.db != nil {
		d.db.Close()
	}
}
