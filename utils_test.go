package utils

import (
	"fmt"
	"testing"
)

func TestGetProxyIps(t *testing.T) {
	a, _ := GetProxyIps()
	fmt.Println(a)

}
