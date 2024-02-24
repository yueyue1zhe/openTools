package envutil

import (
	"e.coding.net/zhechat/magic/taihao/core"
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"fmt"
	"github.com/spf13/viper"
)

func LoadGlobalEnv[T any](env *T) error {
	var structEmpty T
	envList := commutil.StructToTagLabelSlice(structEmpty, "mapstructure")
	for _, targetEnv := range envList {
		if err := viper.BindEnv(targetEnv); err != nil {
			return fmt.Errorf("bind env %v err:%v", targetEnv, err.Error())
		}
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(env); err != nil {
		return fmt.Errorf("global env parse error :%v", err.Error())
	}
	return commutil.StructRequiredJudge(*env)
}

func LoadFileEnv[T any](env *T) error {
	viper.SetConfigFile(core.IaRoot() + "/ini.toml")
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("未找到配置文件:%v", err.Error())
	}
	if err := viper.Unmarshal(env); err != nil {
		return fmt.Errorf("配置解析异常:%v", err.Error())
	}
	return commutil.StructRequiredJudge(*env)
}
