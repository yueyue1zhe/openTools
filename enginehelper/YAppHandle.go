package dto

import (
	"database/sql"
	"e.coding.net/zhechat/magic/taihao/function/commutil"
	"e.coding.net/zhechat/magic/taihao/function/fileutil"
	"e.coding.net/zhechat/magic/taihao/library/dbutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type YAppHandle interface {
	Types() (req []interface{}, res []interface{}, other []interface{})
	Route(r *gin.RouterGroup)
}
type AppHandle struct {
	typesWrite         bool
	typesRequest       []interface{}
	typesRequestPrefix string
	typesResult        []interface{}
	typesResultPrefix  string
	typesOther         []interface{}
	typesOutPath       string
	typesOutNamePrefix string
}

func JustWriteType(req,res,other []interface{}) *AppHandle  {
	return &AppHandle{
		typesRequest:       req,
		typesRequestPrefix: "ApiRequest",
		typesResult:        res,
		typesOther:         other,
		typesOutPath:       ".",
		typesResultPrefix:  "ApiResult",
		typesOutNamePrefix: "",
	}
}
func LoadAppHandle(c *gin.RouterGroup, handle ...YAppHandle) *AppHandle {
	var (
		request []interface{}
		result  []interface{}
		other   []interface{}
	)
	//other = append(other, dbutil.Model{}, sql.NullTime{})
	for _, appHandle := range handle {
		appHandle.Route(c)
		tmpReq, tmpRes, tmpO := appHandle.Types()
		request = append(request, tmpReq...)
		result = append(result, tmpRes...)
		other = append(other, tmpO...)
	}
	return &AppHandle{
		typesRequest:       request,
		typesRequestPrefix: "ApiRequest",
		typesResult:        result,
		typesOther:         other,
		typesOutPath:       ".",
		typesResultPrefix:  "ApiResult",
		typesOutNamePrefix: "",
	}
}
func (a *AppHandle) AppendReqFull() *AppHandle {
	a.typesOther = append(a.typesOther, ReqPage{}, ReqID{})
	return a
}
func (a *AppHandle) AppendOtherFull() *AppHandle {
	a.typesOther = append(a.typesOther, dbutil.Model{}, sql.NullTime{})
	return a
}
func (a *AppHandle) AppendOtherInterface(mode ...interface{}) *AppHandle {
	a.typesOther = append(a.typesOther, mode...)
	return a
}
func (a *AppHandle) WriteTypesWithPower(power bool) {
	if !power {
		a.typesResult = nil
		a.typesResult = nil
		a.typesOther = nil
		return
	}
	a.WriteTypes()
}

func (a *AppHandle) SetTypesOutNamePrefix(prefix string) *AppHandle {
	a.typesOutNamePrefix = prefix
	return a
}
func (a *AppHandle) SetTypesOutPath(outPath string) *AppHandle {
	a.typesOutPath = outPath
	return a
}
func (a *AppHandle) WriteTypes() {
	a.writeDo(a.typesOutNamePrefix+"Request", a.typesRequestPrefix, a.typesRequest)
	a.writeDo(a.typesOutNamePrefix+"Result", a.typesResultPrefix, a.typesResult)
	a.writeDo(a.typesOutNamePrefix+"Other", "", a.typesOther)
}

func (a *AppHandle) writeDo(fileName, typePreName string, types []interface{}) {
	if len(types) <= 0 {
		return
	}
	if err := a.write(types, fileName, typePreName); err != nil {
		fmt.Println(fmt.Sprintf("%v types文件生成失败：%v", fileName, err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("%v types文件已生成", fileName))
}

func (a *AppHandle) write(list []interface{}, name string, typePreName string) error {
	var out []string
	for _, i := range list {
		_, raw := commutil.StructToTsTypes(i, typePreName)
		out = append(out, raw)
	}
	row := strings.Join(out, "\n\n")
	if err := os.RemoveAll(a.typesOutPath + "/" + name + ".d.ts"); err != nil {
		return err
	}
	if err := fileutil.PathMustExists(a.typesOutPath); err != nil {
		return err
	}
	return os.WriteFile(fmt.Sprintf("%v/%v.d.ts", a.typesOutPath, name), []byte(row), 0666)
}
