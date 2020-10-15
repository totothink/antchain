package antchain

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/spacemonkeygo/openssl"
)

// 获取参数签名
func (c *Client) getParamsSign(paramsMap ParamsMap) (sign string, err error) {

	// 拼接原始串
	signStr := sortSignParams(paramsMap)

	// 调用字符串签名
	sign, err = c.getSign(signStr)
	return
}

// 获取字符串签名
func (c *Client) getSign(message string) (sign string, err error) {
	encryptedBytes, err := c.AppPrivateKey.SignPKCS1v15(openssl.SHA256_Method, []byte(message))

	// hex转码
	sign = hex.EncodeToString(encryptedBytes)
	return
}

func (c *Client) GetSign(message string) (sign string, err error) {
	sign, err = c.getSign(message)
	return
}

// 获取根据Key排序后的请求参数字符串
func sortSignParams(paramsMap ParamsMap) string {
	keyList := make([]string, 0)
	for k := range paramsMap {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		buffer.WriteString(fmt.Sprintf("%s=%v&", k, paramsMap[k]))
	}
	s, i := buffer.String(), buffer.Len()
	return s[:i-1]
}
