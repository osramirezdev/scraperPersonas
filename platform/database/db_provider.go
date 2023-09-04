package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/osramirezdev/scraperPersonas/app/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBProvider struct {
	*gorm.DB
}

var DBP DBProvider

func init() {
	enverr := godotenv.Load(".env")
	fmt.Println(enverr)
	if enverr != nil {
		log.Fatalf("Error obteniendo envs params. Err: %s", enverr)
	}

	db_username := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_database := os.Getenv("DB_NAME")
	db_ssl := os.Getenv("DB_SSL_MODE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db_host,
		db_port,
		db_username,
		db_password,
		db_database,
		db_ssl,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB := database

	errM := database.AutoMigrate(
		&entities.Persona{},
		&entities.Ips{},
		&entities.Funcionario{},
		&entities.Asignacion{},
		&entities.Archivo{},
	)
	if errM != nil {
		fmt.Println("error migrate", errM)
	}

	DBP = DBProvider{DB}
}
