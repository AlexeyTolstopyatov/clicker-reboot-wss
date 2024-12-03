package requirements

// 		requirements
// Provides models of needed start-up configuration
// files and etc...)

var EmptyConfigurationString = ""

type DbConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Ssl      string `mapstructure:"ssl"`
}

type AppConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type ConfigConvertable interface {
	ToString() string
}
