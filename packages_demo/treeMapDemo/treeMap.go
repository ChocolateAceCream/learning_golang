/*
* @fileName treeMap.go
* @author Di Sheng
* @date 2023/12/17 19:22:39
* @description: Demo for treeMap pkg
 */

// treeMap pkg provide a map whose keys are ordered using a black/red tree
package TreeMapDemo

import (
	"fmt"

	"github.com/emirpasic/gods/maps/treemap"
)

func Comparator(a, b interface{}) int {
	return a.(int) - b.(int)
}

func Init() {
	tm := treemap.NewWith(Comparator)
	tm.Put(1, 5)
	tm.Put(2, 25)
	tm.Put(-5, 2)
	tm.Put(-15, 12)
	i := tm.Iterator()
	for i.Next() {
		k := i.Key()
		v := i.Value()
		fmt.Println(k, v)
	}
	v, f := tm.Get(5)
	fmt.Println(v, f)
}
