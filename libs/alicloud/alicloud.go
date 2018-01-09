package alicloud

import (
	// "github.com/cuirixin/phoenix_corelib/utils"
	// "github.com/cuirixin/phoenix_corelib/libs/wechat/user"
	"github.com/cuirixin/phoenix_corelib/libs/alicloud/context"
	"github.com/cuirixin/phoenix_corelib/libs/alicloud/openapi"
)

// Alicloud struct
type Alicloud struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppKey         string
	AppSecret      string
	// Token          string
	// EncodingAESKey string
	// Cache          cache.Cache
}

// NewApi init
func NewAlicloud(cfg *Config) *Alicloud {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Alicloud{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppKey = cfg.AppKey
	context.AppSecret = cfg.AppSecret
	// context.Token = cfg.Token
	// context.EncodingAESKey = cfg.EncodingAESKey
	// context.Cache = cfg.Cache
}

// GetOpenApi 获取OpenApi
func (aly *Alicloud) GetOpenApi() *openapi.OpenApi {
	return openapi.NewOpenApi(aly.Context)
}
