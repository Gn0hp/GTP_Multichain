package postgres

import (
	"errors"
	"fmt"
)

type Config struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     uint16 `mapstructure:"port" json:"port" yaml:"port"`
	User     string `mapstructure:"user" json:"user" yaml:"user"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Database string `mapstructure:"database" json:"database" yaml:"database"`

	Params map[string]string
}

// GetDSN returns a PostgreSQL driver compatible data source name.
func GetDSN(c Config) string {
	var params string
	if len(c.Params) > 0 {
		var query string
		for key, value := range c.Params {
			if query != "" {
				query += "&"
			}
			query += key + "=" + value
		}
		params = fmt.Sprintf("%s", query)
	}
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		params)
	return connStr
}

func (c Config) Validate() error {
	if c.Host == "" {
		return errors.New("database host is required")
	}
	if c.Port == 0 {
		return errors.New("database port is required")
	}
	if c.User == "" {
		return errors.New("database user is required")
	}
	if c.Database == "" {
		return errors.New("database name is required")
	}
	return nil
}
