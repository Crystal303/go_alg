// Package wechat 存量支持微信V2接口（含微信V2子商户入网接口）的基础上 需支持了微信V3接口（含微信V3子商户入网接口）
// V2 V3 根据微信资质 MID(Merchant ID) 进行区分，子商户号在注册成功后返回
// MID 在机构层创建 密钥信息也保存在机构层
// V2版接口和V3版接口实际上是基于两种接口标准设计的两套接口。
//	V3 					  		规则差异 	V2
//	JSON 					  	参数格式		XML
//	POST、GET、PATCH、DELETE 	提交方式		POST
//	AES-256-GCM加密 				回调加密		无需加密
//	RSA加密 						敏感加密		无需加密
//	UTF-8 						编码方式		UTF-8
//	非对称密钥SHA256-RSA 			签名方式		MD5或HMAC-SHA256

package wechat
