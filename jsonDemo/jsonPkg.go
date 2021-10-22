package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/*
		To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the keys used by Marshal (either the struct field name or its tag), preferring an exact match but also accepting a case-insensitive match. By default, object keys which don't have a corresponding struct field are ignored (see Decoder.DisallowUnknownFields for an alternative).
	*/
	fmt.Println("--------unmarshal into struct demo-------")
	type Msg struct {
		Name  string
		Order interface{}
	}

	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	var msg []Msg
	err := json.Unmarshal(jsonBlob, &msg)
	if err != nil {
		fmt.Println(err)
	}

	// when printing structs, the plus flag (%+v) adds field names
	fmt.Printf("%+v", msg)
}
