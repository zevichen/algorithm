package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {

	var count int

	fmt.Scanln(&count)

	arrSize := make([]int, count)
	arrStr := make([]string, count)
	arrMatch := make([]string, count)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < count; i++ {
		s, _ := reader.ReadString('\n')
		num, _ := strconv.Atoi(strings.TrimSpace(s))
		arrSize[i] = num

		s, _ = reader.ReadString('\n')
		arrStr[i] = strings.TrimSpace(s)

		s, _ = reader.ReadString('\n')
		arrMatch[i] = strings.TrimSpace(s)
	}

	//var count = 3
	//arrStr := make([]string, count)
	//arrMatch := make([]string, count)
	//arrStr = []string{
	//	"ozkxyhkcst xvglh hpdnb zfzahm",
	//	"gurwgrb maqz holpkhqx aowypvopu",
	//	"a aa aaa aaaa aaaaa aaaaaa aaaaaaa aaaaaaaa aaaaaaaaa aaaaaaaaaa",
	//}
	//arrMatch = []string{
	//	"zfzahm",
	//	"gurwgrb",
	//	"aaaaaaaaaab",
	//}

	for j := 0; j < count; j++ {
		cut := strings.Split(arrStr[j], " ")
		rst := make([]string, 0)

		if notContainLetter(arrMatch[j], arrStr[j]) {
			fmt.Println("WRONG PASSWORD")
			continue
		}

		recursion, b := recursion(cut, rst, arrMatch[j])
		if b {
			fmt.Println(strings.Join(recursion, " "))
		} else {
			fmt.Println("WRONG PASSWORD")
		}
	}
}

func recursion(origin, rst []string, match string) ([]string, bool) {
	if len(match) == 0 {
		return rst, true
	}

	for i := 0; i < len(origin); i++ {

		if strings.HasPrefix(match, origin[i]) {
			match = strings.Replace(match, origin[i], "", 1)
			rst = append(rst, origin[i])
			temp, flag := recursion(origin, rst, match)
			if flag {
				return temp, flag
			}
		}
	}
	return nil, false

}

func notContainLetter(left, right string) bool {
	replace := strings.Replace(left, " ", "", -1)
	letter := strings.Split(replace, "")
	for i := range letter {
		if !strings.Contains(right, letter[i]) {
			return true
		}
	}
	return false

}
