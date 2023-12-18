/*
翻转字符串
问题描述
请实现⼀个算法，在不使⽤【额外数据结构和储存空间】的情况下，翻转⼀个给定的字
符串(可以使⽤单个过程变量)。
给定⼀个string，请返回⼀个string，为翻转后的字符串。保证字符串的⻓度⼩于等于
5000。
*/
package Alg

// func main() {
// 	if str, r := reverse("asdf123"); r {
// 		fmt.Println(str)
// 	}
// }

func reverse(str string) (string, bool) {
	length := len(str)
	if length >= 5000 {
		return str, false
	}
	s := []rune(str)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
	return string(s), true
}
