package ascii

import (
	"errors"
	"strings"
)

// EqualFold is strings.EqualFold, ASCII only. It reports whether s and t
// are equal, ASCII-case-insensitively.
func EqualFold(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if lower(s[i]) != lower(t[i]) {
			return false
		}
	}
	return true
}

// lower returns the ASCII lowercase version of b.
func lower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + ('a' - 'A')
	}
	return b
}

func upper(str string) string {
	return strings.ToUpper(str)
}

// user_name to ID
func GetID(user_name string) (string, error) {
	if len(user_name) == 0 {
		return "", errors.New("string is empty")
	}
	return upper(user_name[0:5]), nil
}
