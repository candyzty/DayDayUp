package main

import "log"

func main() {
	var nums = []int{90, 123, 412, 412, 4, 12, 3245, 3245, 63245, 32463246, 23453245, 123123, 1231, 1, 0, -23, 34234, 6789, 769, 56578, 569590, 1}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	log.Println(nums)
}
