package driver

import (
	"github.com/yunboom/generate/internal/model"
)

type Driver interface {
	GetSchemaName() string
	GetColumnQuery() string
	ToFields(columns []*model.Column) []*model.Field
}
