package antchain

import (
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

func (c *Client) Shakehand() (err error) {
	err = c.getToken()
	if err != nil {
		panic(err)
	}
	return
}

func (c *Client) getToken() (err error) {
	timeUnix := time.Now().Unix()
	timeStr := strconv.FormatInt(timeUnix, 10) + "000"
	tnonce := c.AccessID + timeStr
	signature, err := c.getSign(tnonce)
	if err != nil {
		return
	}

	params := ParamsMap{
		"accessId": c.AccessID,
		"time":     timeStr,
		"secret":   signature,
	}

	bytes, err := c.doRequest(SHAKE_HAND, params)
	if err != nil {
		return
	}

	result := gjson.Get(string(bytes), "success")
	if result.Bool() == false {
		err = &AccessError{
			message: gjson.Get(string(bytes), "data").String(),
		}
		return
	}

	token := gjson.Get(string(bytes), "data")
	c.Token = token.String()
	c.TokenExpiresAt = timeUnix + 1800
	return
}
