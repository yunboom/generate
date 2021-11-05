package driver

import (
	"github.com/yunboom/generate/internal/model"
	"regexp"
	"strings"
)

const (
	Mysql            = "mysql"
	Postgres         = "postgres"
	mysqlColumnQuery = "SELECT TABLE_NAME,COLUMN_NAME,COLUMN_COMMENT,DATA_TYPE,IS_NULLABLE,COLUMN_KEY,COLUMN_TYPE,COLUMN_DEFAULT,EXTRA " +
		"FROM information_schema.COLUMNS " +
		"WHERE table_schema = ? AND table_name =? " +
		"ORDER BY ORDINAL_POSITION"
)

type MysqlDriver struct {
	DriverName string
	Dsn        string
}

func (m MysqlDriver) ToFields(columns []*model.Column) []*model.Field {
	fields := make([]*model.Field, len(columns))
	for i, column := range columns {
		fields[i] = column.ToFieldWithMysql()
	}

	return fields
}

func (m MysqlDriver) GetColumnQuery() string {
	return mysqlColumnQuery
}

var dbNameReg = regexp.MustCompile(`/\w+\??`)

func (m MysqlDriver) GetSchemaName() string {
	dbName := dbNameReg.FindString(m.Dsn)
	if len(dbName) < 3 {
		return ""
	}
	end := len(dbName)
	if strings.HasSuffix(dbName, "?") {
		end--
	}
	return dbName[1:end]
}
