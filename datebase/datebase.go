package datebase

import (
	"github.com/yunboom/generate/internal/model"
)

type Database interface {
	GetStructFields(tableName string) ([]*model.Field, error)
}
