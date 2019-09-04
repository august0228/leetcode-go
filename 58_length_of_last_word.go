package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
逆序遍历 找到第一个空格为止
*/
func lengthOfLastWord(s string) int {
	str := strings.TrimSpace(s)
	if str == "" {
		return 0
	}

	strLength := 0
	for i := len(); i > 0; {
		r, size := utf8.DecodeLastRuneInString(str[:i])
		_r := string(r)

		if _r == " " {
			return strLength
		} else {
			strLength = strLength + 1
		}
		i -= size
	}
	return strLength
}

/**
给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指由字母组成，但不包含任何空格的字符串。
示例:
输入: "Hello World"
输出: 5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/length-of-last-word
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main() {
	fmt.Print(lengthOfLastWord("a"))
}
