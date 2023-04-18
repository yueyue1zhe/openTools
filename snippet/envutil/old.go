package envutil

/*
此方式已废弃




//变量加载次序
//命令行
//环境变量
//本地文件

//参考素材
//https://zhuanlan.zhihu.com/p/272508571
//https://blog.csdn.net/Guzarish/article/details/118626803


type EnvConf struct {
	Debug bool `json:"debug" mapstructure:"DEBUG"` //默认false

	AppPort string   `json:"APP_PORT" mapstructure:"APP_PORT"` //默认8080 云托管时无需配置
	AppCors []string `json:"APP_CORS" mapstructure:"APP_CORS"` //可选配置 跨域数组 需携带http协议

	//mysql 文件读取时 需自行判断 mysql 配置是否加载
	MysqlAddress  string `json:"MYSQL_ADDRESS" mapstructure:"MYSQL_ADDRESS"` //含端口号
	MysqlUsername string `json:"MYSQL_USERNAME" mapstructure:"MYSQL_USERNAME"`
	MysqlPassword string `json:"MYSQL_PASSWORD" mapstructure:"MYSQL_PASSWORD"`
	MysqlDatabase string `json:"MYSQL_DATABASE" mapstructure:"MYSQL_DATABASE"` //程序需自动判断数据库是否已新建
	//utf8mb4  collation utf8mb4_general_ci

	RedisAddress  string `json:"REDIS_ADDRESS" mapstructure:"REDIS_ADDRESS"` //redis地址(程序判断是否有这个变量)
	RedisPort     string `json:"REDIS_PORT" mapstructure:"REDIS_PORT"`
	RedisPassword string `json:"REDIS_PASSWORD" mapstructure:"REDIS_PASSWORD"`

	//微擎特殊参数
	W7AppID     string `json:"w7_app_id" mapstructure:"APP_ID"`
	W7AppSecret string `json:"w7_app_secret" mapstructure:"APP_SECRET"`

	FUllBy int `json:"full_by" mapstructure:"FULL_BY"`

	Mode string `json:"mode" mapstructure:"MODE"` //self w7 wexin 暂未启用
}

func (c *EnvConf) WriteIniToml(conf EnvConf) error {
	t := reflect.TypeOf(conf)
	v := reflect.ValueOf(conf)
	for k := 0; k < t.NumField(); k++ {
		label := t.Field(k).Tag.Get("mapstructure")
		if !v.Field(k).IsZero() {
			switch label {
			case "DEBUG":
				viper.Set(label, v.Field(k).Bool())
			case "FULL_BY":
				viper.Set(label, v.Field(k).Int())
			default:
				viper.Set(label, v.Field(k).String())
			}
		}
	}
	viper.SetConfigFile("./ini.toml")
	return viper.WriteConfig()
}

const (
	KeyDebug   = "DEBUG"
	KeyAppPort = "APP_PORT"

	DefaultAppPort = "8080"
)
const (
	FullWait = iota + 1
	FullByFile
	FullByPFlag
	FullByGlobalEnv
)

var envConfig *EnvConf

func GetEnvConfig() *EnvConf {
	if envConfig == nil {
		envConfig = &EnvConf{}
	}
	return envConfig
}

// Init 加载 本地文件 命令行 环境变量 配置参数
func Init(defaultPort string) error {
	viper.SetDefault(KeyDebug, false)
	appPort := defaultPort
	if appPort == "" {
		appPort = DefaultAppPort
	}
	viper.SetDefault(KeyAppPort, appPort)
	GetEnvConfig()
	op := optutil.FuncErrOptionInclude(fileEnvLoad, globalEnvLoad, pFlagEnvLoad)
	for _, option := range op {
		if err := option(); err == nil {
			break
		}
	}

	if envConfig.FUllBy == FullWait {
		infoF := "启动配置加载失败"
		logutil.InfoF(infoF)
		return fmt.Errorf(infoF)
	}
	return nil
}




func globalEnvLoad() error {
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&envConfig); err != nil {
		return err
	}
	envConfig.FUllBy = FullByGlobalEnv
	return nil
}


func pFlagEnvLoad() error {
	pflag.Bool("DEBUG", false, "")
	pflag.String("APP_HOST", "", "")
	pflag.String("APP_PORT", "", "")
	pflag.StringArray("APP_CORS", []string{}, "")

	pflag.String("MYSQL_ADDRESS", "", "")
	pflag.String("MYSQL_USERNAME", "", "")
	pflag.String("MYSQL_PASSWORD", "", "")
	pflag.String("MYSQL_DATABASE", "", "")

	pflag.String("REDIS_ADDRESS", "", "")
	pflag.String("REDIS_PORT", "", "")
	pflag.String("REDIS_PASSWORD", "", "")

	pflag.String("APP_ID", "", "")
	pflag.String("APP_SECRET", "", "")

	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		return err
	}
	if err := viper.Unmarshal(&envConfig); err != nil {
		return err
	}
	envConfig.FUllBy = FullByPFlag
	return nil
}


func fileEnvLoad() error {
	viper.SetConfigFile(core.IaRoot() + "/ini.toml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if err := viper.Unmarshal(&envConfig); err != nil {
		return err
	}
	envConfig.FUllBy = FullByFile
	return nil
}


*/
