package benchmarks

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type Request struct {
	Body    io.Reader
	Context context.Context
	Headers map[string]string
	Method  string
	Url     string
}

// If destBody is non-nil, it will read and unmarshal the request
// response body into the specified variable.
func (c *Request) Send(destBody any) error {
	if c.Context == nil {
		c.Context = context.Background()
	}

	req, err := http.NewRequestWithContext(c.Context, c.Method, c.Url, c.Body)
	if err != nil {
		return err
	}

	for k, v := range c.Headers {
		req.Header.Add(k, v)
	}

	// set default content-type header (if missing)
	if req.Header.Get("content-type") == "" {
		req.Header.Set("content-type", "application/json")
	}

	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			MaxIdleConns:          0,
			IdleConnTimeout:       120 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", res.StatusCode)
	}

	if destBody != nil {
		bodyRaw, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(bodyRaw, destBody); err != nil {
			return err
		}
	}

	return nil
}
