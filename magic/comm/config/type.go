package config

type WxCloud struct {
	Appid  string
	Envid  string
	Host   string
	Secret string
}

type Mysql struct {
	Charset  string
	Database string
	Addr     string
	Password string
	Username string
	Pre      string
}

// Server 系统配置结构体
type Server struct {
	JwtSecret     string
	JwtIssue      string
	JwtExpireTime int32
	AesKey        string
}

// WxApi 配置结构体
type WxApi struct {
	UseCloudBaseAccessToken bool
	UseComponentAccessToken bool
	UseHttps                bool
}

// Comm 常规配置结构体
type Comm struct {
	Version string
}
