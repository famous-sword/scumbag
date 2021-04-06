package config

import "github.com/spf13/viper"

func String(key string) string {
	return viper.GetString(key)
}

func Integer(key string) int {
	return viper.GetInt(key)
}

func Bool(key string) bool {
	return viper.GetBool(key)
}

func Int32(key string) int32 {
	return viper.GetInt32(key)
}

func Int64(key string) int64 {
	return viper.GetInt64(key)
}

func Uint(key string) uint {
	return viper.GetUint(key)
}

func Uint32(key string) uint32 {
	return viper.GetUint32(key)
}

func Uint64(key string) uint64 {
	return viper.GetUint64(key)
}

func Float(key string) float64 {
	return viper.GetFloat64(key)
}

func Set(key string, value interface{}) {
	viper.Set(key, value)
}
