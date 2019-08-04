package main

import "fmt"

func main() {
	var nums = []int{2, 2, 7, 15}
	var target = 9

	sum := twoSum(nums, target)
	fmt.Println(sum)

}

func twoSum(nums []int, target int) []int {
	var numsMap = make(map[int]int)

	for index, num := range nums {
		var complement = target - num

		_, ok := numsMap[complement]
		if ok {
			return []int{numsMap[complement],index}
		}
		numsMap[num] = index

	}
	return nums
}
