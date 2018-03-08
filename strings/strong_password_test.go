package strings

import (
	"testing"
	"strings"
	"fmt"
)

func TestStrongPassword(t *testing.T) {
	s := []byte("123aA*")

	pat := []string{
		"0123456789",
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"!@#$%^&*()-+",
	}

	fmt.Println(strongPasswd(0, s, pat))
}
func strongPasswd(size int, s []byte, pat []string) int {
	count := 0

	if len(s) >= 6 {
		count++
	}
	for i := range pat {
		if strings.Contains(pat[i], string(s[i])) {
			count++
		} else {
			break
		}

	}
	return count
}
