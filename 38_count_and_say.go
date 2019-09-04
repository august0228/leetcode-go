package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {

	if n < 0 || n > 30 {
		return ""
	}
	if n == 1 {
		return "1"
	}

	r := "1"
	resultStr := ""

	for j := 0; j < n-1; j++ {
		resultStr = func(s string) string {
			result := ""
			sameStr := ""
			next := ""
			for i := 0; i < len(s); i++ {

				current := string(s[i])
				if i+1 < len(s) {
					next = string(s[i+1])
				} else {
					next = ""
				}

				if sameStr == "" {
					sameStr = current
				}

				if next != "" {
					//如果当前和下一个相等 暂存
					if current == next {
						sameStr = sameStr + current
					} else {
						//如果不相等且当前没有暂存的，直接拼接1+当前数字
						if sameStr == "" {
							result = result + "1" + current
						} else {
							//有暂存的 拼接暂存的
							result += strconv.Itoa(len(sameStr)) + current
							sameStr = ""
						}
					}
				} else {
					result += strconv.Itoa(len(sameStr)) + current
				}
			}
			return result
		}(r)

		r = resultStr
	}

	return resultStr
}

/**
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

示例 1:

输入: 1
输出: "1"
示例 2:

输入: 4
输出: "1211"
在真实的面试中遇到过这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/count-and-say
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Printf(countAndSay(22))
}
