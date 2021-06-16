package consts

// HTTP 请求报文 Header 常量
const (
	Authorization  = "Authorization"  // Header 中的 Authorization 字段
	Accept         = "Accept"         // Header 中的 Accept 字段
	ContentType    = "Content-Type"   // Header 中的 ContentType 字段
	ContentLength  = "Content-Length" // Header 中的 ContentLength 字段
	UserAgent      = "User-Agent"     // Header 中的 UserAgent 字段
	AcceptLanguage = "en"
)

const (
	ApplicationJSON  = "application/json"
	XmlContentType   = "application/xml; charset=utf-8"
	UserAgentContent = "CIL-HttpClient/v2"
)

// 返回结构体 Status 枚举值
const (
	Success            = "SUCCESS"
	Fail               = "FAILED"
	UnderReview        = "UNDER_REVIEW"
	ModificationReview = "MODIFICATION_REVIEW"
)

const (
	MerchantTypeINDIVIDUAL = "INDIVIDUAL"
	MerchantTypeENTERPRISE = "ENTERPRISE"

	BusinessTypeBOTH    = "BOTH"
	BusinessTypeONLINE  = "ONLINE"
	BusinessTypeOFFLINE = "OFFLINE"
)
