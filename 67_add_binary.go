package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""
	for i >= 0 || j >= 0 || carry != 0 {
		if i >= 0 {
			//默认ASCII编码
			carry += int(a[i]) - '0'
		}
		if j >= 0 {
			carry += int(b[j]) - '0'
		}
		result = strconv.Itoa(carry%2) + result
		carry /= 2
		i--
		j--
	}

	return result
}

/**
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Print(addBinary("1010", "1011"))
}
