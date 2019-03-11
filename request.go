package go_wechat

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const DeaultTimeout = 10 * time.Second

type Request interface {
	Get(reqUrl string, params map[string]string) (*http.Response, error)
	Post(reqUrl, contentType string, body io.Reader) (*http.Response, error)
}

// 先做一个简单内置的实现
type request struct {
	*http.Client
	Timeout time.Duration
}

func NewRequest(timeout time.Duration) Request {
	if timeout <= 0 {
		timeout = DeaultTimeout
	}
	return &request{
		Timeout: timeout,
		Client:  &http.Client{},
	}
}

func (r *request) Get(reqUrl string, params map[string]string) (*http.Response, error) {
	paramStr, err := GenRequestString(params)
	if err != nil {
		return nil, err
	}
	reqUrl = reqUrl + "?" + paramStr
	return r.do(http.MethodGet, reqUrl, nil, nil)
}

func (r *request) Post(reqUrl, contentType string, body io.Reader) (*http.Response, error) {
	headers := map[string]string{
		"Content-Type": contentType,
	}
	return r.do(http.MethodPost, reqUrl, body, headers)
}

func (r *request) do(method string, reqUrl string, body io.Reader, headers map[string]string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.Timeout)
	defer cancel()

	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return nil, fmt.Errorf("make http req error %v", err)
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req = req.WithContext(ctx)
	return r.Do(req)
}
