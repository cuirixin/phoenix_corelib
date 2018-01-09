package openapi

import (
	"github.com/cuirixin/phoenix_corelib/libs/alicloud/context"
)

// OpenApi struct
type OpenApi struct {
	*context.Context
}

func NewOpenApi(context *context.Context) *OpenApi {
	api := new(OpenApi)
	api.Context = context
	return api
}

// func (api *OpenApi) GenRequestParams(string) 
