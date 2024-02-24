package envutil

import (
	"github.com/spf13/viper"
	"reflect"
)

func WriteIniToml[T any](conf T) error {
	t := reflect.TypeOf(conf)
	v := reflect.ValueOf(conf)
	for k := 0; k < t.NumField(); k++ {
		label := t.Field(k).Tag.Get("mapstructure")
		if !v.Field(k).IsZero() {
			switch label {
			case "DEBUG":
				viper.Set(label, v.Field(k).Bool())
			default:
				viper.Set(label, v.Field(k).String())
			}
		}
	}
	viper.SetConfigFile("./ini.toml")
	return viper.WriteConfig()
}
