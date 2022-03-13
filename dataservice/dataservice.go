package dataservice

import (
	"fmt"

	"github.com/go-pg/pg"
)

type DataServiceRepository struct {
	DBConfig *DBConfig
}

type DBConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	DbName   string
	Timeout  int
	SSLMode  string
	DB       *pg.DB
}

func (dbConfig *DBConfig) OpenConnection() {
	address := fmt.Sprintf("%s:%s", dbConfig.Host, dbConfig.Port)

	fmt.Println("open connection to ", address)

	options := &pg.Options{
		User:     dbConfig.UserName,
		Password: dbConfig.Password,
		Addr:     address,
		Database: dbConfig.DbName,
		PoolSize: 50,
	}

	con := pg.Connect(options)
	if con == nil {
		fmt.Println("connection failed")
	}

	dbConfig.DB = con
}
