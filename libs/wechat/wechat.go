package wechat

import (
	"net/http"
	"sync"

	"github.com/cuirixin/phoenix_corelib/utils/cache"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/context"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/mina"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/js"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/material"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/menu"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/oauth"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/server"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/template"
	"github.com/cuirixin/phoenix_corelib/libs/wechat/user"
)

// Wechat struct
type Wechat struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
	Cache          cache.Cache
}

// NewWechat init
func NewWechat(cfg *Config) *Wechat {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Wechat{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppID = cfg.AppID
	context.AppSecret = cfg.AppSecret
	context.Token = cfg.Token
	context.EncodingAESKey = cfg.EncodingAESKey
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
	context.SetJsAPITicketLock(new(sync.RWMutex))
}


//GetAccessToken 获取access_token
func (wc *Wechat) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetMina 小程序配置
func (wc *Wechat) GetMina() *mina.Mina {
	return mina.NewMina(wc.Context)
}

// GetJs js-sdk配置
func (wc *Wechat) GetJs() *js.Js {
	return js.NewJs(wc.Context)
}

// GetOauth oauth2网页授权
func (wc *Wechat) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Context)
}

// GetMenu 菜单管理接口
func (wc *Wechat) GetMenu() *menu.Menu {
	return menu.NewMenu(wc.Context)
}

// GetTemplate 模板消息接口
func (wc *Wechat) GetTemplate() *template.Template {
	return template.NewTemplate(wc.Context)
}

// GetMaterial 素材管理
func (wc *Wechat) GetMaterial() *material.Material {
	return material.NewMaterial(wc.Context)
}

// GetUser 用户管理接口
func (wc *Wechat) GetUser() *user.User {
	return user.NewUser(wc.Context)
}

func (wc *Wechat) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	wc.Context.Request = req
	wc.Context.Writer = writer
	return server.NewServer(wc.Context)
}

