package global

type App struct {
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type ConfigType struct {
	App    `mapstructure:"app" json:"app" yaml:"app"`
	Author string `mapstructure:"author" json:"author" yaml:"author"`
}

var Config ConfigType
