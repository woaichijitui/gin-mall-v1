package cache

import (
	"fmt"
	"strconv"
)

var (
	RangKey = "rank"
)

// 根据pId 返回一个key
func ProductViemKey(pId uint) string {
	return fmt.Sprintf("viem:product:%s", strconv.Itoa(int(pId)))
}
