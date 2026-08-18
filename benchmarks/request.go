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

var customDialer = &net.Dialer{
	Timeout:   120 * time.Second,
	KeepAlive: 30 * time.Second,
}

// reuse a global HTTP client to minimize the script resources impact on the actual tests
var customClient = http.Client{
	Transport: &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           customDialer.DialContext,
		MaxIdleConns:          2000, // don't set to 0 (unlimited) because some VPS providers have stricter restrictions for the max ephemeral ports
		MaxIdleConnsPerHost:   2000,
		IdleConnTimeout:       120 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	},
}

// If destBodyPtr is non-nil, it will read and unmarshal the request
// response body into the specified pointer.
func (c *Request) Send(destBodyPtr any) error {
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

	res, err := customClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return fmt.Errorf("request failed with status %d", res.StatusCode)
	}

	if destBodyPtr != nil {
		bodyRaw, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(bodyRaw, destBodyPtr); err != nil {
			return err
		}
	}

	return nil
}
