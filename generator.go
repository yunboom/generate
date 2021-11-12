package generate

import (
	"github.com/yunboom/generate/config"
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/internal/check"
)

type Generator struct {
	db       datebase.Database
	executor *Executor
	Err      error
}

func New(config *config.Config) *Generator {
	return &Generator{
		executor: &Executor{
			Data:   make(map[string]*check.BaseStruct),
			Config: config,
		},
	}
}

func (gen *Generator) UseDB(db datebase.Database, err error) {
	if err != nil {
		gen.Err = err
		return
	}
	gen.db = db
}

func (gen *Generator) BindModel(base *check.BaseStruct) {
	if gen.Err != nil {
		return
	}

	gen.executor.Data[base.StructName] = base
}

func (gen *Generator) GenModel(tableName string) *check.BaseStruct {
	if gen.Err != nil {
		return nil
	}
	return gen.GenModelAs(tableName, "")
}

func (gen *Generator) GenModelAs(tableName, modelName string) *check.BaseStruct {
	if gen.Err != nil {
		return nil
	}
	baseStruct, err := check.GenBaseStructs(gen.db, tableName, modelName)
	if err != nil {
		gen.Err = err
		return nil
	}

	return baseStruct
}

func (gen *Generator) Execute() error {
	if gen.Err != nil {
		return gen.Err
	}

	return gen.executor.Execute()
}
