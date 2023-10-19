package conf

type Config struct {
	Server server `mapstructure:"server" json:"server" yaml:"server"`
	MySQL  mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	Logger logger `mapstructure:"logger" json:"logger" yaml:"logger"`
}

type server struct {
	bind string
}

type mysql struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Pass     string `mapstructure:"pass"`
	Database string `mapstructure:"database"`
}

type redis struct {
	Host     string `mapstructure:"host"`
	Port     uint16 `mapstructure:"port"`
	Auth     string `mapstructure:"auth"`
	Database int    `mapstructure:"database"`
}

type logger struct {
}
