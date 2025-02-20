package rpc

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	url        string
}

func NewClient(url string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		url:        url,
	}
}

func (c *Client) CreateRequestPayload(method Method, params ...interface{}) (*Request, error) {
	par, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req := &Request{
		JsonRpc: "1.0",
		Method:  method,
		Params:  par,
		Id:      1,
	}

	return req, nil
}

func (c *Client) HandleRequest(request *Request) (*Response, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	if err = json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
