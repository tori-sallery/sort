package main

import (
	"fmt"
)

func main() {
	//从小到大
	originArr := []int{3, 7, 4, 0, 2, 31}

	quickSort(0, len(originArr)-1, originArr)
	for _, v := range originArr {
		fmt.Println(v)
	}
	//bubbleSort(originArr)
	//bubbleSort2(originArr)
	//for _, v := range originArr {
	//	fmt.Println(v)
	//}
}

/*
	冒泡排序，和快排是一对
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

func quickSort(start, end int, arr []int) {
	i := start
	j := end
	k := arr[start]
	if i < j {
		for i < j {
			for arr[i] < k {
				i++
			}
			for arr[j] > k {
				j--
			}
			arr[i], arr[j] = arr[j], arr[i]
		}
		//第一次排完 左边都比k小 右边比k大
		for _, v := range arr {
			fmt.Print(v)
			fmt.Print(" ")
		}
		quickSort(start, i-1, arr)
		quickSort(j+1, end, arr)
	}
}
