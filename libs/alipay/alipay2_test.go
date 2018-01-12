package alipay

import (
	"testing"
	"fmt"
)

func TestAliPay(t *testing.T) {

	config := &Config{
		AppId: "",
		PartnerId: "",
		PublicKeyCertPath: "cert/alipay_demo_public_key.pem",
		PrivateKeyCertPath: "cert/alipay_demo_private_key.pem",
		AliPublicKeyCertPath: "cert/alipay_demo_public_key.pem",
		IsProduction: true,
	}
	alipay := NewAliPay(config)
	
	println(alipay.appId)
	
	var p = AliPayTradeWapPay{}
	p.NotifyURL = "xxx"
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "商品编码"

	var url, _ = client.TradeWapPay(p)

	fmt.Println("NEW URL: ", url)
}

