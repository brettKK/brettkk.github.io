package main

import "fmt"

// func partition(arr []int) int{
// 	if len(arr) <= 1 {
// 		return 0
// 	}
// 	var pivot int = arr[0]
// 	var latestSmall = 0
// 	for i := 1; i < len(arr); i++ {
// 		if arr[i] < pivot {
// 			latestSmall++
// 			arr[latestSmall], arr[i] = arr[i], arr[latestSmall]
// 		}
// 	}
// 	return latestSmall
// }

func quick_sort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	var start, index int = 0, 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[start] {
			index++
			tmp := arr[index]
			arr[index] = arr[i]
			arr[i] = tmp		
		}
	}
	tmp := arr[index]
	arr[index] = arr[start]
	arr[start] = tmp
	quick_sort(arr[start : index])
	quick_sort(arr[index + 1 :])
}

func quick_sort_simple(arr []int) {
    if len(arr) < 2 {
        return
    }
    var start, index int = 0, 0
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[start] {
            index++
            arr[i], arr[index] = arr[index], arr[i]
        }
    }
    arr[start], arr[index] = arr[index], arr[start]
    quick_sort(arr[start:index])
    quick_sort(arr[index+1:])
}

func main() {
	arr := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	fmt.Println(arr)
	quick_sort_simple(arr)
	fmt.Println(arr)
}
