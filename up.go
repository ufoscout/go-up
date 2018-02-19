package go_up

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ufoscout/go-up/reader"
	"github.com/ufoscout/go-up/reader/decorator"
)

type GoUp interface {
	Exists(key string) bool
	GetBool(key string) bool
	GetBoolOrDefault(key string, defaultValue bool) bool
	GetBoolOrFail(key string) (bool, error)
	GetFloat64(key string) float64
	GetFloat64OrDefault(key string, defaultValue float64) float64
	GetFloat64OrFail(key string) (float64, error)
	GetInt(key string) int
	GetIntOrDefault(key string, defaultValue int) int
	GetIntOrFail(key string) (int, error)
	GetString(key string) string
	GetStringOrDefault(key string, defaultValue string) string
	GetStringOrFail(key string) (string, error)
	GetStringSlice(key string, separator string) []string
	GetStringSliceOrDefault(key string, separator string, defaultValue []string) []string
	GetStringSliceOrFail(key string, separator string) ([]string, error)
}

type goUpImpl struct {
	properties map[string]reader.Property
}

func (up *goUpImpl) Exists(key string) bool {
	_, found := up.properties[key]
	return found
}

func (up *goUpImpl) GetBool(key string) bool {
	result, _ := up.GetBoolOrFail(key)
	return result
}

func (up *goUpImpl) GetBoolOrDefault(key string, defaultValue bool) bool {
	result, err := up.GetBoolOrFail(key)
	if err != nil {
		result = defaultValue
	}
	return result
}

func (up *goUpImpl) GetBoolOrFail(key string) (bool, error) {
	var result bool
	value, errfound := up.GetStringOrFail(key)
	if errfound == nil {
		converted, err := strconv.ParseBool(value)
		if err != nil {
			return result, err
		}
		return converted, nil
	}
	return result, errfound
}

func (up *goUpImpl) GetFloat64(key string) float64 {
	result, _ := up.GetFloat64OrFail(key)
	return result
}

func (up *goUpImpl) GetFloat64OrDefault(key string, defaultValue float64) float64 {
	result, err := up.GetFloat64OrFail(key)
	if err != nil {
		result = defaultValue
	}
	return result
}

func (up *goUpImpl) GetFloat64OrFail(key string) (float64, error) {
	var result float64
	value, errfound := up.GetStringOrFail(key)
	if errfound == nil {
		converted, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return result, err
		}
		return converted, nil
	}
	return result, errfound
}

func (up *goUpImpl) GetInt(key string) int {
	result, _ := up.GetIntOrFail(key)
	return result
}

func (up *goUpImpl) GetIntOrDefault(key string, defaultValue int) int {
	result, err := up.GetIntOrFail(key)
	if err != nil {
		result = defaultValue
	}
	return result
}

func (up *goUpImpl) GetIntOrFail(key string) (int, error) {
	var result int
	value, errfound := up.GetStringOrFail(key)
	if errfound == nil {
		converted, err := strconv.Atoi(value)
		if err != nil {
			return result, err
		}
		return converted, nil
	}
	return result, errfound
}

func (up *goUpImpl) GetString(key string) string {
	value, _ := up.properties[key]
	return value.Value
}

func (up *goUpImpl) GetStringOrDefault(key string, defaultValue string) string {
	if up.Exists(key) {
		return up.GetString(key)
	}
	return defaultValue
}

func (up *goUpImpl) GetStringOrFail(key string) (string, error) {
	var result string
	value, found := up.properties[key]
	if found {
		return value.Value, nil
	}
	return result, fmt.Errorf("Key [%s] not found", key)
}

func (up *goUpImpl) GetStringSlice(key string, separator string) []string {
	result, err := up.GetStringSliceOrFail(key, separator);
	if err!=nil {
		return []string{}
	}
	return result
}

func (up *goUpImpl) GetStringSliceOrDefault(key string, separator string, defaultValue []string) []string {
	result, err := up.GetStringSliceOrFail(key, separator);
	if err!=nil {
		return defaultValue
	}
	return result
}

func (up *goUpImpl) GetStringSliceOrFail(key string, separator string) ([]string, error) {
	result, err := up.GetStringOrFail(key);
	if err!=nil {
		return nil, err
	}
	return strings.Split(result, separator), nil
}

func NewEnvReader(prefix string, toLower bool, underscoreToDot bool) reader.Reader {
	var envReader reader.Reader = &reader.EnvReader{prefix}
	if toLower {
		envReader = &decorator.ToLowerCaseKeyDecoratorReader{envReader}
	}
	if underscoreToDot {
		envReader = &decorator.KeyStringReplacerDecoratorReader{envReader, "_", "."}
	}
	return envReader
}
