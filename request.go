package antchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CommonResponse 公共响应参数
type CommonResponse struct {
	Code    string `json:"code"`               // 网关返回码，参见https://docs.open.alipay.com/common/105806
	Msg     string `json:"msg"`                // 网关返回码描述，参见https://docs.open.alipay.com/common/105806
	SubCode string `json:"sub_code,omitempty"` // 业务返回码，参见具体的API接口文档
	SubMsg  string `json:"sub_msg,omitempty"`  // 业务返回码描述，参见具体的API接口文档
}

// ParamsMap 参数 Map
type ParamsMap map[string]interface{}

func convertToParamsMap(params interface{}) (paramsMap ParamsMap) {
	paramStr, _ := json.Marshal(params)
	_ = json.Unmarshal(paramStr, &paramsMap)
	return
}

// 生成 biz_content 业务字段
func generateBizContent(body interface{}) string {
	bodyStr, _ := json.Marshal(body)
	return string(bodyStr)
}

// 发送Post请求
func (c *Client) doRequest(path string, paramsMap ParamsMap) (data []byte, err error) {

	// 发起请求
	var (
		resp *http.Response
	)

	paramsJson, err := json.Marshal(paramsMap)
	if err != nil {
		return
	}
	fullPath := c.Endpoint + path

	fmt.Printf("doRequest path=%v param=%v\n", fullPath, string(paramsJson))

	req, err := http.NewRequest("POST", fullPath, bytes.NewReader(paramsJson))
	req.Header.Set("Content-Type", "application/json")

	resp, err = c.httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	data, _ = ioutil.ReadAll(resp.Body)
	fmt.Printf("doRequest path=%v got response=%v\n", fullPath, string(data))
	return
}
