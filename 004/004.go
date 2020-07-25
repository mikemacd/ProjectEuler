package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

)

type idx struct{ i, j int }
type productStore map[int]idx

func main() {
	products := productStore{}
	for j := 999; j >= 100; j-- {
		for i := 999; i >= 100; i-- {
			if i < j {
				products[i*j] = idx{i, j}
			}
			if j < i {
				products[i*j] = idx{j, i}
			}
		}
	}

	p := []int{}
	for i := range products {
		p = append(p, i)
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i] > p[j]
	})

	for _, v := range p {
		// if v is palindromic return
		if isPalindrome(v) {
			fmt.Printf("product: %d operands:%v\n", v, products[v])
			os.Exit(0)
		}
	}

}

func isPalindrome(n int) bool {
	v := strconv.Itoa(n)
	l := len(v)

	for i := 0; i <= l/2; i++ {
		if l-i >= 0 && v[i] != v[l-i-1] {
			return false
		}
	}

	return true
}
