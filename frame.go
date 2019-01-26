package frame

import (
	"fmt"
	"strings"
)

const (
	frameChar = "#"
)

// Print prints i in a frame of #
func Print(i interface{}) {
	var (
		str    = fmt.Sprint(i)
		strLen = len(str)
	)

	if strLen == 0 {
		return
	}

	fmt.Println(strings.Repeat(frameChar, strLen+4))
	fmt.Printf("%s %s %s\n", frameChar, str, frameChar)
	fmt.Println(strings.Repeat(frameChar, strLen+4))
}
