package wechat

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Crystal303/go_alg/channel/core"
	"github.com/Crystal303/go_alg/channel/core/query"
	"github.com/Crystal303/go_alg/channel/core/util"
)

type wechatV2 struct{}

var V2 wechatV2

func (w wechatV2) Register(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSettingV2(setting); err != nil {
		return nil, err
	}

	reqBody := initRegister(describer.Describe())

	// 设置签名
	signed, err := signV2(setting.SignKey, reqBody)
	if err != nil {
		return nil, err
	}
	reqBody.Sign = signed

	// 生成tls证书
	certificate, err := tls.X509KeyPair(setting.ClientCert, setting.ClientKey)
	if err != nil {
		return nil, err
	}
	client, err := core.NewClient(ctx, core.WithHTTPsClient(certificate))
	if err != nil {
		return nil, err
	}

	result, err := client.PostWithXMLBody(ctx, UrlWechatV2+ServiceRegisterV2, reqBody)
	if err != nil {
		return nil, err
	}

	respData := new(RegisterResp)
	if err := processXMLResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (w wechatV2) QueryStatus(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSettingV2(setting); err != nil {
		return nil, err
	}

	reqBody := initQueryStatus(describer.Describe())

	// 设置签名
	signed, err := signV2(setting.SignKey, reqBody)
	if err != nil {
		return nil, err
	}
	reqBody.Sign = signed

	// 生成tls证书
	certificate, err := tls.X509KeyPair(setting.ClientCert, setting.ClientKey)
	if err != nil {
		return nil, err
	}
	client, err := core.NewClient(ctx, core.WithHTTPsClient(certificate))
	if err != nil {
		return nil, err
	}

	result, err := client.PostWithXMLBody(ctx, UrlWechatV2+ServiceQueryStatusV2, reqBody)
	if err != nil {
		return nil, err
	}

	respData := new(QueryStatusResp)
	if err := processXMLResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (w wechatV2) Modify(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSettingV2(setting); err != nil {
		return nil, err
	}

	reqBody := initModify(describer.Describe())

	// 设置签名
	signed, err := signV2(setting.SignKey, reqBody)
	if err != nil {
		return nil, err
	}
	reqBody.Sign = signed

	// 生成tls证书
	certificate, err := tls.X509KeyPair(setting.ClientCert, setting.ClientKey)
	if err != nil {
		return nil, err
	}
	client, err := core.NewClient(ctx, core.WithHTTPsClient(certificate))
	if err != nil {
		return nil, err
	}

	result, err := client.PostWithXMLBody(ctx, UrlWechatV2+ServiceModifyV2, reqBody)
	if err != nil {
		return nil, err
	}

	respData := new(ModifyResp)
	if err := processXMLResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func checkPSPSettingV2(setting core.PspSetting) error {
	if strings.TrimSpace(setting.SignKey) == "" {
		return fmt.Errorf("empty %s", "SignKey")
	}
	if setting.ClientCert == nil {
		return fmt.Errorf("empty %s", "ClientCert")
	}
	if setting.ClientKey == nil {
		return fmt.Errorf("empty %s", "ClientKey")
	}
	return nil
}

func initRegister(info core.MerchantInfo) *RegisterReq {
	data := &RegisterReq{
		MerchantParams: MerchantParams{
			AppID: info.AppID,
			MchID: info.ChanMerID,
			Sign:  "",
			// ChannelID:                     "",
			MerchantName:                  info.StoreName,
			MerchantShortName:             info.StoreNameEN,
			MerchantCountryCode:           convertCountry(info.RegisterCountry),
			MerchantType:                  info.SecondaryMerchantType,
			BusinessCategory:              "", // todo 转换至 3 位
			Mcc:                           info.StoreMCC,
			RegistrationCertificateNumber: info.RegistrationNo,
			RegistrationCertificateDate:   info.RegistrationCertificateDate,
			// RegistrationCertificateCopy:   "",
			BusinessType:         info.BusinessType,
			AppDownload:          info.AppDownload,
			BusinessWebsite:      info.BusinessWebsite,
			OfficeAccount:        info.OfficeAccount,
			MiniProgram:          info.MiniProgram,
			StoreAddress:         info.StoreAddress,
			StorePhotos:          "", // todo photos
			DirectorName:         info.ShareholderName,
			DirectorIDNumber:     info.ShareholderID,
			PrincipalName:        info.RepresentativeName,
			PrincipalIDNumber:    info.RepresentativeID,
			OfficePhone:          info.OfficePhone,
			ContactName:          info.ContactName,
			ContactPhone:         info.ContactNo,
			ContactEmail:         info.ContactEmail,
			SettlementBankNumber: info.SettlementNo,
		},
		MerchantRemark: info.StoreNum,
	}
	return data
}

func initQueryStatus(info core.MerchantInfo) *QueryStatusReq {
	return &QueryStatusReq{
		AppID:    info.AppID,
		MchID:    info.ChanMerID,
		Sign:     "",
		SubMchID: info.SubChanMerID,
	}
}

func initModify(info core.MerchantInfo) *ModifyReq {
	data := &ModifyReq{
		MerchantParams: MerchantParams{
			AppID: info.AppID,
			MchID: info.ChanMerID,
			Sign:  "",
			// ChannelID:                     "",
			MerchantName:                  info.StoreName,
			MerchantShortName:             info.StoreNameEN,
			MerchantCountryCode:           convertCountry(info.RegisterCountry),
			MerchantType:                  info.SecondaryMerchantType,
			BusinessCategory:              "", // todo
			Mcc:                           info.StoreMCC,
			RegistrationCertificateNumber: info.RegistrationNo,
			RegistrationCertificateDate:   info.RegistrationCertificateDate,
			// RegistrationCertificateCopy:   "",
			BusinessType:         info.BusinessType,
			AppDownload:          info.AppDownload,
			BusinessWebsite:      info.BusinessWebsite,
			OfficeAccount:        info.OfficeAccount,
			MiniProgram:          info.MiniProgram,
			StoreAddress:         info.StoreAddress,
			StorePhotos:          "", // todo photos
			DirectorName:         info.ShareholderName,
			DirectorIDNumber:     info.ShareholderID,
			PrincipalName:        info.RepresentativeName,
			PrincipalIDNumber:    info.RepresentativeID,
			OfficePhone:          info.OfficePhone,
			ContactName:          info.ContactName,
			ContactPhone:         info.ContactNo,
			ContactEmail:         info.ContactEmail,
			SettlementBankNumber: info.SettlementNo,
		},
		SubMchID: info.SubChanMerID,
	}
	return data
}

func signV2(signKey string, data interface{}) (string, error) {
	values, err := query.Values(data, tagName)
	if err != nil {
		return "", err
	}

	content := query.Encode(values) + "&key=" + signKey

	signed := md5.Sum(util.StringToSliceByte(content))

	return strings.ToUpper(hex.EncodeToString(signed[:])), nil
}

func processXMLResponse(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(bts, &data); err != nil {
		return err
	}
	return nil
}
