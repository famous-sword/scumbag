package config

import "github.com/spf13/viper"

func String(key string) string {
	return viper.GetString(key)
}

func Integer(key string) int64 {
	return viper.GetInt64(key)
}

func Bool(key string) bool {
	return viper.GetBool(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
