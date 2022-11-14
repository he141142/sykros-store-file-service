package pg_database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"gorm.io/gorm"
	"log"
)
import (
	"gorm.io/driver/postgres"
)

func NewPG(connstring string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connstring), &gorm.Config{})
}

func PGMigrate(migrationDir, connstring string) error {
	mig, err := migrate.New(fmt.Sprintf("file://%s", migrationDir), connstring)
	if err != nil {
		return err
	}
	if err := mig.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
		log.Println("No migration needed")
	}
	return nil
}
