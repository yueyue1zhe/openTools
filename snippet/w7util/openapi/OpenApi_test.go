package openapi

import (
	"fmt"
	"testing"
)

func TestW7OauthGetLoginUrl(t *testing.T) {
	r, e := GetClient("347114", "db8591d328e007414d011535e83b1766").GetLoginUrl("http://www.baidu.com")
	fmt.Println(r)
	fmt.Println(e)
}

func TestW7OauthCodeToAccessToken(t *testing.T) {
	r, e := GetClient("347114", "db8591d328e007414d011535e83b1766").CodeToAccessToken("ygWjyJUrDHdJJ60JpWBgWGUdphdQ0gbD")
	fmt.Println(r)
	fmt.Println(e)
}

func TestW7OauthUserInfo(t *testing.T) {
	r, e := GetClient("347114", "db8591d328e007414d011535e83b1766").UserInfo("rxxj3C3cLlzII7YJglb7LxjbjJI33C47l447csGBsgG8c7g7ssJ73GsCY47zbxs7")
	fmt.Println(r)
	fmt.Println(e)
}
