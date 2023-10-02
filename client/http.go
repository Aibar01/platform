package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Aibar01/platform/attrs"
	"io"
	"log/slog"
	"net/http"
)

type Client struct {
	baseUrl      string
	contextAttrs attrs.ContextAttrs
}

func NewClient(baseUrl string, headers ...attrs.ContextAttrs) *Client {
	var contextAttrs attrs.ContextAttrs
	if len(headers) > 0 {
		contextAttrs = headers[0]
	}

	return &Client{baseUrl: baseUrl, contextAttrs: contextAttrs}
}

func (c *Client) Get(path string, params any) ([]byte, error) {
	resp, err := c.request("GET", path, params)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Post(path string, data any) ([]byte, error) {
	resp, err := c.request("POST", path, data)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) Patch(path string, data any) ([]byte, error) {
	resp, err := c.request("PATCH", path, data)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) Put(path string, data any) ([]byte, error) {
	resp, err := c.request("PUT", path, data)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) Delete(path string) ([]byte, error) {
	resp, err := c.request("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) request(
	method, path string, data any, headers ...map[string]string,
) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.baseUrl, path)

	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	logger := slog.Default()

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	defaultHeaders := c.contextAttrs.GetByTagName("header")
	requestHeaders := make(map[string]string)

	if len(headers) > 0 {
		for key, value := range headers[0] {
			if value != "" {
				defaultHeaders[key] = value
			}
		}
	}

	for key, value := range defaultHeaders {
		if value != "" {
			req.Header.Add(key, value)
			requestHeaders[key] = value
		}
	}

	jsonRequestHeaders, err := json.Marshal(requestHeaders)
	if err != nil {
		return nil, err
	}

	logger.Info(
		"Request",
		"headers", string(jsonRequestHeaders),
		"data", string(jsonData),
	)

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)

	f := logger.Info
	if resp.StatusCode >= http.StatusBadRequest {
		f = logger.Error
	}

	f("Response",
		"status_code", resp.StatusCode,
		"status", resp.Status,
		"body", string(respBody),
	)

	return respBody, nil
}
