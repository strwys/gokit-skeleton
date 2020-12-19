package db

import (
	"database/sql"
	"fmt"

	"github.com/cecepsprd/gokit-skeleton/commons/config"
	_ "github.com/go-sql-driver/mysql"
)

func MysqlConnect(cfg config.Config) (*sql.DB, error) {

	dbConnString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MysqlDB.User,
		cfg.MysqlDB.Password,
		cfg.MysqlDB.Host,
		cfg.MysqlDB.Port,
		cfg.MysqlDB.Name,
	)

	db, err := sql.Open("mysql", dbConnString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
