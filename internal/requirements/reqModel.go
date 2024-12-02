package requirements

// 		requirements
// Provides models of needed start-up configuration
// files and etc...)

type Postgres struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Ssl      bool   `mapstructure:"ssl"`
}

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
