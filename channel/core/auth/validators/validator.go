package validators

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// 用于验证签名信息
type Validator interface {
	ValidateResp(ctx context.Context, response *http.Response) error // 对 HTTP 应答报文进行验证
}

type nullValidator struct{}

var NullValidator nullValidator

// Validate 跳过报文签名验证
func (validator nullValidator) ValidateResp(ctx context.Context, response *http.Response) error {
	return nil
}

// 用于验证组装的请求结构体
type ValidateStruct interface {
	Validate(interface{}) error // 只校验结构体
}

type validate struct{}

var DefaultValidate validate

func (v validate) Validate(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	return err
}
