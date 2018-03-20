package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	str, _ := reader.ReadString('\n')
	split := strings.Split(strings.TrimSpace(str), " ")
	fmt.Println(split)
	fmt.Println("gcd:", gcd(split))

}

func gcd(num []string) string {
	if num == nil {
		return ""
	}

	size := len(num)
	if size == 1 {
		return num[0]
	}

	if size == 2 {
		return nonRecursion(num[0], num[1])
	}

	t := nonRecursion(num[0], num[1])
	for i := 2; i < size; i++ {
		t = nonRecursion(num[i], t)
	}
	return t

}

//欧几里得算法，辗转相除
func recursion(sm, sn string) string {
	n, _ := strconv.Atoi(sn)
	if n == 0 {
		return sm
	}
	m, _ := strconv.Atoi(sm)
	i := m % n
	itoa := strconv.Itoa(i)
	return recursion(sn, itoa)
}

func nonRecursion(sm, sn string) string {
	m, _ := strconv.Atoi(sm)
	n, _ := strconv.Atoi(sn)

	for n != 0 {
		r := m % n
		m = n
		n = r
	}
	return strconv.Itoa(m)
}
