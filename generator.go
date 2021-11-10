package generate

import (
	"bytes"
	"fmt"
	"github.com/yunboom/generate/datebase"
	"github.com/yunboom/generate/internal/check"
	tmpl "github.com/yunboom/generate/internal/template"
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
	MapperPath  string
	ModelPkg    string
	ServicePath string
	HandlePath  string
}

func WithModelPath(modelPath string) CfgOpt {
	return func(config *Config) {
		config.ModelPath = modelPath
	}
}

func WithMapperPath(mapperPath string) CfgOpt {
	return func(config *Config) {
		config.MapperPath = mapperPath
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

	if err := gen.genModelFile(); err != nil {
		return err
	}

	if err := gen.genQueryFile(); err != nil {
		return err
	}

	return nil
}

func (gen *Generator) genModelFile() error {
	outPath, err := genAbsPath(gen.ModelPath)
	if err != nil {
		return err
	}

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
		if err := render(tmpl.ModelTemplate, &buf, baseStruct); err != nil {
			return err
		}

		modelFile := fmt.Sprint(outPath, baseStruct.TableName, ".go")
		fmt.Printf("generate model file(table <%s> -> {%s.%s}): %s \n", baseStruct.TableName, baseStruct.ModelPkg, baseStruct.StructName, modelFile)
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

func (gen *Generator) genQueryFile() error {
	_, err := genAbsPath(gen.MapperPath)
	if err != nil {
		return err
	}

	var buf bytes.Buffer

	err = render(tmpl.HeaderTmpl, &buf, "mapper")
	if err != nil {
		return err
	}

	return nil
}

func genAbsPath(path string) (string, error) {
	outPath, err := filepath.Abs(path)
	if err != nil {
		return outPath, err
	}
	if err := mkdirAll(os.ModePerm, outPath); err != nil {
		return outPath, err
	}

	return outPath, nil
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
