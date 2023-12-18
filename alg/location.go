/*
机器⼈坐标问题
问题描述
有⼀个机器⼈，给⼀串指令，L左转 R右转，F前进⼀步，B后退⼀步，问最后机器⼈的
坐标，最开始，机器⼈位于 0 0，⽅向为正Y。 可以输⼊重复指令n ： ⽐如 R2(LF) 这
个等于指令 RLFLF。 问最后机器⼈的坐标是多少？
*/

package Alg

import (
	"fmt"
	"unicode"
)

const (
	Left = iota
	Top
	Right
	Bottom
)

func Location() {
	x, y, z := move("R2(LF)", 0, 0, Top)
	fmt.Printf("(%v, %v), facing %v", x, y, z)
}

func move(cmd string, x0, y0, z0 int) (x, y, z int) {
	x, y, z = x0, y0, z0
	count := 0
	repeatCmd := ""
	for _, s := range cmd {
		switch {
		case unicode.IsNumber(s):
			count = count*10 + (int(s) - '0')
		case s != '(' && s != ')' && count > 0:
			repeatCmd = repeatCmd + string(s)
		case s == ')':
			for i := 0; i < count; i++ {
				x, y, z = move(repeatCmd, x, y, z)
			}
			count = 0
			repeatCmd = ""
		case s == 'L':
			z = (z - 1 + 4) % 4
		case s == 'R':
			z = (z + 1) % 4
		case s == 'F':
			if z == Top || z == Bottom {
				y = y - z + 2
			}
			if z == Right || z == Left {
				x = x + z - 1
			}
		case s == 'B':
			if z == Top || z == Bottom {
				y = y + z - 2
			}
			if z == Right || z == Left {
				x = x - z + 1
			}
		}
	}
	return
}
