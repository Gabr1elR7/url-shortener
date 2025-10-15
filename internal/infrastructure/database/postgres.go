package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string, models ...interface{}) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Error al conectar con PostgreSQL: %v", err)
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatalf("❌ Error en migraciones: %v", err)
	}

	log.Println("✅ Conectado y migrado con éxito a PostgreSQL")
	return db
}