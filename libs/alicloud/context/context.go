package context

import (
	"fmt"
	"crypto/sha1"
	"crypto/hmac"
	"encoding/base64"
	"net/url"
	"sort"
	"strings"
	"bytes"
)

// Context struct
type Context struct {
	AppKey         string
	AppSecret      string
	// Token          string
	// EncodingAESKey string

	// Cache cache.Cache

	// Writer  http.ResponseWriter
	// Request *http.Request
}

// 签名的字符串
func (ctx *Context) GetSignature(stringForSign string) string {
	secreat := []byte(ctx.AppSecret)
	mac := hmac.New(sha1.New, secreat)
	fmt.Println("[alicloud stringForSign]:", stringForSign)
	mac.Write([]byte(stringForSign))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

// 
func (ctx *Context) MakeParams(params url.Values, appKey string) (params_str, sign_str string) {
	var s, p string
	var keys []string
	b := bytes.Buffer{}
	b.WriteString(appKey)
	for k, _ := range params {
			if k != "sign" {
					keys = append(keys, k)
			}
	}
	sort.Strings(keys)
	for _, v := range keys {
			b.WriteString(v)
			b.WriteString(params.Get(v))
	}
	p = b.String()
	b.WriteString(appKey)
	s = b.String()
	p = strings.TrimRight(p, "&")
	return p, s
}
