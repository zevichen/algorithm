package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"strconv"
)

//https://www.hackerrank.com/challenges/password-cracker/problem
func main() {

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

	for j := 0; j < count; j++ {
		cut := strings.Split(arrStr[j], " ")
		rst := make([]string, 0)

		if notContainLetter(arrMatch[j], arrStr[j]) {
			fmt.Println("WRONG PASSWORD")
			continue
		}

		if !recursion(cut, rst, arrMatch[j], 0) {
			fmt.Println("WRONG PASSWORD")
		}
	}
}

func recursion(origin, rst []string, match string, depth int) bool {

	depth += 1
	length := len(match)
	if length == 0 {
		fmt.Println(strings.Join(rst, " "))
		return true
	}
	if depth > 2000 {
		return false
	}
	for i := 0; i < len(origin); i++ {
		orSize := len(origin[i])
		if len(match) >= orSize {
			if match[:orSize] == origin[i] {
				rst = append(rst, origin[i])
				return recursion(origin, rst, match[orSize:], depth)
			}
		}
	}
	return false
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
