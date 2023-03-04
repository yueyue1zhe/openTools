package rgapi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRgApi_GetAccountList(t *testing.T) {
	c := GetClient("w7i2n1xrjttrtdgpgd", "4d7495fd274a7685f435eec76edf6560")
	list, err := c.GetAccountList()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	js, _ := json.Marshal(&list)
	fmt.Println(string(js))
}

func TestRgApi_GetAccessToken(t *testing.T) {
	c := GetClient("w7i2n1xrjttrtdgpgd", "4d7495fd274a7685f435eec76edf6560")
	c.SetUseAccountType(AccountTypeMiniProgram)
	r, e := c.GetAccessToken(true)
	fmt.Println(r, e)
}

func TestRgApi_PayTransactionsJsapi(t *testing.T) {
	c := GetClient("w7i2n1xrjttrtdgpgd", "4d7495fd274a7685f435eec76edf6560")
	c.SetUseAccountType(AccountTypeMiniProgram)
	c.PayTransactionsJsapi()
}
