package driver

import (
	"regexp"
	"strings"
)

const (
	Mysql            = "mysql"
	MysqlColumnQuery = `
		SELECT
			TABLE_NAME,
			COLUMN_NAME,
			COLUMN_COMMENT,
			DATA_TYPE,
			IS_NULLABLE,
			COLUMN_KEY,
			COLUMN_TYPE,
			COLUMN_DEFAULT,
			EXTRA 
		FROM
			information_schema.COLUMNS 
		WHERE
			table_schema = ? 
			AND table_name = ? 
		ORDER BY
			ORDINAL_POSITION
`
)

var mysqlReg = regexp.MustCompile(`/\w+\??`)

func GetMysqlSchemaName(dsn string) string {
	dbName := mysqlReg.FindString(dsn)
	if len(dbName) < 3 {
		return ""
	}
	end := len(dbName)
	if strings.HasSuffix(dbName, "?") {
		end--
	}
	return dbName[1:end]
}
