package core

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type withSignKey struct{ signKey string }

// Apply 将配置添加到PspSetting中
func (w withSignKey) Apply(o *PspSetting) {
	o.SignKey = w.signKey
}

// 设置PSP的signKey
func WithSignKey(signKey string) RegisterOption {
	return withSignKey{signKey: signKey}
}

type withCertificate struct {
	clientCert, clientKey []byte // HTTPS 加密证书
}

func (w withCertificate) Apply(o *PspSetting) {
	o.ClientCert, o.ClientKey = w.clientCert, w.clientKey
}

// 设置PSP的certificate
func WithCertificate(clientCert, clientKey []byte) RegisterOption {
	return withCertificate{
		clientCert: clientCert,
		clientKey:  clientKey,
	}
}

type withPortalCert struct {
	portalCert []byte
}

func (w withPortalCert) Apply(o *PspSetting) {
	o.PortalCert = w.portalCert
}

// 设置PSP的平台公钥
func WithPortalCert(publicKey []byte) RegisterOption {
	return withPortalCert{portalCert: publicKey}
}

type withHTTPsClient struct {
	client *http.Client
}

func (w withHTTPsClient) Apply(settings *dialSettings) {
	settings.HTTPClient = w.client
}

// 设置 Client 的证书
func WithHTTPsClient(certificate tls.Certificate) ClientOption {
	return withHTTPsClient{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				TLSClientConfig: &tls.Config{
					Certificates: []tls.Certificate{certificate},
				},
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		},
	}
}

type withHTTPHeader struct {
	header http.Header
}

func (w withHTTPHeader) Apply(settings *dialSettings) {
	settings.Header = w.header
}

func WithHTTPHeader(header map[string][]string) ClientOption {
	return withHTTPHeader{header: header}
}
