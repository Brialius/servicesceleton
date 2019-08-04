package config

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

type Config struct {
	Port          int
	ConfigFile    string
	WelcomeString string
	Settings      struct {
		key1 string
		key2 []string
	}
	LogConfig LoggerConfig
}

type LoggerConfig struct {
	Level string
	Out   io.Writer
}

func (c Config) String() string {
	return fmt.Sprintf(`Settings:
	port: "%d"
	configFile: "%s"
	welcomeString: "%s"
	config:
		key1: "%s"
		key2: "%v"
	logLevel: "%v"
`, c.Port, c.ConfigFile, c.WelcomeString, c.Settings.key1, c.Settings.key2, c.LogConfig.Level)
}

func LoadConfig(cmd *cobra.Command) *Config {
	viper.AutomaticEnv()
	_ = viper.BindPFlag("port", cmd.Flags().Lookup("port"))
	_ = viper.BindPFlag("welcomeString", cmd.Flags().Lookup("welcomeString"))
	_ = viper.BindPFlag("config", cmd.Flags().Lookup("config"))
	_ = viper.BindPFlag("verbosity", cmd.Flags().Lookup("verbosity"))
	config := viper.GetString("config")
	if config != "" {
		viper.SetConfigFile(config)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			logrus.Fatal(err)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.SetConfigName("servicesceleton")
	}

	viper.SetDefault("key1", "default")
	viper.SetDefault("key2", []string{"default1", "default2"})

	if err := viper.ReadInConfig(); err != nil {
		logrus.Warn("Can't read config: ", err)
	} else {
		viper.Set("config", viper.ConfigFileUsed())
		logrus.Warn("Read config file: ", viper.GetString("config"))
	}

	return newConfig()
}

func newConfig() *Config {
	return &Config{
		Port:          viper.GetInt("port"),
		ConfigFile:    viper.GetString("config"),
		WelcomeString: viper.GetString("welcomeString"),
		Settings: struct {
			key1 string
			key2 []string
		}{
			key1: viper.GetString("key1"),
			key2: viper.GetStringSlice("key2"),
		},
		LogConfig: LoggerConfig{
			Out:   os.Stderr,
			Level: viper.GetString("verbosity"),
		},
	}
}

func ConfigureLogger(config *LoggerConfig) (*logrus.Logger, error) {
	if config.Out != nil {
		logrus.SetOutput(config.Out)
	}

	if config.Level != "" {
		level, err := logrus.ParseLevel(config.Level)
		if err != nil {
			return nil, err
		}
		logrus.SetLevel(level)
	}

	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    false,
		FieldsOrder: []string{"component", "category"},
	})

	return logrus.StandardLogger(), nil
}
