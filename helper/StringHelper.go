package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func JoinStrings(args ...interface{}) string {
	var builder strings.Builder

	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			builder.WriteString(strconv.Itoa(v))
		case string:
			builder.WriteString(v)
		default:
			builder.WriteString(fmt.Sprint(v))
		}
	}

	return builder.String()
}
