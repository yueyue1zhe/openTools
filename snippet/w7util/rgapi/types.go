package rgapi

type commonResult[T any] struct {
	Code    int    `json:"code"` //不为0时异常
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type AccountListResultDataItem struct {
	Name        string `json:"name"`         //号码名称
	Type        int    `json:"type"`         //号码类型
	AppID       string `json:"app_id"`       //号码APPID
	AesKey      string `json:"aes_key"`      //号码加密AESKEY
	Token       string `json:"token"`        //号码TOKEN
	AccessType  string `json:"access_type"`  //未知
	AccountType string `json:"account_type"` //未知
	LogoUrl     string `json:"logo_url"`     //Logo的URL
}

type AccessTokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
