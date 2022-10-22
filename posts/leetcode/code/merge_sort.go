package main

import "fmt"

func merge_sort(arr []int) []int{
	if len(arr) < 2 {return arr}
	mid := len(arr)>>1
	left_arr := merge_sort(arr[:mid])
	right_arr := merge_sort(arr[mid:])
	return merge(left_arr, right_arr)
}

func merge(left, right []int) []int {
	var res []int = make([]int, len(left) + len(right))
	var k, i, j int
	for i <len(left) && j < len(right) {
		if left[i] <= right[j] {
			res[k] = left[i]
			i++ 
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		res[k] = left[i]
		k++; i++
	}
	for j < len(right) {
		res[k] =right[j]
		k++; j++
	}
	return res
}

func main() {
	arr := []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	fmt.Println(arr)
	sorted_arr := merge_sort(arr)
	fmt.Println(sorted_arr)
}