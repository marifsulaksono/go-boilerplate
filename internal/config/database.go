package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/constants"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func (db *Database) ConnectDatabase(ctx context.Context, database string) (DB *gorm.DB, err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.Info,            // Log level
			IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                   // Don't include params in the SQL log
			Colorful:                  false,                  // Disable color
		},
	)

	/*
		Mapping to connect various sql database such as MySQL, Postgres, SQL Server
		Customize to your needs

		more info contact me @marifsulaksono
	*/

	switch database {
	case constants.DB_MYSQL:
		DB, err = db.mysqlConnector(&newLogger)
		if err != nil {
			return nil, err
		}
	case constants.DB_POSTGRESQL:
		DB, err = db.postgreConnector(&newLogger)
		if err != nil {
			return nil, err
		}
	case constants.DB_SQL_SERVER:
		DB, err = db.sqlServerConnector(&newLogger)
		if err != nil {
			return nil, err
		}
	}

	return DB, nil
}

func (db *Database) mysqlConnector(looger *logger.Interface) (*gorm.DB, error) {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db.Username, db.Password, db.Host, db.Port, db.Name)
	log.Println("URL config mysql:", url)
	return gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: *looger,
	})
}

func (db *Database) postgreConnector(looger *logger.Interface) (*gorm.DB, error) {
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", db.Host, db.Username, db.Password, db.Name, db.Port)
	fmt.Println("URL config postgres:", url)
	return gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: *looger,
	})
}

func (db *Database) sqlServerConnector(looger *logger.Interface) (*gorm.DB, error) {
	url := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", db.Username, db.Password, db.Host, db.Port, db.Name)
	return gorm.Open(sqlserver.Open(url), &gorm.Config{
		Logger: *looger,
	})
}
