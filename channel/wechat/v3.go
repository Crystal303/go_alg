package wechat

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Crystal303/go_alg/channel/consts"
	"github.com/Crystal303/go_alg/channel/core"
	"github.com/Crystal303/go_alg/channel/core/util"
)

type wechatV3 struct{}

var V3 wechatV3

func (w wechatV3) Register(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSettingV3(setting); err != nil {
		return nil, err
	}

	reqBody := initRegisterV3(describer.Describe())

	var err error
	reqBody.Contact.Name, err = util.RSAEncryptBase64(reqBody.Contact.Name, setting.PortalCert)
	if err != nil {
		return nil, err
	}
	reqBody.Contact.Phone, err = util.RSAEncryptBase64(reqBody.Contact.Phone, setting.PortalCert)
	if err != nil {
		return nil, err
	}
	reqBody.Contact.Email, err = util.RSAEncryptBase64(reqBody.Contact.Email, setting.PortalCert)
	if err != nil {
		return nil, err
	}

	// 设置签名
	certificateSerialNo, err := getCertificateSerialNo(setting.ClientCert)
	if err != nil {
		return nil, err
	}
	authorization, err := genAuthorizationHeader(reqBody.SpMchID, ServiceRegisterV3, certificateSerialNo, httpPOSTMethod, setting.ClientKey, reqBody)
	if err != nil {
		return nil, err
	}

	// 生成tls证书
	certificate, err := tls.X509KeyPair(setting.ClientCert, setting.ClientKey)
	if err != nil {
		return nil, err
	}
	client, err := core.NewClient(ctx, core.WithHTTPsClient(certificate), core.WithHTTPHeader(map[string][]string{
		consts.Accept:        {consts.ApplicationJSON},
		consts.Authorization: {authorization},
		consts.UserAgent:     {consts.UserAgentContent},
		WechatPaySerial:      {certificateSerialNo},
	}))
	if err != nil {
		return nil, err
	}

	result, err := client.PostWithJSONBody(ctx, UrlWechatV3+ServiceRegisterV3, reqBody)
	if err != nil {
		return nil, err
	}

	respData := new(RespV3)
	if err := processJSONResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (w wechatV3) QueryStatus(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSettingV3(setting); err != nil {
		return nil, err
	}

	reqBody := initQueryStatusV3(describer.Describe())

	// 设置签名
	certificateSerialNo, err := getCertificateSerialNo(setting.ClientCert)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(ServiceQueryStatusV3, reqBody.SubMchID)
	authorization, err := genAuthorizationHeader(reqBody.SpMchID, url, certificateSerialNo, httpGETMethod, setting.ClientKey, reqBody)
	if err != nil {
		return nil, err
	}

	// 生成tls证书
	certificate, err := tls.X509KeyPair(setting.ClientCert, setting.ClientKey)
	if err != nil {
		return nil, err
	}
	client, err := core.NewClient(ctx, core.WithHTTPsClient(certificate), core.WithHTTPHeader(map[string][]string{
		consts.Accept:        {consts.ApplicationJSON},
		consts.Authorization: {authorization},
		consts.UserAgent:     {consts.UserAgentContent},
		WechatPaySerial:      {certificateSerialNo},
	}))
	if err != nil {
		return nil, err
	}

	result, err := client.GETWithJSONBody(ctx, UrlWechatV3+url, reqBody)
	if err != nil {
		return nil, err
	}

	respData := new(QueryStatusV3Resp)
	if err := processJSONResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (w wechatV3) Modify(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSettingV3(setting); err != nil {
		return nil, err
	}

	reqBody := initModifyV3(describer.Describe())

	var err error
	reqBody.Contact.Name, err = util.RSAEncryptBase64(reqBody.Contact.Name, setting.PortalCert)
	if err != nil {
		return nil, err
	}
	reqBody.Contact.Phone, err = util.RSAEncryptBase64(reqBody.Contact.Phone, setting.PortalCert)
	if err != nil {
		return nil, err
	}
	reqBody.Contact.Email, err = util.RSAEncryptBase64(reqBody.Contact.Email, setting.PortalCert)
	if err != nil {
		return nil, err
	}

	// 设置签名
	certificateSerialNo, err := getCertificateSerialNo(setting.ClientCert)
	if err != nil {
		return nil, err
	}
	authorization, err := genAuthorizationHeader(reqBody.SpMchID, ServiceModifyV3, certificateSerialNo, httpPUTMethod, setting.ClientKey, reqBody)
	if err != nil {
		return nil, err
	}

	// 生成tls证书
	certificate, err := tls.X509KeyPair(setting.ClientCert, setting.ClientKey)
	if err != nil {
		return nil, err
	}
	client, err := core.NewClient(ctx, core.WithHTTPsClient(certificate), core.WithHTTPHeader(map[string][]string{
		consts.Accept:        {consts.ApplicationJSON},
		consts.Authorization: {authorization},
		consts.UserAgent:     {consts.UserAgentContent},
		WechatPaySerial:      {certificateSerialNo},
	}))
	if err != nil {
		return nil, err
	}

	result, err := client.PUTWithJSONBody(ctx, UrlWechatV3+ServiceModifyV3, reqBody)
	if err != nil {
		return nil, err
	}

	respData := new(RespV3)
	if err := processJSONResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func checkPSPSettingV3(setting core.PspSetting) error {
	if setting.ClientCert == nil {
		return fmt.Errorf("empty %s", "ClientCert")
	}
	if setting.ClientKey == nil {
		return fmt.Errorf("empty %s", "ClientKey")
	}
	block, _ := pem.Decode(setting.PortalCert)
	if block == nil {
		return fmt.Errorf("invalid %s", "PortalCert")
	}
	_, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	return nil
}

func processJSONResponse(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bts, &data); err != nil {
		return err
	}
	return nil
}

func genAuthorizationHeader(mchID, url, certificateSerialNo, httpMethod string, clientKey []byte, data interface{}) (string, error) {
	signBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	nonce, err := generateNonce()
	if err != nil {
		return "", err
	}

	timestamp := time.Now().Unix()
	message := fmt.Sprintf(SignatureMessageFormat, httpMethod, url, timestamp, nonce, signBody)

	blocks, _ := pem.Decode(clientKey)
	if blocks == nil {
		return "", fmt.Errorf("invalid %s", "ClientKey")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(blocks.Bytes)
	if err != nil {
		return "", err
	}
	signature, err := signSHA256WithRSA(message, privateKey.(*rsa.PrivateKey))
	if err != nil {
		return "", err
	}

	authorization := fmt.Sprintf(HeaderAuthorizationFormat, mchID, nonce, timestamp,
		certificateSerialNo, signature)
	return authorization, nil
}

func generateNonce() (string, error) {
	bytes := make([]byte, NonceLength)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	symbolsByteLength := byte(len(NonceSymbols))
	for i, b := range bytes {
		bytes[i] = NonceSymbols[b%symbolsByteLength]
	}
	return string(bytes), nil
}

func signSHA256WithRSA(source string, privateKey *rsa.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	h := crypto.Hash.New(crypto.SHA256)
	_, err = h.Write([]byte(source))
	if err != nil {
		return "", nil
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}

func getCertificateSerialNo(pemCert []byte) (string, error) {
	block, _ := pem.Decode(pemCert)
	if block == nil {
		return "", fmt.Errorf("invalid %s", "ClientCert")
	}

	certificate, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}
	certID := certificate.SerialNumber.Text(16)

	return certID, nil
}

func initRegisterV3(info core.MerchantInfo) *RegisterV3Req {
	data := &RegisterV3Req{
		SpAppID:     info.AppID,
		SpMchID:     info.ChanMerID,
		Name:        info.StoreName,
		ShortName:   info.StoreNameEN,
		OfficePhone: info.OfficePhone,
		Contact: Contact{
			Name:  info.ContactName,
			Phone: info.ContactNo,
			Email: info.ContactEmail,
		},
		ChannelID:                     "",
		MerchantCountryCode:           convertCountry(info.RegisterCountry),
		MerchantType:                  info.SecondaryMerchantType,
		RegistrationCertificateNumber: info.RegistrationNo,
		RegistrationCertificateDate:   info.RegistrationCertificateDate,
		RegistrationCertificateCopy:   "",
		SettlementBankNumber:          info.SettlementNo,
		Business: Business{
			BusinessType:  info.BusinessType,
			AppDownload:   info.AppDownload,
			Website:       info.BusinessWebsite,
			OfficeAccount: info.OfficeAccount,
			MiniProgram:   info.MiniProgram,
			StoreAddress:  info.StoreAddress,
			StorePhotos:   "", // todo
			Mcc:           info.StoreMCC,
		},
		Director:  nil,
		Principal: nil,
	}
	if data.MerchantType == consts.MerchantTypeENTERPRISE {
		data.Director = &Director{
			Name:   info.ShareholderName,
			Number: info.ShareholderID,
		}
	}
	if data.MerchantType == consts.MerchantTypeINDIVIDUAL {
		data.Principal = &Principal{
			Name:   info.RepresentativeName,
			Number: info.RepresentativeID,
		}
	}
	businessCategory := ""
	data.BusinessCategory, _ = strconv.Atoi(businessCategory)

	return data
}

func initQueryStatusV3(info core.MerchantInfo) *QueryStatusV3Req {
	return &QueryStatusV3Req{
		SpAppID:  info.AppID,
		SpMchID:  info.ChanMerID,
		SubMchID: info.SubChanMerID,
	}
}

func initModifyV3(info core.MerchantInfo) *ModifyV3Req {
	return &ModifyV3Req{
		SubMchID:      info.SubChanMerID,
		RegisterV3Req: *initRegisterV3(info),
	}
}
