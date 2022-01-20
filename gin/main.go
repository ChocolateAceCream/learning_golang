package main

import (
	"fmt"
	"gin/router"
)

func main() {
	r := router.SetupRouter()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
	// r.Run(":8000")
}
