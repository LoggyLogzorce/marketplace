package db

import (
	"authService/internal/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB
var migrate = make([]func(), 0)
var cfg *configs.Config
var connect string

func init() {
	cfg = configs.Get()
	connect = "host=" + cfg.HostDb +
		" port=" + cfg.PortDb +
		" user= " + cfg.User +
		" password=" + cfg.Password +
		" dbname=" + cfg.DbName +
		" sslmode=" + cfg.SslMode
}

func Add(mf func()) {
	migrate = append(migrate, mf)
}

func DB() *gorm.DB {
	return database
}

func Connect() {
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	database = db
	log.Println("Connected to the database")
}

func Migrate() {
	for _, f := range migrate {
		f()
	}
	log.Println("Migration has been completed")
}
