package sorting

import (
	"fmt"
	"testing"
)

var arr [10]int = [10]int{20, 14, 10, 4, 42, 49, 38, 40, 24, 2}
var arr1 [10]int

//ExampleQuick 使用arr[l]基准值，比较right和left和基准的大小,然后互换位置，l,i-1;i+1,r分别递归
func TestQuick(t *testing.T) {

	arr1 = arr

	quick(0, len(arr)-1)
	fmt.Println("old:", arr1)
	fmt.Println("new:", arr)

}
func quick(l, r int) {

	if l > r {
		return
	}

	i, j := l, r

	for i != j {

		for arr[j] > arr[l] && i < j {
			j--
		}

		for arr[i] < arr[l] && i < j {
			i++
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			fmt.Println("i=", i, ",j=", j, " arr：", arr)
		}
	}

	arr[l], arr[i] = arr[i], arr[l]

	quick(l, i-1)
	quick(i+1, r)

}
