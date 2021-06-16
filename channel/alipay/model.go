package alipay

import (
	"github.com/Crystal303/go_alg/channel/consts"
	"github.com/Crystal303/go_alg/channel/core"
)

// 基本参数
type BaseParam struct {
	Service      string `url:"service" validate:"required"`                        // 枚举值
	Partner      string `url:"partner" validate:"required,len=16,startswith=2088"` // PID
	InputCharSet string `url:"_input_charset" validate:"required"`                 // 编码集 UTF-8
	SignType     string `url:"-"`                                                  // 签名类型 RSA RSA2 MD5 DSA
	Sign         string `url:"-"`                                                  // 签名值
	TimeStamp    string `url:"timestamp" validate:"required"`                      // yyyy-MM-dd HH:mm:ss 例 2019-02-01 08:30:10
}

// 支付宝线下
type QueryStatusOffline struct {
	BaseParam

	// 业务参数
	SecondaryMerchantID string `url:"secondary_merchant_id" validate:"required,max=64"`      // 标志二级商户的唯一ID 字母 数字 下划线
	PaymentMethod       string `url:"payment_method" validate:"required,eq=INSTORE_PAYMENT"` // 二级商户支付方式 枚举 INSTORE_PAYMENT
	StoreID             string `url:"store_id" validate:"required"`                          // 二级商户商店的唯一ID PID/MID 下唯一
}

// 支付宝线上
type QueryStatusOnline struct {
	BaseParam

	// 业务参数
	SecondaryMerchantID string `url:"secondary_merchant_id" validate:"required,max=64"`     // 标志二级商户的唯一ID 字母 数字 下划线
	PaymentMethod       string `url:"payment_method" validate:"required,eq=ONLINE_PAYMENT"` // 二级商户支付方式 枚举 ONLINE_PAYMENT
}

// 支付宝线下
type RegisterOffline struct {
	BaseParam

	SecondaryMerchantName   string `url:"secondary_merchant_name" validate:"required,max=128"` // 二级商户的法定注册名称
	SecondaryMerchantID     string `url:"secondary_merchant_id" validate:"required,max=64"`
	StoreID                 string `url:"store_id" validate:"required,max=64"`
	StoreName               string `url:"store_name" validate:"required,max=256"`    // 商店名称 车牌号
	StoreCountry            string `url:"store_country" validate:"required,len=2"`   // 2位代码 eg.HK
	StoreAddress            string `url:"store_address" validate:"required,max=330"` // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	StoreIndustry           string `url:"store_industry" validate:"required,len=4"`  // MCC 4位代码 eg. 4121
	InternalStorePhoto      string `url:"internal_store_photo,omitempty" validate:"max=256"`
	ExternalStorefrontPhoto string `url:"external_storefront_photo,omitempty" validate:"max=256"`
	ExtendParams            string `url:"extend_params,omitempty" validate:"max=1024,required_if=StoreIndustry 4121"` // JSON 格式 报备一位司机并在store_industry参数填写4121时必填
	// 注意：
	// - 每个operation_id的值必须是唯一的。
	// - 一旦报备了某位司机，此字段信息不支持更新。
	// - 此参数最多可输入10位司机的信息。

	// eg :	[
	//			{"operation_id": "1000332", "contact_way": "138xxxx1232", "contact_person": "driverName1"},
	//			{"operation_id": "1082943492", "contact_way": "158xxxx2232", "contact_person": "driverName2"}
	//		]

	SecondaryMerchantType string `url:"secondary_merchant_type" validate:"required,oneof=INDIVIDUAL ENTERPRISE"` // 枚举 INDIVIDUAL ENTERPRISE
	RegistrationNo        string `url:"registration_no,omitempty" validate:"max=128"`
	RegisterCountry       string `url:"register_country" validate:"len=2"`                                                            // 2位代码 eg.HK
	RegisterAddress       string `url:"register_address,omitempty" validate:"max=256"`                                                // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	ShareholderName       string `url:"shareholder_name,omitempty" validate:"max=64,required_if=SecondaryMerchantType ENTERPRISE"`    // ENTERPRISE 必填
	ShareholderID         string `url:"shareholder_id,omitempty" validate:"max=128,required_if=SecondaryMerchantType ENTERPRISE"`     // ENTERPRISE 必填
	RepresentativeName    string `url:"representative_name,omitempty" validate:"max=64,required_if=SecondaryMerchantType INDIVIDUAL"` // INDIVIDUAL 必填
	RepresentativeID      string `url:"representative_id,omitempty" validate:"max=128,required_if=SecondaryMerchantType INDIVIDUAL"`  // INDIVIDUAL 必填
	SettlementNo          string `url:"settlement_no,omitempty" validate:"max=64"`                                                    // 银行账户 仅支持字母和数字
	ContactNo             string `url:"contact_no,omitempty" validate:"max=64"`                                                       // 支持数字及特殊字符+-()
	ContactEmail          string `url:"contact_email,omitempty" validate:"omitempty,email,max=128"`
	CsNo                  string `url:"cs_no,omitempty" validate:"max=64"`                     // 客服电话
	CsEmail               string `url:"cs_email,omitempty" validate:"omitempty,email,max=128"` // 客户邮箱
}

