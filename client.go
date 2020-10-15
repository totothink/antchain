package antchain

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/spacemonkeygo/openssl"
)

type Config struct {
	Endpoint        string // 请求地址
	TenantID        string // 租户ID
	BIZID           string // 链ID
	AccessID        string // 访问ID
	AccessKey       string // 访问密钥
	Account         string // 链账户
	MykmsKeyID      string // 托管标识
	ChainCallForBiz string // 商业接口调用路径
	ChainCall       string // 其他接口调用路径
}

type Client struct {
	Config
	Token          string // 访问token
	TokenExpiresAt int64  // 令牌过期时间

	httpClient    *http.Client
	AppPrivateKey openssl.PrivateKey
}

// Client 发送请求使用的客户端
func NewClient(config Config) (client Client, err error) {
	if config.Endpoint == "" {
		config.Endpoint = ENDPOINT
	}

	if config.BIZID == "" {
		config.BIZID = BIZID
	}

	if config.ChainCallForBiz == "" {
		config.ChainCallForBiz = CHAIN_CALL_FOR_BIZ
	}

	if config.ChainCall == "" {
		config.ChainCall = CHAIN_CALL
	}

	appPrivateKey, err := openssl.LoadPrivateKeyFromPEM([]byte(config.AccessKey))

	if err != nil {
		fmt.Print(err)
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
				DualStack: true,
			}).DialContext,
		},
	}

	return Client{
		Config:        config,
		httpClient:    httpClient,
		AppPrivateKey: appPrivateKey,
	}, nil
}
