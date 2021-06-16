package wechat

import (
	"fmt"

	"github.com/Crystal303/go_alg/channel/consts"
	"github.com/Crystal303/go_alg/channel/core"
)

// 查询子商户审核状态
type QueryStatusReq struct {
	AppID    string `xml:"app_id" json:"app_id" validate:"required,max=32"`                 // 微信分配的公众账号ID
	MchID    string `xml:"mch_id" json:"mch_id" validate:"required,max=32"`                 // 微信支付分配的商户号
	Sign     string `xml:"sign,omitempty" json:"sign,omitempty" validate:"required,max=32"` // 签名
	SubMchID string `xml:"sub_mch_id" json:"sub_mch_id" validate:"required,max=32"`
}

// 子商户注册时公共参数
type MerchantParams struct {
	AppID               string `xml:"app_id" json:"app_id" validate:"required,max=32"`                 // 微信分配的公众账号ID
	MchID               string `xml:"mch_id" json:"mch_id" validate:"required,max=32"`                 // 微信支付分配的商户号
	Sign                string `xml:"sign,omitempty" json:"sign,omitempty" validate:"required,max=32"` // 签名
	ChannelID           string `xml:"channel_id,omitempty" json:"channel_id,omitempty" validate:"max=20"`
	MerchantName        string `xml:"merchant_name" json:"merchant_name" validate:"required,max=20"`
	MerchantShortName   string `xml:"merchant_shortname" json:"merchant_shortname" validate:"required,max=64"`
	MerchantCountryCode string `xml:"merchant_country_code" json:"merchant_country_code" validate:"required,len=3,number"` // 三位国家代码 数字 eg.344
	MerchantType        string `xml:"merchant_type" json:"merchant_type" validate:"required,oneof=ENTERPRISE INDIVIDUAL"`
	BusinessCategory    string `xml:"business_category" json:"business_category" validate:"required,len=3,number"` // 业务类目
	// https://pay.weixin.qq.com/wiki/doc/api/wxpay/ch/sub_merchant_entry/chapter2_3.shtml#part-8

	Mcc                           string `xml:"mcc" json:"mcc" validate:"required,len=4,number"`                                                                                                                         // mcc
	RegistrationCertificateNumber string `xml:"registration_certificate_number,omitempty" json:"registration_certificate_number,omitempty" validate:"max=10,required_if=MerchantType ENTERPRISE"`                        // 公司注册文件编号
	RegistrationCertificateDate   string `xml:"registration_certificate_date,omitempty" json:"registration_certificate_date,omitempty" validate:"max=10,datetime=2006-01-02|eq=N/A,required_if=MerchantType ENTERPRISE"` // 公司注册文件过期时间 YYYY-MM-DD PERMANENT N/A
	RegistrationCertificateCopy   string `xml:"registration_certificate_copy,omitempty" json:"registration_certificate_copy,omitempty" validate:"max=128"`                                                               // 公司注册文件的照片，取值为上传文件API返回的media ID

	BusinessType string `xml:"business_type" json:"business_type" validate:"required,oneof=ONLINE OFFLINE BOTH"` // 业务类型
	// ONLINE BOTH 时 以下四个至少传入一项
	AppDownload     string `xml:"app_download,omitempty" json:"app_download,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=BusinessWebsite OfficeAccount MiniProgram"`     // 商户APP的下载地址
	BusinessWebsite string `xml:"business_website,omitempty" json:"business_website,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload OfficeAccount MiniProgram"` // 业务网站
	OfficeAccount   string `xml:"office_account,omitempty" json:"office_account,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload BusinessWebsite MiniProgram"`   // 公众号
	MiniProgram     string `xml:"mini_program,omitempty" json:"mini_program,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload BusinessWebsite OfficeAccount"`     // 小程序
	StoreAddress    string `xml:"store_address,omitempty" json:"store_address,omitempty" validate:"max=128,required_unless=BusinessType ONLINE"`                                                                   // 门店地址
	StorePhotos     string `xml:"store_photos,omitempty" json:"store_photos,omitempty" validate:"max=1024,required_unless=BusinessType ONLINE"`                                                                    // 商户门店照片
	// 当业务类型取值为OFFLINE或BOTH时必传。至少上传三张门店照片。取值为《图片上传API》返回的media ID.
	// 多个图片，使用Json数组格式提交.
	// 示例值：[Media_id1, Media_id2, Media_id3]
	DirectorName         string `xml:"director_name,omitempty" json:"director_name,omitempty" validate:"max=128,required_if=MerchantType ENTERPRISE"`             // 董事姓名
	DirectorIDNumber     string `xml:"director_id_number,omitempty" json:"director_id_number,omitempty" validate:"max=128,required_if=MerchantType ENTERPRISE"`   // 董事证件号码
	PrincipalName        string `xml:"principal_name,omitempty" json:"principal_name,omitempty" validate:"max=128,required_if=MerchantType INDIVIDUAL"`           // 负责人姓名
	PrincipalIDNumber    string `xml:"principal_id_number,omitempty" json:"principal_id_number,omitempty" validate:"max=128,required_if=MerchantType INDIVIDUAL"` // 负责人证件号码
	OfficePhone          string `xml:"office_phone" json:"office_phone" validate:"max=32"`
	ContactName          string `xml:"contact_name" json:"contact_name" validate:"max=64"`
	ContactPhone         string `xml:"contact_phone" json:"contact_phone" validate:"max=32"`
	ContactEmail         string `xml:"contact_email" json:"contact_email" validate:"max=256,omitempty,email"`
	SettlementBankNumber string `xml:"settlement_bank_number" json:"settlement_bank_number" validate:"max=128"`
}

