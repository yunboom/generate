package datebase

import (
	"errors"
	"fmt"
	"github.com/yunboom/generate/datebase/driver"
	"github.com/yunboom/generate/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type G struct {
	orm    *gorm.DB
	driver driver.Driver
}

func (g G) GetStructFields(tableName string) (result []*model.Field, err error) {
	if g.orm == nil {
		return nil, errors.New("gorm is nil")
	}

	columns, err := getColumns(g.orm, g.driver.GetColumnQuery(), g.driver.GetSchemaName(), tableName)
	if err != nil {
		return nil, err
	}

	return g.driver.ToFields(columns), nil
}

func getColumns(orm *gorm.DB, columnQuery string, schemaName string, tableName string) ([]*model.Column, error) {
	result := make([]*model.Column, 0)
	err := orm.Debug().Raw(columnQuery, schemaName, tableName).Scan(&result).Error

	return result, err
}

func NewGorm(driverName, dsn string) (*G, error) {
	var dia gorm.Dialector
	var d driver.Driver
	switch driverName {
	case driver.Postgres:
		dia = postgres.Open(dsn)
	case driver.Mysql:
		dia = mysql.Open(dsn)
		d = driver.MysqlDriver{DriverName: driverName, Dsn: dsn}
	default:
		return nil, fmt.Errorf("driver must be %s or %s", driver.Postgres, driver.Mysql)
	}
	orm, err := gorm.Open(dia)
	if err != nil {
		return nil, err
	}

	return &G{orm: orm, driver: d}, nil
}
