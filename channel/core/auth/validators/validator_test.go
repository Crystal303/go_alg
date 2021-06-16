package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validate_Validate(t *testing.T) {
	type temp struct {
		BusinessType string `validate:"required,oneof=ONLINE OFFLINE BOTH"` // 业务类型
		// ONLINE BOTH 时 以下四个至少传入一项
		AppDownload     string `validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=BusinessWebsite OfficeAccount MiniProgram"` // 商户APP的下载地址
		BusinessWebsite string `validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload OfficeAccount MiniProgram"`     // 业务网站
		OfficeAccount   string `validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload BusinessWebsite MiniProgram"`   // 公众号
		MiniProgram     string `validate:"max=128,required_unless=BusinessType OFFLINE|required_without_all=AppDownload BusinessWebsite OfficeAccount"` // 小程序
	}

	tests := []temp{
		{
			BusinessType:    "ONLINE",
			AppDownload:     "https://download.qq.com",
			BusinessWebsite: "",
			OfficeAccount:   "",
			MiniProgram:     "",
		},
		{
			BusinessType:    "ONLINE",
			AppDownload:     "https://download.qq.com",
			BusinessWebsite: "https://download.qq.com",
			OfficeAccount:   "https://download.qq.com",
			MiniProgram:     "https://download.qq.com",
		},
		{
			BusinessType:    "ONLINE",
			AppDownload:     "https://download.qq.com",
			BusinessWebsite: "https://download.qq.com",
			OfficeAccount:   "https://download.qq.com",
			MiniProgram:     "",
		},
		{
			BusinessType:    "ONLINE",
			AppDownload:     "https://download.qq.com",
			BusinessWebsite: "https://download.qq.com",
			OfficeAccount:   "",
			MiniProgram:     "",
		},
		{
			BusinessType:    "OFFLINE",
			AppDownload:     "https://download.qq.com",
			BusinessWebsite: "https://download.qq.com",
			OfficeAccount:   "",
			MiniProgram:     "",
		},
	}
	for _, v := range tests {
		err := DefaultValidate.Validate(v)
		assert.NoError(t, err)
	}

	errTest := temp{
		BusinessType:    "BOTH",
		AppDownload:     "",
		BusinessWebsite: "",
		OfficeAccount:   "",
		MiniProgram:     "",
	}
	err := DefaultValidate.Validate(errTest)
	assert.Error(t, err)
}

func TestValidate_Validate(t *testing.T) {
	type temp struct {
		Data string `validate:"datetime=2006-01-02|eq=N/A"`
	}

	test := temp{Data: "N/A"}
	err := DefaultValidate.Validate(test)
	assert.NoError(t, err)

	test = temp{Data: "2020-11-11"}
	err = DefaultValidate.Validate(test)
	assert.NoError(t, err)
}
