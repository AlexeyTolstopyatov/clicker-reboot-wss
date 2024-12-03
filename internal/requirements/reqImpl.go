package requirements

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig
// Creates instance of DbConfig Service,
// using external configuration file
// 'database.json' that contains all information
// about expected database
func (args *DbConfig) LoadConfig() error {
	viper.SetConfigName("database") // file's Name
	viper.SetConfigType("json")     // extention
	viper.AddConfigPath(".")        // file's Full-Name

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read: %w", err)
	}

	var config DbConfig

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

// LoadConfig
// Creates instance of AppConfig configuration struct
// using external configuration 'server.json' file
// server.json contains musthave information about
// running the worker.
func (args *AppConfig) LoadConfig() error {
	viper.SetConfigName("server")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read: %w", err)
	}

	var config AppConfig

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

// ToString
// Returns connection-string based on the structure
// of Server Startup options
// returns template of empty string if struct is empty
// or struct pointer is <nil>
func (args *AppConfig) ToString() string {
	template := AppConfig{}

	if args == nil || template == *args {
		return EmptyConfigurationString
	}

	return fmt.Sprintf("%#v", *args)
}

// ToString
// Returns connection-string based on the structure
// of DB parameters (dbname, host, user, password, ssl)
// etc...
// returns template of empty string if struct is empty
// or struct pointer <nil>
func (args *DbConfig) ToString() string {
	template := DbConfig{}

	if args == nil || template == *args {
		return EmptyConfigurationString
	}

	s := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		(*args).Host,
		(*args).Port,
		(*args).User,
		(*args).Name,
		(*args).Password)

	if (*args).Ssl == "verify-ca" {
		// if SSL-mode arguments not empty -- connect the PEM certificate
		s += " sslmode=verify-ca sslrootcert=root.crt"
	} else {
		s += " sslmode=disable"
	}

	return s
}
