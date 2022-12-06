package envutil

import (
	"e.coding.net/zhechat/magic/jarvis/core"
	"e.coding.net/zhechat/magic/jarvis/function/structutil"
	"fmt"
	"github.com/spf13/viper"
)

func LoadGlobalEnv[T any](env *T) error {
	var structEmpty T
	envList := structutil.ToTagLabelSlice(structEmpty, "mapstructure")
	for _, targetEnv := range envList {
		if err := viper.BindEnv(targetEnv); err != nil {
			return fmt.Errorf("bind env %v err:%v", targetEnv, err.Error())
		}
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&env); err != nil {
		return fmt.Errorf("global env parse error :%v", err.Error())
	}
	return structutil.RequiredJudge(*env)
}

func LoadFileEnv[T any](env *T) error {
	viper.SetConfigFile(core.IaRoot() + "/ini.toml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("未找到配置文件:%v", err.Error())
	}
	if err := viper.Unmarshal(&env); err != nil {
		return fmt.Errorf("配置解析异常:%v", err.Error())
	}
	return structutil.RequiredJudge(*env)
}
