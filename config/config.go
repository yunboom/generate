package config

type CfgOpt func(*Config)

type Config struct {
	ModelPath   string //model代码输出路径
	ModelPkg    string //model层包名
	DaoPath     string //dao代码输出路径
	DaoPkg      string //dao层包名
	ServicePath string //service代码输出路径
	HandlePath  string //handle代码输出路径
}

func WithModelPath(modelPath string) CfgOpt {
	return func(config *Config) {
		config.ModelPath = modelPath
	}
}

func WithModelPkg(modelPkg string) CfgOpt {
	return func(config *Config) {
		config.ModelPkg = modelPkg
	}
}

func WithDaoPath(daoPath string) CfgOpt {
	return func(config *Config) {
		config.DaoPath = daoPath
	}
}

func WithDaoPkg(daoPkg string) CfgOpt {
	return func(config *Config) {
		config.DaoPkg = daoPkg
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

func New(opts ...CfgOpt) *Config {
	defaultCfg := &Config{
		ModelPath:   "./",
		DaoPath:     "./",
		ServicePath: "./",
		HandlePath:  "./",
		ModelPkg:    "model",
		DaoPkg:      "model",
	}
	for _, opt := range opts {
		opt(defaultCfg)
	}

	return defaultCfg
}
