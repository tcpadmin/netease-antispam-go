package core

import (
	"net/http"
	"time"

	"github.com/tcpadmin/netease-antispam-go/common"
)

type Config struct {
	secretID  string // 必填 产品秘钥 id,由易盾内容安全服务分配,产品标识
	secretKey string // 必填 key
	log       common.ILogger
	client    *http.Client
}

func (c *Config) SetLogLevel(level common.LogLevel) {
	c.log.SetLogLevel(level)
}

type ConfigOption func(*Config)

// NewConfig 初始化一个内容安全的 client
func NewConfig(secretId, secretKey string, opts ...ConfigOption) *Config {
	c := &Config{
		secretID:  secretId,
		secretKey: secretKey,
		client:    &http.Client{Timeout: 3 * time.Second}, //默认3s超时
	}
	for i := range opts {
		opts[i](c)
	}
	if c.log == nil {
		log := &common.DefaultLog{}
		log.SetLogLevel(common.Silent)
		c.log = log
	}
	return c
}

// WithLog 设置 client 的logger
func WithLog(l common.ILogger) ConfigOption {
	return func(c *Config) {
		c.log = l
	}
}

// WithTimeout 设置 httpClient 的 超时时间
func WithTimeout(t time.Duration) ConfigOption {
	return func(c *Config) {
		c.client.Timeout = t
	}
}
