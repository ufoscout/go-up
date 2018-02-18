package go_up

import "github.com/ufoscout/go-up/reader"

type GoUp interface {
	GetOrDefault(key string, defaultValue string) string
	Get(key string) *string
}

type goUpImpl struct {
	properties map[string]reader.Property
}

func (up *goUpImpl) GetOrDefault(key string, defaultValue string) string {
	value, found := up.properties[key]
	if !found {
		return defaultValue
	}
	return value.Value
}

func (up *goUpImpl) Get(key string) *string {
	value, found := up.properties[key]
	if !found {
		return nil
	}
	return &value.Value
}