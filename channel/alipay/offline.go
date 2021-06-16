package alipay

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Crystal303/go_alg/channel/core"
	"github.com/Crystal303/go_alg/channel/core/auth/validators"
	"github.com/Crystal303/go_alg/channel/core/util"
)

type alipayOffline struct{}

var Offline alipayOffline

func (a alipayOffline) Register(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSetting(setting); err != nil {
		return nil, err
	}

	reqBody := initRegisterOffline(describer.Describe())

	// 数据校验
	err := validators.DefaultValidate.Validate(reqBody)
	if err != nil {
		return nil, err
	}

	uri, err := genURI(alipayOfflineUrl, setting.SignKey, reqBody)
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

func (a alipayOffline) QueryStatus(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	var setting core.PspSetting
	for _, opt := range opts {
		opt.Apply(&setting)
	}

	if err := checkPSPSetting(setting); err != nil {
		return nil, err
	}

	reqBody := initQueryOffline(describer.Describe())

	// 数据校验
	err := validators.DefaultValidate.Validate(reqBody)
	if err != nil {
		return nil, err
	}

	uri, err := genURI(alipayOfflineUrl, setting.SignKey, reqBody)
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

	respData := new(QueryStatusOfflineResp)
	if err := processResponse(result.Response, respData); err != nil {
		return nil, err
	}

	return respData, nil
}

func (a alipayOffline) Modify(ctx context.Context, describer core.SecMerchantDescriber, opts ...core.RegisterOption) (core.RespBuilder, error) {
	return a.Register(ctx, describer, opts...)
}

func initRegisterOffline(info core.MerchantInfo) *RegisterOffline {
	now := time.Now()
	data := &RegisterOffline{
		BaseParam: BaseParam{
			Service:      ServiceOfflineRegister,
			Partner:      info.ChanMerID,
			InputCharSet: charsetDefault,
			SignType:     "",
			Sign:         "",
			TimeStamp:    now.Format(dateFormat),
		},
		SecondaryMerchantName:   info.SecondaryMerchantName,
		SecondaryMerchantID:     info.SubChanMerID,
		StoreID:                 info.StoreID,
		StoreName:               info.StoreName,
		StoreCountry:            convertCountry(info.StoreCountry),
		StoreAddress:            info.StoreAddress,
		StoreIndustry:           info.StoreMCC,
		InternalStorePhoto:      info.InternalStorePhoto,
		ExternalStorefrontPhoto: info.ExternalStorefrontPhoto,
		ExtendParams:            "",
		SecondaryMerchantType:   info.SecondaryMerchantType,
		RegistrationNo:          info.RegistrationNo,
		RegisterCountry:         convertCountry(info.RegisterCountry),
		RegisterAddress:         info.RegisterAddress,
		ShareholderName:         info.ShareholderName,
		ShareholderID:           info.ShareholderID,
		RepresentativeName:      info.RepresentativeName,
		RepresentativeID:        info.RepresentativeID,
		SettlementNo:            info.SettlementNo,
		ContactNo:               info.ContactNo,
		ContactEmail:            info.ContactEmail,
		CsNo:                    info.CsNo,
		CsEmail:                 info.CsEmail,
	}
	if data.StoreIndustry == mccDriver {
		bts, _ := json.Marshal(info.ExtendParams)
		data.ExtendParams = util.SliceByteToString(bts)
	}
	return data
}

func initQueryOffline(info core.MerchantInfo) *QueryStatusOffline {
	now := time.Now()
	data := &QueryStatusOffline{
		BaseParam: BaseParam{
			Service:      ServiceOfflineQueryStatus,
			Partner:      info.ChanMerID,
			InputCharSet: charsetDefault,
			SignType:     "",
			Sign:         "",
			TimeStamp:    now.Format(dateFormat),
		},
		SecondaryMerchantID: info.SubChanMerID,
		PaymentMethod:       paymentMethodOffline,
		StoreID:             info.StoreID,
	}
	return data
}
