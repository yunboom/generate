package datebase

import (
	"github.com/jinzhu/gorm"
	"github.com/yunboom/generate/datebase/driver"
	"github.com/yunboom/generate/internal/model"
	"github.com/yunboom/generate/internal/model/column"
)

type MysqlGorm struct {
	orm *gorm.DB
	dsn string
}

func (g MysqlGorm) GetStructFields(tableName string) (result []*model.Field, err error) {
	columns := make([]column.MysqlColumn, 0)
	schemaName := driver.GetMysqlSchemaName(g.dsn)
	if err = g.orm.Debug().Raw(driver.MysqlColumnQuery, schemaName, tableName).Scan(&columns).Error; err != nil {
		return nil, err
	}

	for _, mysqlColumn := range columns {
		field := mysqlColumn.ToField()
		field.Name = snakeToHump(field.ColumnName)
		result = append(result, field)
	}
	return result, nil
}
