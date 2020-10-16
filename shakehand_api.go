package antchain

import (
	"strconv"
	"time"
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

	token, err := c.doRequest(SHAKE_HAND, params)
	if err != nil {
		return
	}

	c.Token = string(token)
	c.TokenExpiresAt = timeUnix + 1800
	return
}
