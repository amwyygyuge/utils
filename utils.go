package pkg

import (
	"math/rand"

	pkgConstans "github.com/amwyygyuge/utils/constant"
)

func GetRandomUserAgent() string {
	return pkgConstans.USER_AGENTS[rand.Intn(len(pkgConstans.USER_AGENTS))]
}
