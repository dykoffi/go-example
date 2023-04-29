package configs

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() *gorm.DB {
	// Ouvre une connexion à la base de données Postgres
	dsn := "host=localhost user=user password=4134 dbname=db port=21822 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
