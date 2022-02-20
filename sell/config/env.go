package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"yueyue-sell/y"
)

type envConf struct {
	Http  http
	Mysql mysql
}

type http struct {
	Port  string
	Debug bool
	Pid   string
}
type mysql struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Pre      string
	Charset  string
}

var (
	Env     *envConf
	envPath string
)

func init() {
	flag.StringVar(&envPath, "d", y.Global().IaRoot+"/", " set ini config file path")
}

func registerEnv() (err error) {
	Env = newConfig()
	viper.SetConfigName("ini")
	viper.SetConfigType("toml")
	viper.AddConfigPath(envPath)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}
	if err = viper.Unmarshal(&Env); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}
	return nil
}

func newConfig() *envConf {
	return &envConf{
		Http: http{
			Port:  "8088",
			Debug: false,
		},
		Mysql: mysql{
			Host:    "127.0.0.1",
			Port:    "3306",
			Pre:     "y_",
			Charset: "utf8mb4",
		},
	}
}

func InstallSet(host, port, database, username, password, pre string) error {
	f, err := os.Create(y.Global().IaRoot + "/ini.toml")
	if err != nil {
		return err
	}
	f.Close()
	viper.SetConfigName("ini")
	viper.SetConfigType("toml")
	viper.AddConfigPath(envPath)
	viper.Set("mysql.host", host)
	viper.Set("mysql.port", port)
	viper.Set("mysql.charset", "utf8mb4")
	viper.Set("mysql.database", database)
	viper.Set("mysql.username", username)
	viper.Set("mysql.password", password)
	viper.Set("mysql.pre", pre)
	_ = viper.WriteConfig()
	return nil
}
