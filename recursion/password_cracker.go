package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {

	var count int
	//var size int
	//var str string
	//var match string

	fmt.Scanln(&count)
	arrSize := make([]int, count)
	arrStr := make([]string, count)
	arrMatch := make([]string, count)
	fmt.Println(count)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < count; i++ {
		s, _ := reader.ReadString('\n')
		atoi, _ := strconv.Atoi(s)
		arrSize[i] = atoi
		s, _ = reader.ReadString('\n')
		arrStr[i] = s
		s, _ = reader.ReadString('\n')
		arrMatch[i] = s
	}
	fmt.Println(arrSize, arrStr, arrMatch)

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

	//for j := 0; j < count; j++ {
	//	cut := strings.Split(arrStr[j], " ")
	//	rst := make([]string, 0)
	//
	//	s, _ := recursion(cut, rst, arrMatch[j])
	//	fmt.Println(s)
	//}
}

func recursion(origin, rst []string, match string) (string, bool) {
	if len(match) == 0 {
		return strings.Join(rst, " "), true
	}

	for i := 0; i < len(origin); i++ {

		if strings.HasPrefix(match, origin[i]) {
			match = strings.Replace(match, origin[i], "", 1)
			rst = append(rst, origin[i])
		} else {
			continue
		}
		s, b := recursion(origin, rst, match)
		if b {
			return s, b
		}

	}
	return "WRONG PASSWORD", false
}