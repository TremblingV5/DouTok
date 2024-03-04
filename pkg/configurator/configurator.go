package configurator

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/caarlos0/env/v9"
	"gopkg.in/yaml.v3"

	"github.com/TremblingV5/DouTok/pkg/constants"
	"github.com/TremblingV5/DouTok/pkg/dtviper"
)

const (
	ErrNotPtrOfStruct = "given config is not a pointer of a struct"
)

func InitConfig(t any, configName string) error {
	configPath, err := getConfigFilesPath(configName)
	if err != nil {
		return err
	}

	file, _ := os.ReadFile(configPath)
	if err := yaml.Unmarshal(file, t); err != nil {
		return err
	}

	return nil
}

func getConfigFilesPath(configName string) (string, error) {
	pathList := [5]string{
		"./config/",
		"../../config/",
		"../../../config/",
		"../../../../config/",
		"../../../../../config/",
	}

	for i := range pathList {
		_, err := os.Stat(pathList[i] + configName)
		if err == nil {
			p, _ := filepath.Abs(pathList[i] + configName)
			return p, nil
		}
	}

	return "", errors.New(constants.ErrConfigFileNotFound + ", file name: " + configName)
}

func Load(ctx context.Context, config interface{}, envPrefix, configName string) error {
	if err := loadFromEnv(ctx, config); err != nil {
		return fmt.Errorf("could not load env variables: %w", err)
	}

	v := dtviper.ConfigInit(envPrefix, configName)

	if err := loadFromFile(ctx, config, v); err != nil {
		return fmt.Errorf("could not load config from files: %w", err)
	}

	return nil
}

func loadFromEnv(ctx context.Context, config interface{}) error {
	if err := env.Parse(config); err != nil {
		return err
	}
	return nil
}

func loadFromFile(ctx context.Context, config interface{}, v *dtviper.Config) error {
	configType := reflect.TypeOf(config)

	if configType.Kind() != reflect.Ptr {
		return errors.New(ErrNotPtrOfStruct)
	}

	elemType := configType.Elem()
	if elemType.Kind() != reflect.Struct {
		return errors.New(ErrNotPtrOfStruct)
	}

	value := reflect.ValueOf(config).Elem()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := value.Type().Field(i)
		configPath := fieldType.Tag.Get("configPath")
		valueOfConfig := v.Viper.GetString(configPath)

		switch fieldType.Type.Kind() {
		case reflect.Int:
			setInt(&fieldValue, valueOfConfig)
		case reflect.Int8:
			setInt(&fieldValue, valueOfConfig)
		case reflect.Int16:
			setInt(&fieldValue, valueOfConfig)
		case reflect.Int32:
			setInt(&fieldValue, valueOfConfig)
		case reflect.Int64:
			setInt(&fieldValue, valueOfConfig)
		case reflect.Uint:
			setUint(&fieldValue, valueOfConfig)
		case reflect.Uint8:
			setUint(&fieldValue, valueOfConfig)
		case reflect.Uint16:
			setUint(&fieldValue, valueOfConfig)
		case reflect.Uint32:
			setUint(&fieldValue, valueOfConfig)
		case reflect.Uint64:
			setUint(&fieldValue, valueOfConfig)
		case reflect.Bool:
			if valueOfConfig == "true" || valueOfConfig == "True" || valueOfConfig == "TRUE" {
				fieldValue.SetBool(true)
			} else {
				fieldValue.SetBool(false)
			}
		case reflect.Struct:
			if err := loadFromFile(ctx, fieldValue.Addr().Interface(), v); err != nil {
				continue
			}
		case reflect.String:
			if valueOfConfig != "" {
				fieldValue.Set(reflect.ValueOf(valueOfConfig))
			}
		default:
			continue
		}
	}

	return nil
}

func setInt(value *reflect.Value, configPath string) {
	valueI64, err := strconv.ParseInt(configPath, 10, 64)
	if err != nil {
		return
	}
	value.SetInt(valueI64)
}

func setUint(value *reflect.Value, configPath string) {
	valueI64, err := strconv.ParseInt(configPath, 10, 64)
	if err != nil {
		return
	}
	value.SetUint(uint64(valueI64))
}
