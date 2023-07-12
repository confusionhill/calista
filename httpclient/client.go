package httpclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	timeout time.Duration
	client  *http.Client
}

const (
	defaultHTTPTimeout = 30 * time.Second
)

func NewClient(opts ...Option) (*HTTPClient, error) {
	calista := HTTPClient{
		timeout: defaultHTTPTimeout,
	}
	for _, ops := range opts {
		ops(&calista)
	}
	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       calista.timeout,
	}
	calista.client = &client
	return &calista, nil
}

func (h *HTTPClient) Request(
	url string,
	method string,
	header map[string]string,
	body interface{},
	response interface{},
) error {
	jsonBody, _ := json.Marshal(body)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	for key, value := range header {
		req.Header.Set(key, value)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(resp.Status)
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(respBody, &response); err != nil {
		return err
	}
	return nil
}

func (h *HTTPClient) Get(
	url string,
	header map[string]string,
	response interface{},
) error {
	return h.Request(url, http.MethodGet, header, nil, response)
}

func (h *HTTPClient) Post(
	url string,
	header map[string]string,
	body interface{},
	response interface{},
) error {
	return h.Request(url, http.MethodPost, header, body, response)
}

func (h *HTTPClient) Put(
	url string,
	header map[string]string,
	body interface{},
	response interface{},
) error {
	return h.Request(url, http.MethodPut, header, body, response)
}

func (h *HTTPClient) Patch(
	url string,
	header map[string]string,
	body interface{},
	response interface{},
) error {
	return h.Request(url, http.MethodPatch, header, body, response)
}

func (h *HTTPClient) Delete(
	url string,
	header map[string]string,
	body interface{},
	response interface{},
) error {
	return h.Request(url, http.MethodDelete, header, body, response)
}
