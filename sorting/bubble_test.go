package sorting

import (
	"fmt"
	"testing"
)

//ExampleBubble 循环相邻值比较，大的置后，i-1,再次循环（确保每次大循环大值都在后面依次排序）
func TestBubble(t *testing.T) {

	arr := []int{20, 14, 10, 4, 42, 49, 38, 40, 24, 2}

	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Print(arr)

}
