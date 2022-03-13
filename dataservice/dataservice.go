package dataservice

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/spf13/viper"
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

func (dbConfig *DBConfig) InitDbConfig() {
	dbConfig.UserName = viper.GetString("db.postgres.username")
	dbConfig.Password = viper.GetString("db.postgres.password")
	dbConfig.Host = viper.GetString("db.postgres.host")
	dbConfig.Port = viper.GetString("db.postgres.port")
	dbConfig.DbName = viper.GetString("db.postgres.name")
	dbConfig.Timeout = 20
	dbConfig.SSLMode = "disabled"
}

func (dbConfig *DBConfig) OpenConnection() {

	fmt.Println("open connection to ", dbConfig.Host)

	options := &pg.Options{
		User:     dbConfig.UserName,
		Password: dbConfig.Password,
		Addr:     dbConfig.Host,
		Database: dbConfig.DbName,
		PoolSize: 50,
	}

	con := pg.Connect(options)
	if con == nil {
		fmt.Println("connection failed")
	}

	dbConfig.DB = con
}
