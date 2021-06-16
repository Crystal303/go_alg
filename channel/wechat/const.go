package wechat

// wechat 接口uri
const (
	ServiceQueryStatusV2 = "/secapi/mch/queryInstitutionsub"
	ServiceRegisterV2    = "/secapi/mch/addInstitutionsub"
	ServiceModifyV2      = "/secapi/mch/modifyInstitutionsub"

	ServiceRegisterV3    = "/hk/v3/merchants"
	ServiceModifyV3      = "/hk/v3/merchants"
	ServiceQueryStatusV3 = "/hk/v3/merchants/%s" // /hk/v3/merchants/{sub_mchid}

	UrlWechatV2 = ""
	UrlWechatV3 = ""
)

const (
	VerificationStatusUnderReview = "Under Review"
	VerificationStatusApproved    = "Approved"
)

const (
	ResultCodeSuccess = "SUCCESS"
	ResultCodeFail    = "FAIL"
)

const (
	ReturnCodeSuccess = ResultCodeSuccess
	ReturnCodeFail    = ResultCodeFail
)

const (
	httpPOSTMethod = "POST"
	httpPUTMethod  = "PUT"
	httpGETMethod  = "GET"

	//RegisterUrl    = "https://api.mch.weixin.qq.com/secapi/mch/addInstitutionsub"
	//QueryStatusUrl = "https://api.mch.weixin.qq.com/secapi/mch/queryInstitutionsub"
	//ModifyUrl      = "https://api.mch.weixin.qq.com/secapi/mch/modifyInstitutionsub"
)

// 请求报文签名相关常量
const (
	NonceSymbols           = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 随机字符串可用字符集
	NonceLength            = 32                                                               // 随机字符串的长度
	SignatureMessageFormat = "%s\n%s\n%d\n%s\n%s\n"                                           // 数字签名原文格式
	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "WECHATPAY2-SHA256-RSA2048 mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
)

const (
	tagName = "xml"
)

// HTTP 应答报文 Header 相关常量
const (
	WechatPayTimestamp = "Wechatpay-Timestamp" // 微信支付回包时间戳
	WechatPayNonce     = "Wechatpay-Nonce"     // 微信支付回包随机字符串
	WechatPaySignature = "Wechatpay-Signature" // 微信支付回包签名信息
	WechatPaySerial    = "Wechatpay-Serial"    // 微信支付回包平台序列号
	RequestID          = "Request-Id"          // 微信支付回包请求ID
)
