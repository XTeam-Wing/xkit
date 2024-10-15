package push

import (
	"context"
	"crypto/tls"
	"errors"
	"github.com/imroc/req/v3"
	"time"
)

func NewHttpClient() *req.Client {
	client := req.C()
	customTransport := &tls.Config{
		InsecureSkipVerify: true, // 忽略 SSL 证书验证
	}

	client.
		ImpersonateChrome().
		SetTLSClientConfig(customTransport).
		SetTimeout(10 * time.Second).
		SetCommonRetryCount(3).
		SetCookieJar(nil).
		SetCommonRetryInterval(func(resp *req.Response, attempt int) time.Duration {
			if errors.Is(resp.Err, context.Canceled) {
				return 0
			}
			return time.Second * 5
		}).
		SetCommonRetryHook(func(resp *req.Response, err error) {
		}).SetCommonRetryCondition(func(resp *req.Response, err error) bool {
		if err != nil {
			return !errors.Is(err, context.Canceled)
		}
		return false
	})
	return client
}
