/*
* @fileName regexDemo.go
* @author Di Sheng
* @date 2022/07/04 10:46:35
* @description
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
)

type Payload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r, _ := regexp.Compile(`^[+-]?([0-9]+\.?[0-9]*((e|E)[+-]?[0-9]+)?|\.[0-9]+((e|E)[+-]?[0-9]+)?|)$`)
	fmt.Println(r.MatchString("e"))

	r2, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r2.MatchString("peach"))
	url := "http://localhost:4050/api/public/auth/login"
	payload := &Payload{
		Username: "admin",
		Password: "fbfb386efea67e816f2dda0a8c94a98eb203757aebb3f55f183755a192d44467",
	}
	p, _ := json.Marshal(payload)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(p)))
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	fmt.Println(string(resp.Status))

}
