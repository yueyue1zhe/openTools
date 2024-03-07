package bootstrap

import (
	"e.coding.net/zhechat/magic/taihao/core"
	"e.coding.net/zhechat/magic/taihao/library/dbutil"
	"e.coding.net/zhechat/magic/taihao/library/ginutil"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/url"
	"reflect"
	"time"
)

type AppModelBase struct {
	dbutil.Model
	TenantID string `json:"-" gorm:"index;type:string;size:30;not null"`
}

type Facade struct {
	TenantID string
	FacadeDBAction
}

func FacadeGetByEmpty() *Facade {
	return &Facade{}
}
func FacadeGetByTenantID(tenantID string) *Facade {
	return &Facade{
		TenantID:       tenantID,
		FacadeDBAction: FacadeDBAction{TenantID: tenantID},
	}
}
func FacadeGetByHandle(c *gin.Context) *Facade {
	tenantID := TenantIdGet(c)
	return &Facade{
		TenantID:       tenantID,
		FacadeDBAction: FacadeDBAction{TenantID: tenantID},
	}
}
func GetCurSiteDomain(c *gin.Context) string {
	appConf := core.GetAppRegisterConf()
	if appConf.Debug && appConf.DevProxySiteDomain != "" {
		return appConf.DevProxySiteDomain
	}
	return ginutil.CurSiteDomain(c)
}

func (f *Facade) EncodeCBQuery(query url.Values) string {
	appConf := core.GetAppRegisterConf()
	if appConf.MultiTenant && appConf.MultiTenantIDQueryKey != "" {
		query.Set(appConf.MultiTenantIDQueryKey, f.TenantID)
	}
	return query.Encode()
}
func (f *Facade) EncodeWechatPayCBUrl(raw string) string {
	useUrl := raw
	appConf := core.GetAppRegisterConf()
	if appConf.MultiTenant && appConf.MultiTenantIDQueryKey != "" {
		useUrl = raw + "/" + f.TenantID
	}
	return useUrl
}

func (f *Facade) cacheKeyFull(key string) string {
	return fmt.Sprintf("facade-cache-%v-%v", f.TenantID, key)
}
func (f *Facade) CacheGet(key string, val interface{}) (has bool) {
	if save, ok := dbutil.GetCache().Get(f.cacheKeyFull(key)); ok {
		use := reflect.ValueOf(val)
		use.Elem().Set(reflect.ValueOf(save))
		has = true
	}
	return
}
func (f *Facade) CacheSetWithExpire(key string, val interface{}, d time.Duration) {
	dbutil.GetCache().Set(f.cacheKeyFull(key), val, d)
}
func (f *Facade) CacheSet(key string, val interface{}) {
	dbutil.GetCache().SetDefault(f.cacheKeyFull(key), val)
}
func (f *Facade) CacheDel(key string) {
	dbutil.GetCache().Delete(f.cacheKeyFull(key))
}
func (f *Facade) FullFilePath(str string) string {
	//各账号单独配置远程附件 无所谓分割不分割
	return str
	//var out []string
	//if p.TenantID != "" {
	//	out = append(out, p.TenantID)
	//}
	//out = append(out, str)
	//return strings.Join(out, "/")
}

// Spared 使用指定db 创建新的基础facade
func (f *Facade) Spared(db *gorm.DB) *FacadeDBAction {
	return &FacadeDBAction{
		TenantID: f.TenantID,
		DB:       db,
	}
}
func (f *Facade) DBFull(db func(db *gorm.DB) *gorm.DB) *FacadeDBAction {
	return &FacadeDBAction{
		DB:       db(dbutil.Get()),
		TenantID: f.TenantID,
	}
}

type FacadeDBAction struct {
	DB       *gorm.DB
	alias    string
	TenantID string
}

// make safe when just action run
func (p *FacadeDBAction) safeDb() *gorm.DB {
	if p.DB == nil {
		p.DB = dbutil.Get()
	}
	return p.DB
}
func (p *FacadeDBAction) First(out interface{}, condS ...interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).First(&out, condS...)
}
func (p *FacadeDBAction) Last(out interface{}, condS ...interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).Last(&out, condS...)
}
func (p *FacadeDBAction) FirstOrCreate(val interface{}, condS ...interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).FirstOrCreate(val, condS...)
}
func (p *FacadeDBAction) Count(count *int64) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).Count(count)
}
func (p *FacadeDBAction) Save(val any) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantCreateScope).Save(val)
}
func (p *FacadeDBAction) Create(val any) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantCreateScope).Create(val)
}
func (p *FacadeDBAction) Find(dest any, condS ...interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).Find(dest, condS...)
}
func (p *FacadeDBAction) Updates(val interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Omit(core.GetAppRegisterConf().StructBasic).Updates(val)
}
func (p *FacadeDBAction) Update(column string, value interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).Update(column, value)
}
func (p *FacadeDBAction) Delete(val interface{}, condS ...interface{}) *gorm.DB {
	defer func() { p.DB = nil }()
	return p.safeDb().Scopes(p.tenantQueryScope).Delete(val, condS...)
}
func (p *FacadeDBAction) SetAlias(alias string) *FacadeDBAction {
	p.alias = alias
	return p
}
func (p *FacadeDBAction) tenantQueryScope(db *gorm.DB) *gorm.DB {
	useKey := core.GetAppRegisterConf().StructBasic
	if useKey == "" {
		return db
	}
	if p.alias != "" {
		useKey = fmt.Sprintf("%v.%v", p.alias, useKey)
	}
	return db.Where(useKey, p.TenantID)
}
func (p *FacadeDBAction) tenantCreateScope(db *gorm.DB) *gorm.DB {
	return db.Set(core.GetAppRegisterConf().StructBasic, p.TenantID)
}

type TokenGenerateOpt struct {
	UID         uint
	Role        string
	AuthID      uint
	AuthChannel string
	OpenID      string
}

//func (p *Facade) TokenGenerate(client string, opt TokenGenerateOpt) (string, error) {
//	return tokenGenerate(TokenPayload{
//		Uid:         opt.UID,
//		AuthID:      opt.AuthID,
//		AuthChannel: opt.AuthChannel,
//		Role:        opt.Role,
//		TenantID:    p.TenantID,
//		OpenID:      opt.OpenID,
//	}, client)
//}

// CurBaseSql	根据自身tenant_id 生成sql
func (a *AppModelBase) CurBaseSql() string {
	return getBaseSqlWithTenantID(a.TenantID)
}
func (a *AppModelBase) GetFacade() *Facade {
	return FacadeGetByTenantID(a.TenantID)
}

func getBaseSqlWithAlias(tenantID, alias string) string {
	if !core.GetAppRegisterConf().MultiTenant {
		return ""
	}
	if alias != "" {
		alias += "."
	}
	return fmt.Sprintf("%vtenant_id = %v", alias, tenantID)
}
func getBaseSqlWithTenantID(tenantID string) string {
	return getBaseSqlWithAlias(tenantID, "")
}
func GetBaseSqlWithAlias(c *gin.Context, alias string) string {
	return getBaseSqlWithAlias(TenantIdGet(c), alias)
}
func GetBaseSql(c *gin.Context) string {
	return GetBaseSqlWithAlias(c, "")
}
