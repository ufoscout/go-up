package go_up

import (
	"github.com/ufoscout/go-up/reader"
	"strconv"
	"strings"
	"github.com/ufoscout/go-up/reader/decorator"
)

type GoUp interface {
	Exists(key string) bool
	GetBool(key string) bool
	GetBoolOrDefault(key string, defaultValue bool) bool
	GetFloat64(key string) float64
	GetFloat64OrDefault(key string, defaultValue float64) float64
	GetInt(key string) int
	GetIntOrDefault(key string, defaultValue int) int
	GetString(key string) string
	GetStringOrDefault(key string, defaultValue string) string
}

type goUpImpl struct {
	properties map[string]reader.Property
}

func (up *goUpImpl) Exists(key string) bool {
	_, found := up.properties[key]
	return found
}

func (up *goUpImpl) GetBool(key string) bool {
	var result bool
	value, found := up.properties[key]
	if !found {
		converted, err := strconv.ParseBool(value.Value)
		if err==nil {
			result = converted
		}
	}
	return result
}


func (up *goUpImpl) GetBoolOrDefault(key string, defaultValue bool) bool {
	if up.Exists(key) {
		return up.GetBool(key)
	}
	return defaultValue
}


func (up *goUpImpl) GetFloat64(key string) float64 {
	var result float64
	value, found := up.properties[key]
	if !found {
		converted, err := strconv.ParseFloat(value.Value, 64)
		if err==nil {
			result = converted
		}
	}
	return result
}


func (up *goUpImpl) GetFloat64OrDefault(key string, defaultValue float64) float64 {
	if up.Exists(key) {
		return up.GetFloat64(key)
	}
	return defaultValue
}


func (up *goUpImpl) GetInt(key string) int {
	var result int
	value, found := up.properties[key]
	if !found {
		converted, err := strconv.Atoi(value.Value)
		if err==nil {
			result = converted
		}
	}
	return result
}


func (up *goUpImpl) GetIntOrDefault(key string, defaultValue int) int {
	if up.Exists(key) {
		return up.GetInt(key)
	}
	return defaultValue
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


func (up *goUpImpl) GetStringSlice(key string, separator string) []string {
	return strings.Split(up.GetString(key), separator)
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