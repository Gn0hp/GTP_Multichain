package config

import (
	"errors"
	"eth_bsc_multichain/internal/service/db/postgres"
	"eth_bsc_multichain/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

type appConfig struct {
	// Http Server Address
	HttpAddress string `mapstructure:"httpAddr"`

	// GRpc Server Address
	GRpcAddress string `mapstructure:"grpcAddr"`

	// Storage is the storage backend of the app
	Storage string `mapstructure:"storage"`
}
type Configuration struct {
	// Log is the configuration for the logger.
	Log logger.Config `mapstructure:"logger"`

	Telemetry struct {
		Addr string `mapstructure:"address"`
	} `mapstructure:"telemetry"`

	App appConfig `mapstructure:"app"`

	Database postgres.Config `mapstructure:"database"`
}

// Process post-processes configuration after loading it.
func (c Configuration) Process() error {
	return nil
}
func (a appConfig) Validate() error {
	if a.HttpAddress == "" {
		return errors.New("http app server address is required")
	}

	if a.GRpcAddress == "" {
		return errors.New("grpc app server address is required")
	}
	if a.Storage != "in-memory" && a.Storage != "database" {
		return errors.New("storage must be either in-memory or database")
	}
	return nil
}
func (c Configuration) Validate() error {
	if c.Telemetry.Addr == "" {
		return errors.New("telemetry http server address is required")
	}

	if err := c.App.Validate(); err != nil {
		return err
	}

	if err := c.Database.Validate(); err != nil {
		return err
	}

	return nil
}

// setDefaultConfig sets the default configuration if not exist any values in conf file for the application.
func setDefaultConfig(v *viper.Viper, f *pflag.FlagSet) {
	v.SetConfigName("config") // name of config file (without extension)
	v.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name
	// Viper settings
	v.AddConfigPath(".")
	//v.AddConfigPath("$CONFIG_DIR/")

	// Environment variable settings
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	// Global configuration
	v.SetDefault("shutdownTimeout", 15*time.Second)
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		v.SetDefault("no_color", false)
	}

	// Log configuration
	v.SetDefault("logger.format", "json")
	v.SetDefault("logger.level", "info")
	v.RegisterAlias("full_timestamp", "logger.full_timestamp") // alias for multiple ways of calling the same thing
	v.RegisterAlias("no_color", "logger.no_color")

	// App configuration
	f.String("http-addr", "8080", "App HTTP server address")
	_ = v.BindPFlag("app.httpAddr", f.Lookup("http-addr"))
	v.SetDefault("app.httpAddr", "8080")

	// Database configuration
	_ = v.BindEnv("database.host")
	v.SetDefault("database.port", 5432)
	_ = v.BindEnv("database.user")
	_ = v.BindEnv("database.password")
	_ = v.BindEnv("database.database")
	v.SetDefault("database.params", map[string]string{
		"collation": "utf8mb4_general_ci",
	})
}

func New(v *viper.Viper, f *pflag.FlagSet) Configuration {
	setDefaultConfig(v, f)
	f.String("config", "", "Configuration file")
	_ = f.Parse(os.Args[1:])

	if version, _ := f.GetBool("version"); version {
		os.Exit(0)
	}

	if c, _ := f.GetString("config"); c != "" {
		v.SetConfigFile(c)
	}
	err := v.ReadInConfig()
	if err != nil {
		logrus.Panicf("failed to read configuration: %v", err)
	}
	notFoundErr, configFileNotFound := err.(viper.ConfigFileNotFoundError)
	if configFileNotFound {
		panic(notFoundErr.Error())
	}
	var config Configuration
	err = v.Unmarshal(&config)
	return config
}
