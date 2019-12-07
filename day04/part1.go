package main

import "fmt"

func check(num int) bool {
	var seq bool
	for num > 0 {
		digit := num % 10
		num /= 10
		next := num % 10
		if next > digit {
			return false
		} else if next == digit {
			seq = true
		}
	}
	return seq
}

func main() {
	lower := 264360
	upper := 746325
	var count int
	for i := lower; i <= upper; i++ {
		ok := check(i)
		if ok {
			count++
		}
	}
	fmt.Println(count)
}
