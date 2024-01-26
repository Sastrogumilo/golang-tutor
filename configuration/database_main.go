package configuration

import (
	"os"
)

type DBConfig struct {
	Host                  string
	User                  string
	Password              string
	Database              string
	Port                  string
	MultipleStatements    bool
	EnableKeepAlive       bool
	KeepAliveInitialDelay int
}

func MainDBConfig() *DBConfig {
	// port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	return &DBConfig{
		Host:     os.Getenv("DB_MAIN_HOST"),
		User:     os.Getenv("DB_MAIN_USER"),
		Password: os.Getenv("DB_MAIN_PASS"),
		Port:     os.Getenv("DB_MAIN_PORT"),
		// MultipleStatements: false,
		// EnableKeepAlive:    true,
		// KeepAliveInitialDelay: keepAliveInitialDelay,
	}
}
