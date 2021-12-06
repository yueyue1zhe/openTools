package env

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"openTools/y"
	"os"
)

type Config struct {
	Http  confHttp
	Mysql confMysql
}

type confHttp struct {
	Port  string
	Debug bool
	Pid   string
}
type confMysql struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
	Pre      string
	Charset  string
}

var (
	Conf     *Config
	confPath string
)

func init() {
	flag.StringVar(&confPath, "d", y.Global().IaRoot+"/", " set ini config file path")
}

func Register() (err error) {
	Conf = NewConfig()
	viper.SetConfigName("ini")
	viper.SetConfigType("toml")
	viper.AddConfigPath(confPath)
	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	if err = viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("unable to decode into struct：  %s \n", err))
	}
	//viper.Set("http.pid", os.Getpid())
	//_ = viper.WriteConfig()
	//viper.WatchConfig()
	//viper.OnConfigChange(func(in fsnotify.Event) {
	//	fmt.Println(Conf.Mysql.Pre)
	//})
	return nil
}

func NewConfig() *Config {
	return &Config{
		Http: confHttp{
			Port:  "8089",
			Debug: false,
		},
		Mysql: confMysql{
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
	viper.AddConfigPath(confPath)
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
