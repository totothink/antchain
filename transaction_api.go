package antchain

import "github.com/tidwall/gjson"

func (c *Client) CreateAccount(orderID string, newAccount string, newKmsID string) (publicKey string, err error) {
	c.Shakehand()

	params := ParamsMap{
		"orderId":         orderID,
		"bizid":           c.BIZID,
		"method":          "TENANTCREATEACCUNT",
		"account":         c.Account,
		"mykmsKeyId":      c.MykmsKeyID,
		"newAccountId":    newAccount,
		"newAccountKmsId": newKmsID,
		"accessId":        c.AccessID,
		"token":           c.Token,
		"tenantid":        c.TenantID,
	}

	data, err := c.doRequest(CHAIN_CALL_FOR_BIZ, params)
	if err != nil {
		return
	}

	result := gjson.Get(string(data), "success")
	if result.Bool() == false {
		err = &AccessError{
			message: gjson.Get(string(data), "data").String(),
		}
		return
	}

	publicKey = gjson.Get(string(data), "data").String()
	return
}

func (c *Client) Deposit(orderID string, content string, gas int) (hash string, err error) {
	c.Shakehand()

	params := ParamsMap{
		"orderId":    orderID,
		"bizid":      c.BIZID,
		"account":    c.Account,
		"content":    content,
		"mykmsKeyId": c.MykmsKeyID,
		"method":     "DEPOSIT",
		"accessId":   c.AccessID,
		"token":      c.Token,
		"gas":        gas,
		"tenantid":   c.TenantID,
	}

	data, err := c.doRequest(CHAIN_CALL_FOR_BIZ, params)

	result := gjson.Get(string(data), "success")
	if result.Bool() == false {
		err = &AccessError{
			message: gjson.Get(string(data), "data").String(),
		}
		return
	}

	hash = gjson.Get(string(data), "data").String()
	return
}
