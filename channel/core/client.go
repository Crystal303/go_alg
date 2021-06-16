package core

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/Crystal303/go_alg/channel/consts"
	"github.com/Crystal303/go_alg/channel/core/auth/validators"
)

type APIResult struct {
	// 本次请求所获得的 HTTPResponse
	Response *http.Response
}

type Client struct {
	httpClient    *http.Client
	defaultHeader http.Header
	// signer        signers.Signer 生成签名
	postValidator validators.Validator      // 校验返回response
	preValidator  validators.ValidateStruct // 校验请求参数 用于 POST PUT
}

func NewClient(ctx context.Context, opts ...ClientOption) (client *Client, err error) {
	settings, err := initSettings(opts)
	if err != nil {
		return nil, fmt.Errorf("init client setting err:%v", err)
	}
	client = &Client{
		postValidator: settings.PostValidator,
		httpClient:    settings.HTTPClient,
		defaultHeader: settings.Header,
		// signer:        settings.Signer,
		preValidator: settings.PreValidator,
	}
	return client, nil
}

func initSettings(opts []ClientOption) (*dialSettings, error) {
	var o dialSettings
	for _, opt := range opts {
		opt.Apply(&o)
	}
	if o.HTTPClient == nil {
		o.HTTPClient = http.DefaultClient
	}
	if o.Timeout != 0 {
		o.HTTPClient.Timeout = o.Timeout
	}
	if o.PreValidator == nil {
		o.PreValidator = validators.DefaultValidate
	}
	if o.PostValidator == nil {
		o.PostValidator = validators.NullValidator
	}
	return &o, nil
}

func (client *Client) Get(ctx context.Context, requestURL string) (*APIResult, error) {
	return client.doRequest(ctx, http.MethodGet, requestURL, nil, nil)
}

func (client *Client) PostWithXMLBody(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	if err := client.preValidator.Validate(requestBody); err != nil {
		return nil, err
	}

	bodyBuf := new(bytes.Buffer)
	if err := xml.NewEncoder(bodyBuf).Encode(requestBody); err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Set(consts.ContentType, consts.XmlContentType)

	return client.doRequest(ctx, http.MethodPost, requestURL, nil, bodyBuf)
}

func (client *Client) PostWithJSONBody(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	if err := client.preValidator.Validate(requestBody); err != nil {
		return nil, err
	}

	bodyBuf := new(bytes.Buffer)
	if err := json.NewEncoder(bodyBuf).Encode(requestBody); err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Set(consts.ContentType, consts.ApplicationJSON)

	return client.doRequest(ctx, http.MethodPost, requestURL, header, bodyBuf)
}

func (client *Client) GETWithJSONBody(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	if err := client.preValidator.Validate(requestBody); err != nil {
		return nil, err
	}

	bodyBuf := new(bytes.Buffer)
	if err := json.NewEncoder(bodyBuf).Encode(requestBody); err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Set(consts.ContentType, consts.ApplicationJSON)

	return client.doRequest(ctx, http.MethodGet, requestURL, header, bodyBuf)
}

func (client *Client) PUTWithJSONBody(ctx context.Context, requestURL string, requestBody interface{}) (*APIResult, error) {
	if err := client.preValidator.Validate(requestBody); err != nil {
		return nil, err
	}

	bodyBuf := new(bytes.Buffer)
	if err := json.NewEncoder(bodyBuf).Encode(requestBody); err != nil {
		return nil, err
	}

	header := make(http.Header)
	header.Set(consts.ContentType, consts.ApplicationJSON)

	return client.doRequest(ctx, http.MethodPut, requestURL, header, bodyBuf)
}

func (client *Client) doRequest(ctx context.Context, method string, requestURL string, header http.Header, reqBody io.Reader) (*APIResult, error) {
	var (
		err     error
		request *http.Request
	)
	if request, err = http.NewRequestWithContext(ctx, method, requestURL, reqBody); err != nil {
		return nil, err
	}

	// header
	for key, values := range client.defaultHeader {
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}
	if header != nil {
		for key, values := range header {
			for _, v := range values {
				request.Header.Add(key, v)
			}
		}
	}

	result, err := client.doHTTP(request)
	if err != nil {
		return result, err
	}

	// 验证签名
	if err := client.postValidator.ValidateResp(ctx, result.Response); err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) doHTTP(req *http.Request) (result *APIResult, err error) {
	result = new(APIResult)

	result.Response, err = client.httpClient.Do(req)
	return result, err
}
