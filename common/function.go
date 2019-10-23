package common

import (
	"hash/fnv"
	"strconv"
)

// HashText: common hash text function
func HashText(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return strconv.FormatUint(uint64(h.Sum32()), 16)
}
