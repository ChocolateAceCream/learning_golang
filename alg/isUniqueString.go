/*
判断字符串中字符是否全都不同
问题描述
请实现⼀个算法，确定⼀个字符串的所有字符【是否全都不同】。这⾥我们要求【不允
许使⽤额外的存储结构】。 给定⼀个string，请返回⼀个bool值,true代表所有字符全都
不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的⻓
度⼩于等于【3000】。
*/

package Alg

import (
	"strings"
)

func isUniqueString(str string) bool {
	if strings.Count(str, "") > 3000 {
		return false
	}
	for _, v := range str {
		//ASCII only contains 128 characters which can be typied in through keyboard.
		// v > 127 will filter out any characters that out of range of ASCII character set
		if v > 127 {
			return false
		}
		if strings.Count(str, string(v)) > 1 {
			return false
		}
	}
	return true
}
