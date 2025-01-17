package go360

import (
	"fmt"
	"strconv"
	"strings"
)

func ToHex(str string) string {
	var result string
	for _, c := range str {
		result += fmt.Sprintf("%02x", c)
	}
	return result
}

func ParseHexToDec(hexStr string) (int, error) {
	hexStr = strings.TrimSpace(hexStr)
	parsedInt, err := strconv.ParseInt(hexStr, 16, 0)
	if err != nil {
		return 0, err
	}
	return int(parsedInt), nil
}
