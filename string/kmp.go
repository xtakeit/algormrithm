package main

import (
	"fmt"
	"strconv"
)

// https://www.cnblogs.com/zzuuoo666/p/9028287.html
func main() {
	s := "BBC ABCDAB ABCDABCDABDE"
	p := "ABCDABD"
	pos := Kmp(s, p)
	fmt.Println("pos:" + strconv.Itoa(pos))
}

// kmp 算法实现
func Kmp(s, p string) int {
	i, j := 0, 0
	next := getNext(p)
	for i < len(s) && j < len(p) {
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(p) {
		return i - j
	} else {
		return -1
	}
}

// 获取子串的next数组
func getNext(p string) []int {
	pLen := len(p)
	start := -1
	last := 0

	next := make([]int, pLen)
	next[0] = -1

	for last < pLen-1 {

		if start == -1 || p[start] == p[last] {
			start++
			last++
			next[last] = start
			if p[last] != p[start] { //TODO 待理解
				next[last] = start
			} else {
				next[last] = next[start]
			}

		} else {
			start = next[start]
		}
	}

	return next
}
