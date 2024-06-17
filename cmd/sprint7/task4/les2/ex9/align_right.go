package ex9

import (
	"bytes"
	"strings"
)

func AlignRightAppend(s string, length int, lead rune) string {
	for len(s) < length {
		s = string(lead) + s
	}
	return s
}

func AlignRightBuffer(s string, length int, lead rune) string {
	buf := bytes.Buffer{}
	for i := 0; i < length-len(s); i++ {
		buf.WriteRune(lead)
	}
	buf.WriteString(s)
	return buf.String()
}

func AlignRightRepeat(s string, length int, lead rune) string {
	if len(s) < length {
		return strings.Repeat(string(lead), length-len(s)) + s
	}
	return s
}
