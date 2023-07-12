package httpclient

import "time"

type Option func(client *HTTPClient)

func WithTimeout(timeout time.Duration) Option {
	return func(c *HTTPClient) {
		c.timeout = timeout
	}
}
