package datastore

import (
	"github.com/go-sql-driver/mysql"
	"github.com/shabacha/pkg/config"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	mysqlConfig := &mysql.Config{
		User:                 config.C.Database.User,
		Passwd:               config.C.Database.Password,
		Net:                  config.C.Database.Net,
		Addr:                 "127.0.0.1:3306",
		DBName:               config.C.Database.DBName,
		AllowNativePasswords: config.C.Database.AllowNativePasswords,
		Params: map[string]string{
			"parseTime": config.C.Database.Params.ParseTime,
		},
	}

	db, err := gorm.Open(gormmysql.Open(mysqlConfig.FormatDSN()), &gorm.Config{})
	if err != nil {
		return nil
	}
	// Set up a connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// Return the database connection
	return db
}
