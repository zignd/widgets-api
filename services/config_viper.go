package services

import (
	"os"
	"path"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// ConfigViper configures the application flags.
func ConfigViper() error {
	// This application conforms with the XDG Base Directory Specification
	// https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html

	viper.SetConfigName("config")
	if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
		viper.AddConfigPath(path.Join(xdgConfigHome, "widget_api"))
	} else {
		viper.AddConfigPath("$HOME/.config/widget_api")
	}

	if err := viper.ReadInConfig(); err != nil {
		return errors.Wrap(err, "failed to read the configuration file (named config)")
	}

	if viper.GetInt("port") == 0 {
		return errors.New("port is required")
	}

	if len(viper.GetString("postgresql_connection_string")) == 0 {
		return errors.New("postgresql_connection_string is required")
	}

	return nil
}
