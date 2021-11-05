package model

type Field struct {
	Name          string
	Type          string
	ColumnName    string
	ColumnComment string
	JSONTag       string
	GORMTag       string
	XORMTag       string
}
