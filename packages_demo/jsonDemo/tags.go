package main

import (
	"encoding/json"
	"fmt"
)

type Demo struct {
	Id    int `json:"uuid"` //use alias name uuid
	Name  string
	Count int `json:"-"`            //ignore this field
	Fu    int `json:"aa,omitempty"` //ignore this filed if value is empty or default value, such as false、0、nil、 or any array，map，slice，string with length === 0
}

func main() {

	dd := Demo{
		Id:    3,
		Name:  "@data&%",
		Count: 4,
		Fu:    0,
	}
	data, err := json.Marshal(dd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", string(data))

	d := Demo{Id: 1, Name: "d"}
	fmt.Printf("%+v\n", d)
	jsonData := []byte(`{"uuid": 221, "Name": "dddd", "Count":123, "Fu":12}`)
	var parsedJson Demo
	err = json.Unmarshal(jsonData, &parsedJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", parsedJson)
}
