package column

import (
	"bytes"
	"fmt"
	"github.com/yunboom/generate/internal/model"
)

type MysqlColumn struct {
	TableName     string `gorm:"column:TABLE_NAME"`
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
	DataType      string `gorm:"column:DATA_TYPE"`
	ColumnKey     string `gorm:"column:COLUMN_KEY"`
	ColumnType    string `gorm:"column:COLUMN_TYPE"`
	ColumnDefault string `gorm:"column:COLUMN_DEFAULT"`
	Extra         string `gorm:"column:EXTRA"`
	IsNullable    string `gorm:"column:IS_NULLABLE"`
}

func (c MysqlColumn) ToField() *model.Field {
	memberType := dataType.Get(c.DataType, c.ColumnType)
	return &model.Field{
		Name:          c.ColumnName,
		Type:          memberType,
		ColumnName:    c.ColumnName,
		ColumnComment: c.ColumnComment,
		GORMTag:       c.buildGormTag(),
		JSONTag:       c.ColumnName,
		XORMTag:       c.buildGormTag(),
	}
}

func (c *MysqlColumn) IsPrimaryKey() bool {
	return c != nil && c.ColumnKey == "PRI"
}

func (c *MysqlColumn) IsAutoIncrement() bool {
	return c != nil && c.Extra == "auto_increment"
}

func (c *MysqlColumn) buildGormTag() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("column:%s;type:%s", c.ColumnName, c.ColumnType))
	if c.IsPrimaryKey() {
		buf.WriteString(";primaryKey")
		if !c.IsAutoIncrement() {
			buf.WriteString(";autoIncrement:false")
		}
	} else if c.IsNullable != "YES" {
		buf.WriteString(";not null")
	}

	if c.ColumnDefault != "" {
		buf.WriteString(fmt.Sprintf(";default:%s", c.ColumnDefault))
	}
	return buf.String()
}
