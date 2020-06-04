package query

import "strings"

type values struct {
	strings.Builder
}

func (v *values) add(key, value string) {
	v.WriteString(key)
	v.WriteByte('=')

	n := len(value)
	if n > 0 {
		// Hint the compiler to remove bounds checks in the loop below.
		_ = value[n-1]
	}
	for i := 0; i < n; i++ {
		c := value[i]

		// See http://www.w3.org/TR/html5/forms.html#form-submission-algorithm
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' ||
			c == '-' || c == '.' || c == '_' {
			v.WriteByte(c)
		} else {
			if c == ' ' {
				v.WriteByte('+')
			} else {
				v.Write([]byte{'%', hexCharUpper(c >> 4), hexCharUpper(c & 15)})
			}
		}
	}
	v.WriteByte('&')
}

func hexCharUpper(c byte) byte {
	if c < 10 {
		return '0' + c
	}
	return c - 10 + 'A'
}
