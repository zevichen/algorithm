package strings

import (
	"testing"
	"fmt"
)

func TestSuperReducedString(t *testing.T) {

	arrList := [][]byte{
		[]byte("aaabbbcdddeff"),
		[]byte("abba"),
		[]byte("aaaaa"),
		[]byte("bbaaabbbcdaaabbb"),
		[]byte("abbccbb"),
	}

	for n := range arrList {
		bytes := reduce(arrList[n])
		fmt.Println("arr:", string(arrList[n]), "str:", string(bytes))
	}
}

func reduce(arr []byte) []byte {
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			if i == 1 {
				arr = arr[i+1:]
			} else {
				t := make([]byte, 0)
				arr = append(append(t, arr[:i-1]...), arr[i+1:]...)
			}
			i = 0
		}
	}
	if len(arr) == 0 {
		return []byte("Empty String")
	} else {
		return arr
	}
}
