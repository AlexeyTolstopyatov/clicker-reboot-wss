package requirements

import (
	"fmt"
	"github.com/spf13/viper"
)

// LoadPostgres
// Creates instance of Postgres Service,
// using external configuration file
// 'database.json' that contains all information
// about expected database
func (args *Postgres) LoadPostgres() error {
	viper.SetConfigName("database") // file's Name
	viper.SetConfigType("json")     // extention
	viper.AddConfigPath(".")        // file's Full-Name

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read: %w", err)
	}

	var config Postgres

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}

// LoadServer
// Creates instance of Server configuration struct
// using external configuration 'server.json' file
// server.json contains musthave information about
// running the worker.
func (args *Server) LoadServer() error {
	viper.SetConfigName("server")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("unable to read: %w", err)
	}

	var config Server

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to unmarshal: %w", err)
	}

	return nil
}
