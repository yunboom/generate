package tmpl

const DaoTemplate = NotEditMark + `
package {{.PkgName}}

import (
	"gorm.io/gorm"
)

type {{.StructName}}Wrapper struct {
	{{.StructFieldName}}    *{{.StructName}}
	where   map[string]interface{}
	orderBy string
}

type {{.StructName}}Dao struct {
	db *gorm.DB
}

func New{{.StructName}}Dao(db *gorm.DB) *{{.StructName}}Dao {
	return &{{.StructName}}Dao{db: db}
}

func (dao {{.StructName}}Dao) QueryWrapper({{.StructFieldName}} *{{.StructName}}) *{{.StructName}}Wrapper {
	return &{{.StructName}}Wrapper{
		{{.StructFieldName}}:  {{.StructFieldName}},
		where: make(map[string]interface{}),
	}
}

func (wrap *{{.StructName}}Wrapper) Where(query string, arg interface{}) *{{.StructName}}Wrapper {
	wrap.where[query] = arg
	return wrap
}

func (wrap *{{.StructName}}Wrapper) OrderBy(query string) *{{.StructName}}Wrapper {
	wrap.orderBy = query
	return wrap
}

func (dao {{.StructName}}Dao) Insert({{.StructFieldName}} *{{.StructName}}) error {
	return dao.db.Create({{.StructFieldName}}).Error
}

func (dao {{.StructName}}Dao) DeleteByWrapper(wrap *{{.StructName}}Wrapper) error {
	orm := dao.db.Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}
	return orm.Delete({{.StructName}}{}).Error
}

func (dao {{.StructName}}Dao) DeleteById(id interface{}) error {
	return dao.db.Where("id = ?", id).Delete({{.StructName}}{}).Error
}

func (dao {{.StructName}}Dao) DeleteBatchIds(ids ...interface{}) error {
	return dao.db.Where("id in ?", ids).Delete({{.StructName}}{}).Error
}

func (dao {{.StructName}}Dao) UpdateById(id interface{}, {{.StructFieldName}} *{{.StructName}}) error {
	return dao.db.Where("id = ?", id).Updates({{.StructFieldName}}).Error
}

func (dao {{.StructName}}Dao) SelectById(id interface{}) (*{{.StructName}}, error) {
	{{.StructFieldName}} := &{{.StructName}}{}
	err := dao.db.Table(TableName{{.StructName}}).Where("id = ?", id).First({{.StructFieldName}}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return {{.StructFieldName}}, err
}

func (dao {{.StructName}}Dao) SelectOne(wrap *{{.StructName}}Wrapper) (*{{.StructName}}, error) {
	{{.StructFieldName}} := &{{.StructName}}{}
	orm := dao.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}
	err := orm.Order(wrap.orderBy).First({{.StructFieldName}}).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return {{.StructFieldName}}, err
}

func (dao {{.StructName}}Dao) SelectList(wrap *{{.StructName}}Wrapper) ({{.StructFieldName}}List []{{.StructName}}, err error) {
	orm := dao.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}

	return {{.StructFieldName}}List, orm.Order(wrap.orderBy).Find(&{{.StructFieldName}}List).Error
}

func (dao {{.StructName}}Dao) SelectMaps(wrap *{{.StructName}}Wrapper) (list []map[string]interface{}, err error) {
	{{.StructFieldName}}List := make([]{{.StructName}}, 0)
	orm := dao.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}
	err = orm.Find(&userList).Error
	for _, {{.StructFieldName}} := range {{.StructFieldName}}List {
		m := make(map[string]interface{})
		obj := {{.StructFieldName}}
		{{range .Fields}} m["{{.Name}}"] = obj.{{.Name}}
		{{end}}
		list = append(list, m)
	}

	return
}

func (dao {{.StructName}}Dao) SelectCount(wrap *{{.StructName}}Wrapper) (count int64, err error) {
	orm := dao.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}

	return count, orm.Count(&count).Error
}

func (dao {{.StructName}}Dao) SelectPage(limit, offset int, wrap *{{.StructName}}Wrapper) ({{.StructFieldName}}List []{{.StructName}}, err error) {
	orm := dao.db.Table(TableName{{.StructName}}).Where(wrap.{{.StructFieldName}})
	for query, arg := range wrap.where {
		orm = orm.Where(query, arg)
	}

	err = orm.Limit(limit).Offset(offset).Order(wrap.orderBy).Find(&{{.StructFieldName}}List).Error
	return {{.StructFieldName}}List, err
}

`
