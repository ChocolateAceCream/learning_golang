package main

import (
	"fmt"
	"reflect"
	"time"
)

type Note struct {
	Time    time.Time `must:"true"` //Caution: no space in between must: and "true"
	Content string    `must:"true"`
}

func main() {
	note := Note{
		Time:    time.Now(),
		Content: "gg",
	}
	ValidateMustField(&note)
	note2 := Note{}
	ValidateMustField(&note2)
}

func ValidateMustField(value interface{}) {
	val := reflect.ValueOf(value).Elem()
	typ := reflect.TypeOf(value).Elem()
	for i := 0; i < typ.NumField(); i++ {
		if typ.Field(i).Tag.Get("must") == "true" && val.Field(i).IsZero() {
			fmt.Println("Fail:", typ.Field(i).Name, "is null")
			return
		}
	}
	fmt.Println("pass")

}
