package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	var numsMap = make(map[int]int)

	for index, num := range nums {
		var complement = target - num

		_, ok := numsMap[complement]
		if ok {
			return []int{numsMap[complement], index}
		}
		numsMap[num] = index
	}
	return nums
}

func main() {
	var nums = []int{2, 2, 7, 15}
	var target = 9

	sum := TwoSum(nums, target)
	fmt.Println(sum)

}
