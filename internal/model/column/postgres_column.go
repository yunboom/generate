package column

import (
	"bytes"
	"fmt"
	"github.com/yunboom/generate/internal/model"
)

type PostgresColumn struct {
	ColumnName    string `gorm:"column:column_name"`
	ColumnComment string `gorm:"column:column_comment"`
	DataType      string `gorm:"column:data_type"`
	IsNullable    string `gorm:"column:is_nullable"`
	ColumnType    string `gorm:"column:column_type"`
}

func (p PostgresColumn) ToField() *model.Field {
	memberType := dataType.Get(p.DataType, "")
	return &model.Field{
		Name:          p.ColumnName,
		Type:          memberType,
		ColumnName:    p.ColumnName,
		ColumnComment: p.ColumnComment,
		GORMTag:       p.buildGormTag(),
		JSONTag:       p.ColumnName,
		XORMTag:       p.buildGormTag(),
	}
}

func (p *PostgresColumn) buildGormTag() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("column:%s;type:%s", p.ColumnName, p.DataType))
	if p.IsNullable != "YES" {
		buf.WriteString(";not null")
	}

	return buf.String()
}
