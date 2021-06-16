package core

import (
	"net/http"
	"time"

	"github.com/Crystal303/go_alg/channel/core/auth/validators"
)

// 这是一个综合的结构体
// Country 统一为3位字母
type MerchantInfo struct {
	ChanMerID    string
	SubChanMerID string
	// 门店信息
	StoreID                 string
	StoreName               string
	StoreCountry            string
	StoreAddress            string
	StoreMCC                string
	InternalStorePhoto      string // URL
	ExternalStorefrontPhoto string // URL
	ExtendParams            []ExtendParam
	// 商户信息
	SecondaryMerchantType string
	SecondaryMerchantName string
	RegistrationNo        string
	RegisterCountry       string
	RegisterAddress       string
	ShareholderName       string
	ShareholderID         string
	RepresentativeName    string
	RepresentativeID      string
	SettlementNo          string     // 子商户的结算银行账户信息
	ContactNo             string     // 联系电话
	ContactEmail          string     // 联系邮箱
	CsNo                  string     // 客服电话
	CsEmail               string     // 客服邮箱
	SecondaryMerchantMCC  string     // 商户MCC 线上使用
	SiteInfos             []SiteInfo // 网站信息 线上使用
	// wechat
	AppID           string
	BusinessType    string
	AppDownload     string
	BusinessWebsite string
	OfficeAccount   string
	OfficePhone     string
	MiniProgram     string
	StorePhotos     []string // Media ID 数组
	// 门店
	StoreNameEN                 string
	StoreNum                    string
	RegistrationCertificateDate string // 公司注册文件过期时间 YYYY-MM-DD
	ContactName                 string // 联系人姓名
}

// 出租车信息
type ExtendParam struct {
	OperationID   string `json:"operation_id"`          // M
	ContactWay    string `json:"contact_way,omitempty"` // O
	ContactPerson string `json:"contact_person"`        // M
}

// 网站信息
type SiteInfo struct {
	SiteType string `json:"site_type"`           // M
	SiteUrl  string `json:"site_url"`            // M
	SiteName string `json:"site_name,omitempty"` // O
}

// 通用返回信息
type CommonResp struct {
	SubMerchantID string // 二级商户号
	Status        string // 入网处理状态 枚举 SUCCESS UNDER_REVIEW MODIFICATION_REVIEW FAILED
	ResultCode    string // 返回的信息码
	StoreStatus   string // 门店状态 枚举 ARCHIVED ACTIVATED INACTIVATE
	Message       string // 返回的信息详情
}

// 渠道配置，区分渠道个性化配置信息
type PspSetting struct {
	SignKey               string
	ClientCert, ClientKey []byte // HTTPS 加密证书
	PortalCert            []byte // 微信获取的平台证书(公钥)
}

// http 配置
type dialSettings struct {
	HTTPClient    *http.Client         // 自定义所使用的 HTTPClient 实例
	Header        http.Header          // 自定义额外请求头
	PostValidator validators.Validator // 应答包签名校验器
	Timeout       time.Duration        // HTTP 请求超时时间，将覆盖 HTTPClient 中的 Timeout
	// Signer        signers.Signer            // 生成签名
	PreValidator validators.ValidateStruct // 前置结构体校验
}
