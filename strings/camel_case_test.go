package strings

import (
	"testing"
	"fmt"
)

func TestCamelCase(t *testing.T) {
	a := []byte("saveChangesInTheEditor")
	fmt.Println(countWord(a))
}
func countWord(s []byte) int {
	count := 0
	for i := 0; i < len(s); i++ {
		if 'A' <= s[i] && s[i] <= 'Z' {
			count++
		}
	}
	return count + 1
}
