package main

import "fmt"

// Majority returns the value occurring more than len(nums)/2 times, found by
// Boyer–Moore voting and confirmed by a verification pass. ok is false when
// no such value exists.
func Majority(nums []int) (value int, ok bool) {
	cand, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			cand = v
		}
		if v == cand {
			count++
		} else {
			count--
		}
	}
	occurrences := 0
	for _, v := range nums {
		if v == cand {
			occurrences++
		}
	}
	if len(nums) > 0 && occurrences > len(nums)/2 {
		return cand, true
	}
	return 0, false
}

func main() {
	v, ok := Majority([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(v, ok)
}
