package generate

import (
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/internal/check"
)

type CfgOpt func(*Config)

type Config struct {
	db datebase.Database

	ModelPath   string
	ServicePath string
	HandlePath  string
}

func WithModelPath(modelPath string) CfgOpt {
	return func(config *Config) {
		config.ModelPath = modelPath
	}
}

func WithServicePath(servicePath string) CfgOpt {
	return func(config *Config) {
		config.ServicePath = servicePath
	}
}

func WithHandlePath(handlePath string) CfgOpt {
	return func(config *Config) {
		config.HandlePath = handlePath
	}
}

func NewConfig(opts ...CfgOpt) *Config {
	defaultCfg := &Config{
		ModelPath:   "",
		ServicePath: "",
		HandlePath:  "",
	}
	for _, opt := range opts {
		opt(defaultCfg)
	}

	return defaultCfg
}

type Generator struct {
	*Config
}

func New(config *Config) *Generator {
	return &Generator{Config: config}
}

func (gen *Generator) UseDB(db datebase.Database, err error) error {
	if err != nil {
		return err
	}
	gen.db = db
	return nil
}

func (gen *Generator) BindModel(base *check.BaseStruct) {

}

func (gen *Generator) GenModel(tableName string) (*check.BaseStruct, error) {
	return gen.GenModelAs(tableName, "")
}

func (gen *Generator) GenModelAs(tableName, modelName string) (*check.BaseStruct, error) {
	return check.GenBaseStructs(gen.db, tableName, modelName)
}
