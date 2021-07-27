package hello

import (
	"rsc.io/quote/v3"
)

// func HelloV1() string {
// 	return quoteV1.Hello()
// }

// func V3() string {
// 	return quoteV3.Concurrency()
// }

//after re-write program using quoteV3
func Hello() string {
	return quote.HelloV3()
}

func V3() string {
	return quote.Concurrency()
}
