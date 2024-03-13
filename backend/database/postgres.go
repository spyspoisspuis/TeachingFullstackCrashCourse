package database

import (
	"backend/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

func NewPostgresDatabase(cfg *config.Config) Database {
	// log => current time
	// fmt => print
	log.Println("Connecting to Postgres....")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		cfg.Db.Host,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Dbname,
		cfg.Db.Port,
		cfg.Db.Sslmode,
		cfg.Db.Timezone,
	)
	log.Println("With config", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// panic => print ; exit program
		panic(err)
	}

	log.Println("Connected to Postgres!")

	return &postgresDatabase{
		Db: db,
	}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return p.Db
}
