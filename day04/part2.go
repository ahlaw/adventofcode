package main

import "fmt"

func check(num int) bool {
	var seqOk bool
	seqLength := 1
	for num > 0 {
		digit := num % 10
		num /= 10
		next := num % 10
		if next > digit {
			return false
		} else if next == digit {
			seqLength++
		} else {
			if seqLength == 2 {
				seqOk = true
			}
			seqLength = 1
		}
	}
	return seqOk
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
