package column

import "strings"

type typeMap map[string]func(string) string

var (
	defaultDataType         = "string"
	dataType        typeMap = map[string]func(detailType string) string{
		"int":        func(string) string { return "int32" },
		"integer":    func(string) string { return "int32" },
		"smallint":   func(string) string { return "int32" },
		"mediumint":  func(string) string { return "int32" },
		"bigint":     func(string) string { return "int64" },
		"float":      func(string) string { return "float32" },
		"double":     func(string) string { return "float64" },
		"decimal":    func(string) string { return "float64" },
		"char":       func(string) string { return "string" },
		"varchar":    func(string) string { return "string" },
		"tinytext":   func(string) string { return "string" },
		"mediumtext": func(string) string { return "string" },
		"longtext":   func(string) string { return "string" },
		"binary":     func(string) string { return "[]byte" },
		"varbinary":  func(string) string { return "[]byte" },
		"tinyblob":   func(string) string { return "[]byte" },
		"blob":       func(string) string { return "[]byte" },
		"mediumblob": func(string) string { return "[]byte" },
		"longblob":   func(string) string { return "[]byte" },
		"text":       func(string) string { return "string" },
		"json":       func(string) string { return "string" },
		"enum":       func(string) string { return "string" },
		"time":       func(string) string { return "time.Time" },
		"date":       func(string) string { return "time.Time" },
		"datetime":   func(string) string { return "time.Time" },
		"timestamp":  func(string) string { return "time.Time" },
		"year":       func(string) string { return "int32" },
		"bit":        func(string) string { return "[]uint8" },
		"boolean":    func(string) string { return "bool" },
		"tinyint": func(detailType string) string {
			if strings.HasPrefix(detailType, "tinyint(1)") {
				return "bool"
			}
			return "int32"
		},
	}
)

func (m typeMap) Get(dataType, detailType string) string {
	if convert, ok := m[dataType]; ok {
		return convert(detailType)
	}
	return defaultDataType
}
