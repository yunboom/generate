package datebase

import (
	"fmt"
	"github.com/yunboom/generate/datebase/driver"
	"github.com/yunboom/generate/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	GetStructFields(tableName string) ([]*model.Field, error)
}

func OpenGorm(driverName, dsn string) (Database, error) {
	switch driverName {
	case driver.Postgres:
		orm, err := gorm.Open(postgres.Open(dsn))
		return &PostgresGorm{dsn: dsn, orm: orm}, err
	case driver.Mysql:
		orm, err := gorm.Open(mysql.Open(dsn))
		return &MysqlGorm{dsn: dsn, orm: orm}, err
	default:
		return nil, fmt.Errorf("driver must be %s or %s", driver.Postgres, driver.Mysql)
	}
}
