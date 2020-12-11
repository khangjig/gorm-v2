package datastores

import (
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// Register using Golang migrate.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	getDD, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(getDD, &mysql.Config{
		MigrationsTable: "custom_schema_migrations",
	})
	if err != nil {
		log.Fatal(err)
	}

	var m *migrate.Migrate

	m, err = migrate.NewWithDatabaseInstance("file://./migrations", os.Getenv("DB_NAME"), driver)

	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()

	if err != nil {
		if !strings.Contains(err.Error(), "no change") && err.Error() != "file does not exist" {
			log.Fatal(err)
		}
	}

	log.Println("Migrate done!")
}

type BatchDB struct {
	*gorm.DB
}
