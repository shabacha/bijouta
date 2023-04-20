package data

import (
	"fmt"

	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/shabacha/utils"
)

func NewConnection(config *utils.Configurations, logger hclog.Logger) (*sqlx.DB, error) {
	var conn string
	if config.DBConn != "" {
		conn = config.DBConn
	} else {
		host := config.DBHost
		port := config.DBPort
		user := config.DBUser
		dbName := config.DBName
		password := config.DBPass
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	}
	logger.Debug("connection string", conn)

	db, err := sqlx.Connect("mysql", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