// 子商户注册请求
type RegisterReq struct {
	MerchantParams
	MerchantRemark string `xml:"merchant_remark" json:"merchant_remark" validate:"required,max=20"`
}

// 子商户修改请求
type ModifyReq struct {
	MerchantParams
	SubMchID string `xml:"sub_mch_id" json:"sub_mch_id" validate:"required,max=32"`
}

// 应答基本返回字段
type BaseResp struct {
	ReturnCode string `xml:"return_code" json:"return_code"` // 通信标识
	ReturnMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"`
	ResultCode string `xml:"result_code" json:"result_code"` // 业务标识
	ErrCode    string `xml:"err_code,omitempty" json:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`
	SubMchID   string `xml:"sub_mch_id" json:"sub_mch_id"`
	Sign       string `xml:"sign" json:"sign"`
}

func (b BaseResp) Describe() core.CommonResp {
	data := core.CommonResp{
		SubMerchantID: b.SubMchID,
		Status:        consts.Fail,
		ResultCode:    b.ResultCode,
		StoreStatus:   "",
		Message:       "",
	}
	if b.ReturnCode == ReturnCodeSuccess && b.ResultCode == ResultCodeSuccess {
		data.Status = consts.Success
	} else {
		data.Message = fmt.Sprintf("return_msg: %s, err_code: %s, description: %s", b.ReturnMsg, b.ErrCode, b.ErrCodeDes)
	}
	return data
}

// 注册返回应答
type RegisterResp struct {
	BaseResp
	VerificationStatus string `xml:"verification_status,omitempty" json:"verification_status,omitempty"` // 验证状态
	Description        string `xml:"description,omitempty" json:"description,omitempty"`
}

func (r RegisterResp) Describe() core.CommonResp {
	data := core.CommonResp{
		SubMerchantID: r.SubMchID,
		Status:        consts.Fail,
		ResultCode:    r.ResultCode,
		StoreStatus:   "",
		Message:       "",
	}
	if r.ReturnCode == ReturnCodeSuccess && r.ResultCode == ResultCodeSuccess {
		data.Status = consts.UnderReview
		if r.VerificationStatus == VerificationStatusApproved {
			data.Status = consts.Success
		}
	} else {
		data.Message = fmt.Sprintf("return_msg: %s, err_code: %s, description: %s", r.ReturnMsg, r.ErrCode, r.Description)
	}
	return data
}

// 修改返回应答
type ModifyResp struct {
	BaseResp
}

// 查询返回应答
// 查询返回商户注册信息相关字段，此处解析进行省略
type QueryStatusResp struct {
	BaseResp
}

type QueryStatusV3Req struct {
	SpAppID  string `json:"sp_appid" validate:"required,max=32"`  // 机构APP ID
	SpMchID  string `json:"sp_mchid" validate:"required,max=32"`  // 机构商户号
	SubMchID string `json:"sub_mchid" validate:"required,max=32"` // 子商户号
}

