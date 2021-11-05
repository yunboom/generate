package datebase

import (
	"github.com/yunboom/generate/internal/model/column"
	"xorm.io/xorm"
)

type X struct {
	orm *xorm.Engine
}

func (x X) GetTbColumns(tableName string) (*column.MysqlColumn, error) {
	panic("implement me")
}

func (x X) Schema(modelName string) string {
	panic("implement me")
}
