package alipay

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Crystal303/go_alg/channel/core"
	"github.com/Crystal303/go_alg/channel/core/auth/validators"
	"github.com/Crystal303/go_alg/channel/core/util"
)

type alipayOnline struct{}

var Online alipayOnline

func (a alipayOnline) Register(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSetting(setting); err != nil {
		return nil, err
	}

	reqBody := initRegisterOnline(describer.Describe())

	// 数据校验
	err := validators.DefaultValidate.Validate(reqBody)
	if err != nil {
		return nil, err
	}

	uri, err := genURI(alipayOnlineUrl, setting.SignKey, reqBody)
	if err != nil {
		return nil, err
	}

	client, err := core.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	result, err := client.Get(ctx, uri)
	if err != nil {
		return nil, err
	}

	respData := new(RegisterResp)
	if err := processResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (a alipayOnline) QueryStatus(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSetting(setting); err != nil {
		return nil, err
	}

	reqBody := initQueryOnline(describer.Describe())

	// 数据校验
	err := validators.DefaultValidate.Validate(reqBody)
	if err != nil {
		return nil, err
	}

	uri, err := genURI(alipayOnlineUrl, setting.SignKey, reqBody)
	if err != nil {
		return nil, err
	}

	client, err := core.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	result, err := client.Get(ctx, uri)
	if err != nil {
		return nil, err
	}

	respData := new(QueryStatusOnlineResp)
	if err := processResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (a alipayOnline) Modify(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	return a.Register(ctx, describer, opts...)
}

func initRegisterOnline(info core.MerchantInfo) *RegisterOnline {
	now := time.Now()
	data := &RegisterOnline{
		BaseParam: BaseParam{
			Service:      ServiceOnlineRegister,
			Partner:      info.ChanMerID,
			InputCharSet: charsetDefault,
			SignType:     "",
			Sign:         "",
			TimeStamp:    now.Format(dateFormat),
		},
		SecondaryMerchantName:     info.SecondaryMerchantName,
		SecondaryMerchantID:       info.SubChanMerID,
		SecondaryMerchantIndustry: info.SecondaryMerchantMCC,
		RegisterCountry:           convertCountry(info.RegisterCountry),
		RegisterAddress:           info.RegisterAddress,
		SiteInfos:                 "",
		SecondaryMerchantType:     info.SecondaryMerchantType,
		RegistrationNo:            info.RegistrationNo,
		ShareholderName:           info.ShareholderName,
		ShareholderID:             info.ShareholderID,
		RepresentativeName:        info.RepresentativeName,
		RepresentativeID:          info.RepresentativeID,
		SettlementNo:              info.SettlementNo,
		ContactNo:                 info.ContactNo,
		ContactEmail:              info.ContactEmail,
		CsNo:                      info.CsNo,
		CsEmail:                   info.CsEmail,
	}
	bts, _ := json.Marshal(info.SiteInfos)
	data.SiteInfos = util.SliceByteToString(bts)

	return data
}

func initQueryOnline(info core.MerchantInfo) *QueryStatusOnline {
	now := time.Now()
	data := &QueryStatusOnline{
		BaseParam: BaseParam{
			Service:      ServiceOnlineQueryStatus,
			Partner:      info.ChanMerID,
			InputCharSet: charsetDefault,
			SignType:     "",
			Sign:         "",
			TimeStamp:    now.Format(dateFormat),
		},
		SecondaryMerchantID: info.SubChanMerID,
		PaymentMethod:       paymentMethodOnline,
	}
	return data
}
