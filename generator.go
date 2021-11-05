package generate

import (
	"bytes"
	"fmt"
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/internal/check"
	"github.com/yunboom/generate/internal/template/model"
	"golang.org/x/tools/imports"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

type CfgOpt func(*Config)

type Config struct {
	db datebase.Database

	ModelPath   string
	ModelPkg    string
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
		ModelPath:   "../dal",
		ServicePath: "",
		HandlePath:  "",
		ModelPkg:    "model",
	}
	for _, opt := range opts {
		opt(defaultCfg)
	}

	return defaultCfg
}

type Generator struct {
	*Config
	Data map[string]*check.BaseStruct
	Err  error
}

func New(config *Config) *Generator {
	return &Generator{Config: config, Data: make(map[string]*check.BaseStruct)}
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

	gen.Data[base.TableName] = base
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

	abs, err := filepath.Abs(gen.ModelPath)
	if err != nil {
		return err
	}

	if err := mkdirAll(os.ModePerm, abs); err != nil {
		return err
	}

	if err := gen.genModelFile(abs); err != nil {
		return err
	}

	return nil
}

func (gen *Generator) genModelFile(outPath string) error {
	outPath = fmt.Sprint(outPath, "/", filepath.Base(gen.ModelPkg), "/")

	for _, baseStruct := range gen.Data {
		if baseStruct == nil {
			continue
		}
		if err := mkdirAll(os.ModePerm, outPath); err != nil {
			return err
		}

		var buf bytes.Buffer
		baseStruct.ModelPkg = gen.ModelPkg
		if err := render(model.Template, &buf, baseStruct); err != nil {
			return err
		}

		modelFile := fmt.Sprint(outPath, baseStruct.TableName, ".go")
		fmt.Printf("\ngenerate model file path : %s \n", modelFile)
		if err := gen.output(modelFile, buf.Bytes()); err != nil {
			return err
		}

	}

	return nil
}

func (gen *Generator) output(fileName string, content []byte) error {
	result, err := imports.Process(fileName, content, nil)
	if err != nil {
		fmt.Println(string(content))
		return err
	}
	return outputFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, result)
}

func outputFile(filename string, flag int, data []byte) error {
	out, err := os.OpenFile(filename, flag, 0640)
	if err != nil {
		return fmt.Errorf("open out file fail: %w", err)
	}
	return output(out, data)
}

func output(wr io.WriteCloser, data []byte) (err error) {
	defer func() {
		if e := wr.Close(); e != nil {
			err = fmt.Errorf("close file fail: %w", e)
		}
	}()

	if _, err = wr.Write(data); err != nil {
		return fmt.Errorf("write file fail: %w", err)
	}
	return nil
}

func render(tmpl string, wr io.Writer, data interface{}) error {
	t, err := template.New(tmpl).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(wr, data)
}

func mkdirAll(perm fs.FileMode, path ...string) error {
	for _, p := range path {
		if _, err := os.Stat(p); err != nil {
			if err := os.MkdirAll(p, perm); err != nil {
				return err
			}
		}
	}

	return nil
}
