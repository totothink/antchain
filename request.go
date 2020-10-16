package antchain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

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

	result := gjson.Get(string(data), "success")
	if result.Bool() == false {
		err = &AccessError{
			message: gjson.Get(string(data), "data").String(),
		}
		return
	}

	data = []byte(gjson.Get(string(data), "data").String())
	return
}