// 支付宝线上
type RegisterOnline struct {
	BaseParam

	SecondaryMerchantName     string `url:"secondary_merchant_name" validate:"required,max=64"` // 二级商户的法定注册名称
	SecondaryMerchantID       string `url:"secondary_merchant_id" validate:"required,max=64"`
	SecondaryMerchantIndustry string `url:"secondary_merchant_industry" validate:"required,len=4"`
	RegisterCountry           string `url:"register_country" validate:"required,len=2"`   // 2位代码 eg.HK
	RegisterAddress           string `url:"register_address" validate:"required,max=256"` // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	SiteInfos                 string `url:"site_infos" validate:"required"`
	//
	// 二级商户网站或APP URL
	// 格式： [
	//	{"site_type":"WEB","site_url":"https://alipay.com","site_name":"website"},
	//	{"site_type":"APP","site_url":"https://alipay.com","site_name":"website"}
	//	]

	SecondaryMerchantType string `url:"secondary_merchant_type" validate:"required"` // 枚举 INDIVIDUAL ENTERPRISE
	RegistrationNo        string `url:"registration_no,omitempty" validate:"max=128"`
	ShareholderName       string `url:"shareholder_name,omitempty" validate:"max=64"`    // ENTERPRISE 必填
	ShareholderID         string `url:"shareholder_id,omitempty" validate:"max=128"`     // ENTERPRISE 必填
	RepresentativeName    string `url:"representative_name,omitempty" validate:"max=64"` // INDIVIDUAL 必填
	RepresentativeID      string `url:"representative_id,omitempty" validate:"max=128"`  // INDIVIDUAL 必填
	SettlementNo          string `url:"settlement_no,omitempty" validate:"max=64"`       // 银行账户 仅支持字母和数字
	ContactNo             string `url:"contact_no,omitempty" validate:"max=64"`          // 支持数字及特殊字符+-()
	ContactEmail          string `url:"contact_email,omitempty" validate:"omitempty,email,max=128"`
	CsNo                  string `url:"cs_no,omitempty" validate:"max=64"`                     // 客服电话
	CsEmail               string `url:"cs_email,omitempty" validate:"omitempty,email,max=128"` // 客户邮箱
}

// 支付宝 线下 线上
type RegisterResp struct {
	IsSuccess  string `xml:"is_success"`            // T F
	SignType   string `xml:"sign_type,omitempty"`   // MD5 RSA RSA2
	Sign       string `xml:"sign,omitempty"`        // 签名值
	Error      string `xml:"error,omitempty"`       // 错误码
	ResultCode string `xml:"result_code,omitempty"` // 当 is_success字段的值为T时，返回此字段 eg.SUCCESS
}

// 注册接口返回成功，代表支付宝接受信息 还在审核中
func (r *RegisterResp) Describe() core.CommonResp {
	resp := core.CommonResp{
		Status:     consts.Fail,
		ResultCode: r.Error,
		Message:    r.ResultCode,
	}
	if r.IsSuccess == success {
		resp.Status = consts.UnderReview
	}
	return resp
}

