package main

import "fmt"

func main() {
	fmt.Println("%d", KMP([]byte{'1','2', '3'}, []byte{'2'}))
}

/*
* return index when pattern first occur in string
*/
func KMP(str []byte, pattern []byte) int {
	next := get_next_array(pattern)
	i, j := 0, 0
	for i < len(str) {
		if str[i] == pattern[j] {
			i++;
			j++;
		} else if j > 0 { //跳过pattern中的一些子符
			j = next[j-1]
		} else { // str的第一个字符串就不匹配pattern
			i++
		}
		if j == len(pattern) {
			return i - j
		}
	}
	return  -1
}

func get_next_array(pattern []byte) []int {
	// 当前字符串中前缀 和 后缀 相同的长度
	var prefix_len, i int = 0, 1
	var next []int = make([]int, len(pattern))
	next[0] = 0
	for i < len(pattern) {
		if pattern[i] == pattern[prefix_len] {
			prefix_len++
			next[i] = prefix_len
			i++
		} else {
			// p[prefix_len] != p[i]
			// 回退prefix_len到某个位置
			// a b c d e f a b c g
			//                   i prefix_len = 3
			//    prefix_len        pattern[i] != pattern[3] , g != d
			//    下一次prefix_len将缩小到next[prefix_len]的位子
			prefix_len = next[prefix_len - 1]
			if prefix_len == 0 {
				next[i] = 0
				i++
			}
		}
	}
	return next
}