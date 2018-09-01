package common

import (
	"fmt"
	"strings"
)

func SplitSpacesWithQuotes(s string) ([]string, error) {
	var ret []string
	var buf = new(strings.Builder) // not in quote string buffer
	var scanned string
	var err error

	for pos := 0; pos < len(s); pos++ {
		char := s[pos]

		switch char {
		case ' ':
			if buf.Len() > 0 {
				ret = append(ret, buf.String())
				buf.Reset()
			}
		case '"', '\'':
			pos, scanned, err = scanForByte(s, pos, char)
			if err != nil {
				return nil, err
			}
			ret = append(ret, scanned)
		default:
			buf.WriteByte(char)
		}
	}

	if buf.Len() > 0 {
		ret = append(ret, buf.String())
	}

	return ret, nil
}

func scanForByte(s string, pos int, r byte) (int, string, error) {
	var ret = new(strings.Builder)
	for pos++; pos < len(s); pos++ {
		char := s[pos]

		switch char {
		case '\\':
			if pos >= len(s)-1 {
				return 0, "", fmt.Errorf("unbalanced quotes")
			}
			pos++
			ret.WriteByte(s[pos])
		case r:
			return pos, ret.String(), nil
		default:
			ret.WriteByte(char)
		}
	}

	return 0, "", fmt.Errorf("unbalanced quotes")
}
