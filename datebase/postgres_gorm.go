package datebase

import (
	"github.com/yunboom/generate/datebase/driver"
	"github.com/yunboom/generate/internal/model"
	"github.com/yunboom/generate/internal/model/column"
	"gorm.io/gorm"
)

type PostgresGorm struct {
	orm *gorm.DB
	dsn string
}

func (p PostgresGorm) GetStructFields(tableName string) (result []*model.Field, err error) {
	columns := make([]column.PostgresColumn, 0)
	if err = p.orm.Raw(driver.PostgresColumnQuery, tableName).Scan(&columns).Error; err != nil {
		return nil, err
	}

	for _, postgresColumn := range columns {
		field := postgresColumn.ToField()
		field.Name = snakeToHump(field.ColumnName)
		result = append(result, field)
	}
	return result, nil
}
