package datebase

import (
	"github.com/yunboom/generate/datebase/driver"
	"github.com/yunboom/generate/internal/model"
	"xorm.io/xorm"
)

type X struct {
	orm *xorm.Engine
}

func (x X) GetTbColumns(tableName string) (*model.Column, error) {
	panic("implement me")
}

func NewX(orm driver.Driver) (*X, error) {
	panic("implement me")
}

func (x X) Schema(modelName string) string {
	panic("implement me")
}
