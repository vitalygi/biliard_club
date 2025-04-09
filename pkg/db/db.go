package db

import (
	"biliard_club/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
	"strings"
)

type Db struct {
	*gorm.DB
}

func ConnectDb(Dsn string) *gorm.DB {
	slog.Info("connecting to database")
	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		slog.Error("failed to connect to database, trying to create it from default db")
		// trying to create necessary db
		db, err = CreateDb(Dsn)
		if err != nil {
			panic(err)
		}
	}
	slog.Info("successfully connected to database")
	return db
}

func CreateDb(Dsn string) (*gorm.DB, error) {
	els := strings.Split(Dsn, " ")
	var defaultDsn string
	var dbname string
	for _, el := range els {
		if strings.HasPrefix(el, "dbname=") {
			// get dbname from dsn to create it
			dbname = strings.Split(el, "=")[1]
			// set postgres dbname
			defaultDsn = fmt.Sprintf("%s %s", defaultDsn, "dbname=postgres")
		} else {
			defaultDsn = fmt.Sprintf("%s %s", defaultDsn, el)
		}
	}
	defaultDb, err := gorm.Open(postgres.Open(defaultDsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to default database")
		return nil, err
	}
	if err := defaultDb.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname)).Error; err != nil {
		slog.Error("failed to create database",
			"database", dbname)
		return nil, err
	}

	db, err := gorm.Open(postgres.Open(Dsn), &gorm.Config{})
	if err != nil {
		slog.Error("failed to connect to created database",
			"database", dbname)
		return nil, err
	}
	return db, err
}
func NewDb(config *config.DbConfig) *Db {
	return &Db{
		ConnectDb(config.Dsn),
	}
}
