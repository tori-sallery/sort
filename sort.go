package main

import "fmt"

func main() {
	//从小到大
	originArr := []int{3, 7, 4, 0, 9, 31}
	bubbleSort(originArr)
	bubbleSort2(originArr)
	for _, v := range originArr {
		fmt.Println(v)
	}
}

/*
	冒泡排序
*/
func bubbleSort(arr []int) {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := i + 1; j < arrLen; j++ {
			//谁比第一个小放到最前边，这样前i个就是有序的
			if arr[j] < arr[i] {
				tmp := arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}
		for _, v := range arr {
			fmt.Print(v)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

//优化一下  整体有序不需要跑循环了
func bubbleSort2(arr []int) {
	arrLen := len(arr)
	exchange := true
	for exchange {
		exchange = false
		for j := 0; j < arrLen-1; j++ {
			//大的放到最后边，两两比较，如果一次都没发生交换，说明整体有序
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				exchange = true
			}
		}
		for _, v := range arr {
			fmt.Print(v)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
