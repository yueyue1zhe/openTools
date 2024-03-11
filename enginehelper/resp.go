package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AdminAuthBaseRes struct {
	TenantID string   `json:"tenant_id,omitempty" y-ts-types-spacer:"?:"`
	Domain   []string `json:"domain,omitempty"  y-ts-types-spacer:"?:" y-ts-types:"string[]"`

	HomePage string `json:"home_page,omitempty" y-ts-types-spacer:"?:"`
}

const (
	ErrorToast    = 1
	ErrorRedirect = 2

	ErrorInvalidParam = 1001 //参数异常
	ErrorSystem       = 1002 //系统错误 taiyi bridge中使用

	ErrorTenantIdFail = 40010 //租户编号异常
	ErrorAuthFail     = 40019 //未登录
	ErrorPowerFail    = 40020 //权限不足

	ErrorProxyUri = 50001 //代理uri异常
	ErrorProxyReq = 50002 //代理请求异常
	ErrorProxyOwn = 50003 //代理服务自身异常

	ErrorServiceFail    = 60000
	ErrorRecordNotfound = 60004 //记录不存在

)

// Result 返回结果
type Result interface {
	SetData(data interface{}) Result
	AppendMsg(msg string) Result
	SetErrno(errno int64) Result
	HttpOk(c *gin.Context)
}

type JsonResult struct {
	Errno   int64       `json:"errno"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
type JsonResultCommon[T any] struct {
	Errno   int64  `json:"errno"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func (j *JsonResult) SetData(data interface{}) Result {
	j.Data = data
	return j
}
func (j *JsonResult) AppendMsg(msg string) Result {
	if j.Message != "" {
		j.Message = strings.Join([]string{j.Message, msg}, "｜")
	} else {
		j.Message = msg
	}
	return j
}
func (j *JsonResult) SetErrno(errno int64) Result {
	j.Errno = errno
	return j
}
func (j *JsonResult) HttpOk(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, j)
}
func JsonOk() Result {
	return &JsonResult{}
}

func JsonOkData(data interface{}) Result {
	return &JsonResult{Data: data}
}

func JsonToast(message string, appendMsg ...string) Result {
	r := JsonOk().SetErrno(ErrorToast).AppendMsg(message)
	for _, s := range appendMsg {
		r.AppendMsg(s)
	}
	return r
}
func JsonRedirect(tips, path string) Result {
	return JsonOk().SetErrno(ErrorRedirect).AppendMsg(tips).SetData(path)
}
func JsonErrorInvalidParam() Result {
	return JsonOk().SetErrno(ErrorInvalidParam).AppendMsg("参数异常")
}
func JsonErrorSystem() Result {
	return JsonOk().SetErrno(ErrorSystem).AppendMsg("系统错误")
}
func JsonErrorTenantIDFail() Result {
	return JsonOk().SetErrno(ErrorTenantIdFail).AppendMsg("t 信息 no 异常")
}
func JsonErrorAuthFail() Result {
	return JsonOk().SetErrno(ErrorAuthFail).AppendMsg("登录状态异常")
}
func JsonErrorPowerFail() Result {
	return JsonOk().SetErrno(ErrorPowerFail).AppendMsg("权限不足")
}
func JsonErrorProxyUri() Result {
	return JsonOk().SetErrno(ErrorProxyUri).AppendMsg("a u 异常")
}
func JsonErrorProxyReq() Result {
	return JsonOk().SetErrno(ErrorProxyReq).AppendMsg("a 请求异常")
}
func JsonErrorProxyOwn() Result {
	return JsonOk().SetErrno(ErrorProxyOwn).AppendMsg("a 服务异常")
}
func JsonErrorServiceFail() Result {
	return JsonOk().SetErrno(ErrorServiceFail).AppendMsg("业务服务异常")
}
func JsonErrorRecordNotfound() Result {
	return JsonOk().SetErrno(ErrorRecordNotfound).AppendMsg("数据记录不存在")
}
