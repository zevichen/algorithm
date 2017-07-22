package sorting

import (
	"testing"
	"fmt"
)

//TestInsertion 从a[1]开始取值,把大于a[1]的值向后移动一位，然后把该值插入那个空位
func TestInsertion(t *testing.T) {

	arr := []int{20, 14, 10, 4, 42, 49, 38, 40, 24, 2}
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			j := i - 1
			t := arr[i]
			for j >= 0 && t < arr[j] {
				arr[j+1] = arr[j]
				j--
			}
			arr[j+1] = t
		}
	}
	fmt.Println(arr)

}
