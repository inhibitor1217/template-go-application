package env

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	// You may want to change this prefix to avoid conflict with other env variables.
	envVariablePrefix = ""
)

func Init() {
	fillDefaultValues()

	viper.SetEnvPrefix(envVariablePrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

func Load() (*Env, error) {
	stage, err := readStage()
	if err != nil {
		return nil, errors.Wrapf(err, "viper failed to load env")
	}

	viper.SetConfigName(stage)
	viper.SetConfigType("yaml")

	configDir, err := configDir()
	if err != nil {
		return nil, errors.Wrapf(err, "viper failed to load env")
	}

	viper.AddConfigPath(configDir)

	env := &Env{}

	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, errors.Wrapf(err, "viper failed to load env")
		}
	}

	if err := bindEnvs(env); err != nil {
		return nil, errors.Wrapf(err, "viper failed to load env")
	}

	if err := viper.Unmarshal(env); err != nil {
		return nil, errors.Wrapf(err, "viper failed to load env")
	}

	return env, nil
}

func fillDefaultValues() {
	viper.SetDefault("stage", string(StageDevelopment))
}

func readStage() (Stage, error) {
	stage := viper.GetString("stage")
	switch stage {
	case string(StageDevelopment):
		return StageDevelopment, nil
	case string(StageTest):
		return StageTest, nil
	default:
		return "", errors.Errorf("invalid stage: %s", stage)
	}
}

func configDir() (string, error) {
	configAbsPath, err := filepath.Abs(".")
	if err != nil {
		return "", errors.Wrapf(err, "fail to read config path")
	}
	configAbsPath += "/config"
	return configAbsPath, nil
}

func bindEnvs(env *Env) error {
	return bindEnvToKey("", reflect.TypeOf(env))
}

func bindEnvToKey(prefix string, dataType reflect.Type) error {
	if dataType.Kind() == reflect.Ptr {
		dataType = dataType.Elem()
	}

	if dataType.Kind() == reflect.Struct {
		for i := 0; i < dataType.NumField(); i++ {
			field := dataType.Field(i)
			nextPrefix := ""
			if len(prefix) > 0 {
				nextPrefix = fmt.Sprintf("%s.%s", prefix, field.Name)
			} else {
				nextPrefix = field.Name
			}
			if err := bindEnvToKey(nextPrefix, field.Type); err != nil {
				return err
			}
		}
	} else {
		err := viper.BindEnv(prefix)
		if err != nil {
			return err
		}
	}

	return nil
}
