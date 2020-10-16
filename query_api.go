package antchain

// 查询交易
func (c *Client) QueryTransaction(hash string) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"bizid":    c.BIZID,
		"method":   "QUERYTRANSACTION",
		"accessId": c.AccessID,
		"hash":     hash,
		"token":    c.Token,
	}

	data, err = c.doRequest(CHAIN_CALL, params)
	return
}

// 查询交易回执
func (c *Client) QueryReceipt(hash string) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"bizid":    c.BIZID,
		"method":   "QUERYRECEIPT",
		"accessId": c.AccessID,
		"hash":     hash,
		"token":    c.Token,
	}

	data, err = c.doRequest(CHAIN_CALL, params)
	return
}

// 查询块头
func (c *Client) QueryBlockHeader(blockNumber int64) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"bizid":      c.BIZID,
		"method":     "QUERYBLOCK",
		"accessId":   c.AccessID,
		"requestStr": blockNumber,
		"token":      c.Token,
	}

	data, err = c.doRequest(CHAIN_CALL, params)
	return
}

// 查询块体
func (c *Client) QueryBlockBody(blockNumber int64) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"bizid":      c.BIZID,
		"method":     "QUERYBLOCKBODY",
		"accessId":   c.AccessID,
		"requestStr": blockNumber,
		"token":      c.Token,
	}

	data, err = c.doRequest(CHAIN_CALL, params)
	return
}

// 查询最新块高
func (c *Client) QueryLastBlock() (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"bizid":    c.BIZID,
		"method":   "QUERYLASTBLOCK",
		"accessId": c.AccessID,
		"token":    c.Token,
	}

	data, err = c.doRequest(CHAIN_CALL, params)
	return
}

// 查询账户
func (c *Client) QueryAccount(account string) (data []byte, err error) {
	c.Shakehand()

	params := ParamsMap{
		"bizid":      c.BIZID,
		"method":     "QUERYACCOUNT",
		"requestStr": "{\"queryAccount\":\"" + account + "\"}",
		"accessId":   c.AccessID,
		"token":      c.Token,
	}

	data, err = c.doRequest(CHAIN_CALL, params)
	return
}
