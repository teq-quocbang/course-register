package migration

//nolint:revive
import (
	"git.teqnological.asia/teq-go/teq-pkg/teqlogger"
	"git.teqnological.asia/teq-go/teq-pkg/teqsentry"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/teq-quocbang/course-register/config"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func Up(db *gorm.DB) {
	getDB, err := db.DB()
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	driver, err := mysql.WithInstance(getDB, &mysql.Config{MigrationsTable: "migration"})
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migration", config.GetConfig().MySQL.DBName, driver)
	if err != nil {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		teqsentry.Fatal(err)
		teqlogger.GetLogger().Fatal(err.Error())
	}

	teqlogger.GetLogger().Info("Up done!")
}