// 支付宝线下
type QueryStatusOfflineResp struct {
	RegisterResp
	SecondaryMerchantID     string `xml:"secondary_merchant_id"`
	StoreID                 string `xml:"store_id"`
	Status                  string `xml:"status"` // 枚举 SUCCESS UNDER_REVIEW MODIFICATION_REVIEW FAILED
	RejectReason            string `xml:"reject_reason"`
	PaymentMethod           string `xml:"payment_method"`          // 二级商户支付方式 枚举 INSTORE_PAYMENT
	Partner                 string `xml:"partner"`                 // PID
	SecondaryMerchantName   string `xml:"secondary_merchant_name"` // 二级商户的法定注册名称
	SecondaryMerchantType   string `xml:"secondary_merchant_type"` // 枚举 INDIVIDUAL ENTERPRISE
	RegistrationNo          string `xml:"registration_no,omitempty"`
	RegistrationCountry     string `xml:"register_country"`           // 2位代码 eg.HK
	RegisterAddress         string `xml:"register_address,omitempty"` // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	ShareholderName         string `xml:"shareholder_name,omitempty"`
	ShareholderID           string `xml:"shareholder_id,omitempty"`
	RepresentativeName      string `xml:"representative_name,omitempty"`
	RepresentativeID        string `xml:"representative_id,omitempty"`
	SettlementNo            string `xml:"settlement_no,omitempty"`
	ContactNo               string `xml:"contact_no,omitempty"`
	ContactEmail            string `xml:"contact_email,omitempty"`
	CsNo                    string `xml:"cs_no,omitempty"`
	CsEmail                 string `xml:"cs_email,omitempty"`
	StoreName               string `xml:"store_name"`     // 商店名称 车牌号
	StoreCountry            string `xml:"store_country"`  // 2位代码 eg.HK
	StoreAddress            string `xml:"store_address"`  // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	StoreIndustry           string `xml:"store_industry"` // MCC 4位代码 eg. 4121
	InternalStorePhoto      string `xml:"internal_store_photo,omitempty"`
	ExternalStorefrontPhoto string `xml:"external_storefront_photo,omitempty"`
	StoreStatus             string `xml:"store_status"` // 枚举 ARCHIVED ACTIVATED INACTIVATE
}

func (q QueryStatusOfflineResp) Describe() core.CommonResp {
	resp := q.RegisterResp.Describe()
	if q.IsSuccess == success {
		resp.Status = consts.Success
	}
	resp.StoreStatus = q.StoreStatus
	resp.SubMerchantID = q.SecondaryMerchantID
	resp.Message = q.RejectReason

	return resp
}

// 支付宝线上
type QueryStatusOnlineResp struct {
	RegisterResp
	SecondaryMerchantID       string `xml:"secondary_merchant_id"`
	Status                    string `xml:"status"`
	RejectReason              string `xml:"reject_reason,omitempty"`
	PaymentMethod             string `xml:"payment_method,omitempty"`    // 二级商户支付方式 枚举 INSTORE_PAYMENT
	Partner                   string `xml:"partner"`                     // PID
	SecondaryMerchantName     string `xml:"secondary_merchant_name"`     // 二级商户的法定注册名称
	SecondaryMerchantType     string `xml:"secondary_merchant_type"`     // 枚举 INDIVIDUAL ENTERPRISE
	SecondaryMerchantIndustry string `xml:"secondary_merchant_industry"` // MCC 四位
	RegistrationNo            string `xml:"registration_no,omitempty"`
	RegistrationCountry       string `xml:"register_country"`           // 2位代码 eg.HK
	RegisterAddress           string `xml:"register_address,omitempty"` // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	ShareholderName           string `xml:"shareholder_name,omitempty"`
	ShareholderID             string `xml:"shareholder_id,omitempty"`
	RepresentativeName        string `xml:"representative_name,omitempty"`
	RepresentativeID          string `xml:"representative_id,omitempty"`
	SettlementNo              string `xml:"settlement_no,omitempty"`
	ContactNo                 string `xml:"contact_no,omitempty"`
	ContactEmail              string `xml:"contact_email,omitempty"`
	CsNo                      string `xml:"cs_no,omitempty"`
	CsEmail                   string `xml:"cs_email,omitempty"`
	StoreName                 string `xml:"store_name"`     // 商店名称 车牌号
	StoreCountry              string `xml:"store_country"`  // 2位代码 eg.HK
	StoreAddress              string `xml:"store_address"`  // 邮政地址格式 eg. No.276, Road YinCheng, Shanghai
	StoreIndustry             string `xml:"store_industry"` // MCC 4位代码 eg. 4121
	InternalStorePhoto        string `xml:"internal_store_photo,omitempty"`
	ExternalStorefrontPhoto   string `xml:"external_storefront_photo,omitempty"`
	StoreStatus               string `xml:"store_status"` // 枚举 ARCHIVED ACTIVATED INACTIVATE
}

func (q QueryStatusOnlineResp) Describe() core.CommonResp {
	resp := q.RegisterResp.Describe()
	if q.IsSuccess == success {
		resp.Status = consts.Success
	}
	resp.StoreStatus = q.StoreStatus
	resp.SubMerchantID = q.SecondaryMerchantID
	resp.Message = q.RejectReason

	return resp
}
