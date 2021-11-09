package tools

import (
	"strconv"
	"strings"
)

func GetVersionNum(ver string) int {
	digs := strings.Split(ver, ".")
	a, _ := strconv.Atoi(digs[0])
	b, _ := strconv.Atoi(digs[1])
	return a*10 + b
}
