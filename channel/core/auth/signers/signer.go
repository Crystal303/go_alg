package signers

import "context"

type Signer interface {
	Sign(context.Context, interface{}) (string, error) // 对信息进行签名
}
