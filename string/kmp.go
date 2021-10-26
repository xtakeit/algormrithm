package main

import (
    "fmt"
)

func main() {
    arr := getNext("abab")
    for _, v := range arr {
        fmt.Println(v)
    }
}

// kmp 算法实现
func Kmp(s, p string) {

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
            if p[last] != p[start] {
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
