package alipay

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Crystal303/go_alg/channel/core"
	"github.com/Crystal303/go_alg/channel/core/query"
)

// data 为 nil
func processResponse(resp *http.Response, data interface{}) error {
	defer resp.Body.Close()

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := xml.Unmarshal(bts, &data); err != nil {
		return err
	}
	return nil
}

func checkPSPSetting(setting core.PspSetting) error {
	if strings.TrimSpace(setting.SignKey) == "" {
		return fmt.Errorf("empty %s", "SignKey")
	}
	return nil
}

func genURI(url, signKey string, data interface{}) (string, error) {
	values, err := query.Values(data, tagName)
	if err != nil {
		return "", err
	}

	content := query.Encode(values)

	signed := md5.Sum([]byte(content + signKey))

	values.Add(sign, hex.EncodeToString(signed[:]))
	values.Add(signType, signDefault)

	return url + "?" + values.Encode(), nil
}

// 国际信息转为ISO 3166格式 3位字母到2位字母
func convertCountry(countryCode string) string {
	panic("todo")
}
