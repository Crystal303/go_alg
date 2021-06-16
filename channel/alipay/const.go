package alipay

// service 默认值
const (
	ServiceOnlineQueryStatus = "alipay.overseas.secmerchant.maintain.queryStatus"
	ServiceOnlineRegister    = "alipay.overseas.secmerchant.online.maintain"

	ServiceOfflineQueryStatus = "alipay.overseas.secmerchant.maintain.queryStatus"
	ServiceOfflineRegister    = "alipay.overseas.secmerchant.offline.maintain"

	alipayOfflineUrl = ""
	alipayOnlineUrl  = ""
)

// 添加到 url 的另外信息
const (
	signType = "sign_type"
	sign     = "sign"
)

// Input_charset
const (
	charsetDefault = "UTF-8"
	signDefault    = "MD5"
	dateFormat     = "2006-01-02 15:04:05"
	tagName        = "url"
)

const (
	success = "T"
	fail    = "F"

	paymentMethodOnline  = "ONLINE_PAYMENT"
	paymentMethodOffline = "INSTORE_PAYMENT"
)

const (
	mccDriver = "4121"
)
