package core

import "context"

// 子商户自动入网接入
type SecMerchantRegister interface {
	// 注册
	Register(context.Context, SecMerchantDescriber, ...RegisterOption) (RespBuilder, error)
	// 查询注册/更新结果
	QueryStatus(context.Context, SecMerchantDescriber, ...RegisterOption) (RespBuilder, error)
	// 更新
	Modify(context.Context, SecMerchantDescriber, ...RegisterOption) (RespBuilder, error)
}

// PSP 进行自动入网需要实现以下接口方法
type SecMerchantDescriber interface {
	Describe() MerchantInfo
}

// PSP 自动入网返回结果处理
type RespBuilder interface {
	Describe() CommonResp
}

// 用于 渠道配置
type RegisterOption interface {
	Apply(*PspSetting)
}

// 用于 http 客户端
type ClientOption interface {
	Apply(*dialSettings)
}
