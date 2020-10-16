package antchain

// 创建账户
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

	publicKey = string(data)
	return
}

// 存证
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

	hash = string(data)
	return
}

// 部署Solidity合约
func (c *Client) DeploySol(orderID string, contractName string, contractCode string, gas int) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"orderId":      orderID,
		"bizid":        c.BIZID,
		"account":      c.Account,
		"contractName": contractName,
		"contractCode": contractCode,
		"mykmsKeyId":   c.MykmsKeyID,
		"method":       "DEPLOYCONTRACTFORBIZ",
		"accessId":     c.AccessID,
		"token":        c.Token,
		"gas":          gas,
		"tenantid":     c.TenantID,
	}

	data, err = c.doRequest(CHAIN_CALL_FOR_BIZ, params)
	return
}

// 异步调用Solidity合约
func (c *Client) CallSol(orderID string, contractName string, methodSignature string, inputParams string, outTypes string, gas int) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"orderId":           orderID,
		"bizid":             c.BIZID,
		"account":           c.Account,
		"contractName":      contractName,
		"methodSignature":   methodSignature,
		"mykmsKeyId":        c.MykmsKeyID,
		"method":            "CALLCONTRACTBIZASYNC",
		"inputParamListStr": inputParams,
		"outTypes":          outTypes,
		"accessId":          c.AccessID,
		"token":             c.Token,
		"gas":               gas,
		"tenantid":          c.TenantID,
	}

	data, err = c.doRequest(CHAIN_CALL_FOR_BIZ, params)
	return
}
