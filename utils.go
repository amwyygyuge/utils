package utils

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	pkgConstans "github.com/amwyygyuge/utils/constant"
)

func ZhToUnicode(raw []byte) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	str = strings.ReplaceAll(str, `\/`, "/")
	return str, nil
}

func GetRandomUserAgent() string {
	return pkgConstans.USER_AGENTS[rand.Intn(len(pkgConstans.USER_AGENTS))]
}
func GetProxyIps() pkgConstans.Ips {
	res, _ := http.Get("https://ip.jiangxianli.com/api/proxy_ips")
	result, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	str, _ := ZhToUnicode(result)
	var data pkgConstans.Ips
	json.Unmarshal([]byte(str), &data)
	return data
}
