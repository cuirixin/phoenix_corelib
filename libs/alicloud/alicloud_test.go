package alicloud

import (
	"fmt"
	"testing"
)

func TestAliCloudOpenApi(t *testing.T) {

	//配置微信参数
	config := &Config{
		AppKey:         "23484070",
		AppSecret:      "c2c9b3b0311fe9bb6aa313ffacb09999",
	}
		
	aly := NewAlicloud(config)
	api := aly.GetOpenApi()

	fmt.Println(api, api.AppKey, api.AppSecret)

	fmt.Println(api.GetSignature("22321"))

}