type RegisterV3Req struct {
	SpAppID                       string     `json:"sp_appid" validate:"required,max=32"`                                                                                      // 机构APP ID
	SpMchID                       string     `json:"sp_mchid" validate:"required,max=32"`                                                                                      // 机构商户号
	Name                          string     `json:"name" validate:"required,max=128"`                                                                                         // 子商户全称
	ShortName                     string     `json:"shortname" validate:"required,max=64"`                                                                                     // 子商户简称
	OfficePhone                   string     `json:"office_phone" validate:"required,max=32"`                                                                                  // 公司电话
	Contact                       Contact    `json:"contact" validate:"required"`                                                                                              // 联需人信息
	BusinessCategory              int        `json:"business_category" validate:"required,len=3"`                                                                              // 类目
	ChannelID                     string     `json:"channel_id,omitempty" validate:"max=20"`                                                                                   // 渠道号
	MerchantCountryCode           string     `json:"merchant_country_code" validate:"required,len=3"`                                                                          // 国家区域
	MerchantType                  string     `json:"merchant_type" validate:"required,oneof=ENTERPRISE INDIVIDUAL"`                                                            // 商户类型
	RegistrationCertificateNumber string     `json:"registration_certificate_number,omitempty" validate:"max=50,required_if=MerchantType ENTERPRISE"`                          // 商业登记号
	RegistrationCertificateDate   string     `json:"registration_certificate_date,omitempty" validate:"max=10,datetime=2006-01-02|eq=N/A,required_if=MerchantType ENTERPRISE"` // 商业证书有效期
	RegistrationCertificateCopy   string     `json:"registration_certificate_copy,omitempty" validate:"max=128"`                                                               // 公司注册文件的照片，取值为上传文件API返回的media ID
	SettlementBankNumber          string     `json:"settlement_bank_number,omitempty" validate:"max=128"`                                                                      // 开户银行编号
	Business                      Business   `json:"business" validate:"required"`
	Director                      *Director  `json:"director,omitempty" validate:"required_if=MerchantType ENTERPRISE"`
	Principal                     *Principal `json:"principal,omitempty" validate:"required_if=MerchantType INDIVIDUAL"`
}

type Business struct {
	BusinessType  string `json:"business_type" validate:"required,oneof=ONLINE OFFLINE BOTH"`                                                                                   // 业务类型
	AppDownload   string `json:"app_download,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=BusinessWebsite OfficeAccount MiniProgram"` // 商户APP的下载地址
	Website       string `json:"website,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload OfficeAccount MiniProgram"`          // 业务网站
	OfficeAccount string `json:"office_account,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload BusinessWebsite MiniProgram"` // 公众号
	MiniProgram   string `json:"mini_program,omitempty" validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload BusinessWebsite OfficeAccount"` // 小程序
	StoreAddress  string `json:"store_address,omitempty" validate:"max=128,required_unless=BusinessType ONLINE"`                                                                // 门店地址
	StorePhotos   string `json:"store_photos,omitempty" validate:"max=1024,required_unless=BusinessType ONLINE"`
	Mcc           string `xml:"mcc" json:"mcc" validate:"required,len=4,number"` // mcc
}

type Director struct {
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
}

type Principal struct {
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
}

type Contact struct {
	Name  string `json:"name" validate:"required,max=64"`
	Phone string `json:"phone" validate:"required,max=32"`
	Email string `json:"email" validate:"required,email,max=256"`
}

type ModifyV3Req struct {
	SubMchID string `json:"sub_mchid" validate:"required,max=32"` // 子商户号
	RegisterV3Req
}

type RespV3 struct {
	Code    string  `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Detail  *Detail `json:"detail,omitempty"`

	// 正常返回
	SubMchID           string `json:"sub_mchid,omitempty"`           // 子商户号
	VerificationStatus string `json:"verification_status,omitempty"` // 验证状态
	Description        string `json:"description,omitempty"`
}

func (r RespV3) Describe() core.CommonResp {
	data := core.CommonResp{
		SubMerchantID: r.SubMchID,
		Status:        consts.Fail,
		ResultCode:    r.Code,
		StoreStatus:   "",
	}
	if r.SubMchID != "" {
		data.Status = consts.UnderReview
	}
	if r.Detail != nil {
		data.Message = fmt.Sprintf("field: %s, value: %s, issue: %s, location: %s", r.Detail.Field, r.Detail.Value, r.Detail.Issue, r.Detail.Location)
	}
	return data
}

type QueryStatusV3Resp struct {
	RespV3
}

func (r QueryStatusV3Resp) Describe() core.CommonResp {
	data := r.RespV3.Describe()
	if r.SubMchID != "" {
		data.Status = consts.Success
	}
	return data
}

type Detail struct {
	Field    string `json:"field"`
	Value    string `json:"value"`
	Issue    string `json:"issue"`
	Location string `json:"location,omitempty"`
}
