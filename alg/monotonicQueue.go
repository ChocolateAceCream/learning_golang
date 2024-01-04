/*leetcode #239
for a monotonic increasing queue, q[0] is always the minimum element in the queue
e.g.
for input [3,-1,2,6,7], after each insertion, the monotonic decreasing queue will be:
[3]
[3, -1]
[3, 2]
[6]
[7]
and the monotonic increasing queue will be
[3]
[-1]
[-1, 2]
[-1, 2, 6]
[-1, 2, 6, 7]
*/
package Alg

import (
	"fmt"
	"math/rand"
)

func IncreasingMonotonicDemon() {
	q := []int{}
	for i := 0; i < 20; i++ {
		fmt.Println("before: ", q)
		insert := rand.Intn(100)
		fmt.Println("insert: ", insert)
		if len(q) == 0 {
			q = append(q, insert)
		} else {
			for idx, n := range q {
				if insert > n {
					q = append(q[:idx], insert)
					break
				} else if idx == len(q)-1 {
					q = append(q, insert)
				}
			}
		}
		fmt.Println("after: ", q)

	}

}
