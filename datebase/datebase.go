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

func snakeToHump(str string) string {
	toStr := make([]byte, 0, len(str))
	i := 0
	for i < len(str) {
		if (str[i] < 97 || str[i] > 122) && str[i] != '_' {
			toStr = append(toStr, str[i])
			i++
			continue
		}

		if i == 0 && str[i] >= 97 && str[i] <= 122 {
			toStr = append(toStr, str[i]-32)
			i++
			continue
		}

		if i == len(str)-1 && str[i] == '_' {
			i++
			continue
		}

		if str[i] == '_' {
			if str[i+1] >= 97 && str[i+1] <= 122 {
				toStr = append(toStr, str[i+1]-32)
				i += 2
			}
			continue
		}

		toStr = append(toStr, str[i])
		i++
	}

	return string(toStr)
}
