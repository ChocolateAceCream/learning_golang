/*
* @fileName mapstructure.go
* @author Di Sheng
* @date 2022/07/22 11:09:45
* @description demo of github.com/mitchellh/mapstructure pkg
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Person struct {
	Name string
	Age  int
	Job  string
}

type Cat struct {
	Name  string
	Age   int
	Breed string
}

func main() {
	BasicUsage()
	Demo2()
	Demo3()
	Demo4()
	MetadataDemo()
	WeakDecodeDemo()
	CustomizedDecoder()
}

func BasicUsage() {
	dataset := []string{`
    { 
      "type": "person",
      "name":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}
	for _, data := range dataset {
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(data), &m); err != nil {
			fmt.Println(err)
		}
		switch m["type"].(string) {
		case "person":
			var p Person
			mapstructure.Decode(m, &p)
			fmt.Println("person: ", p)
		case "cat":
			var c Cat
			mapstructure.Decode(m, &c)
			fmt.Println("Cat: ", c)
		}
	}
}

// usage of alias, field name is case insensitive.
type Person2 struct {
	//mapstructure will search for name field (case insensitive) in map[string]interface{}
	Name string `mapstructure:"username"`
	Age  int
	Job  string
}

func Demo2() {
	dataset := []string{`
    { 
      "type": "person",
      "username":"dj",
      "age":18,
      "job": "programmer"
    }
  `,
		`
    {
      "type": "cat",
      "name": "kitty",
      "age": 1,
      "breed": "Ragdoll"
    }
  `,
	}
	for _, data := range dataset {
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(data), &m); err != nil {
			fmt.Println(err)
		}
		switch m["type"].(string) {
		case "person":
			var p Person2
			mapstructure.Decode(m, &p)
			fmt.Println("person2: ", p)
		case "cat":
			var c Cat
			mapstructure.Decode(m, &c)
			fmt.Println("Cat: ", c)
		}
	}
}

type P3 struct {
	Name string
}

type Friend1 struct {
	P3
}

type Friend2 struct {
	Name string
	// this way in JSON string there's no need to specify p3 field.
	// if P3 has name field and friend2 also has name field, then mapstructure will decode json name field into both of them
	P3 `mapstructure:",squash"`
}

func Demo3() {
	fmt.Println("---------Demo3-----------------")
	dataset := []string{`
    { 
      "type": "friend1",
      "p3": {
        "name":"dj"
      }
    }
  `,
		`
    {
      "type": "friend2",
      "name": "dj2"
    }
  `,
	}
	for _, val := range dataset {
		var m map[string]interface{}
		if err := json.Unmarshal([]byte(val), &m); err != nil {
			fmt.Println("Unmarshal error : ", err)
		} else {
			switch m["type"].(string) {
			case "friend1":
				var f1 Friend1
				if err := mapstructure.Decode(m, &f1); err != nil {
					fmt.Println("mapstructure err: ", err)
				} else {
					fmt.Println("friend1: ", f1)
				}
			case "friend2":
				var f2 Friend2
				if err := mapstructure.Decode(m, &f2); err != nil {
					fmt.Println("mapstructure err: ", err)
				} else {
					fmt.Println("friend2: ", f2)
					fmt.Println("friend2: ", f2.P3)
				}
			}
		}

	}
}

// demo of remain field: JSON string has more fields than defined in the struct
type P4 struct {
	Name string
	Age  int
	Job  string

	//the data type of remain field must be of map[string]interface{}
	Other map[string]interface{} `mapstructure:",remain"`
}

func Demo4() {
	fmt.Println("---------Demo4-----------------")
	dataset := `
    { 
      "name": "dj",
      "age":18,
      "job":"programmer",
      "height":"1.8m",
      "handsome": true
    }
  `
	var m map[string]interface{}
	json.Unmarshal([]byte(dataset), &m)
	var p P4
	mapstructure.Decode(m, &p)
	fmt.Println("p4: ", p)
}

//reverse decode: decode map[string]interface{} into struct
type P5 struct {
	Name string
	Age  int
	Job  string `mapstructure: ",omitempty"`
}

func Demo5() {
	p1 := &P5{
		Name: "dj",
		Age:  18,
	}
	var m map[string]interface{}
	mapstructure.Decode(p1, &m)
	data, _ := json.Marshal(m)
	fmt.Println("data: ", data)
}

// metadata demo:
type P6 struct {
	Name  string
	Age   int
	Color string
}

func MetadataDemo() {
	m := map[string]interface{}{
		"name": "dj",
		"age":  18,
		"job":  "programmer",
	}
	p := &P6{
		Color: "gree123n",
	}
	var meta mapstructure.Metadata
	/*
		type Metadata struct{
			Keys []string	//success decoded keys
			Unused []string	// other unused keys: exist in source but not specified in target struct
			Unset []string // keys which exist on target struct bue missing on source
		}
	*/
	if err := mapstructure.DecodeMetadata(m, &p, &meta); err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println("meta: ", meta)
	fmt.Println("meta.unused: ", meta.Unused)
	fmt.Println("meta.Keys: ", meta.Keys)
	fmt.Println("meta.Unset: ", meta.Unset)
	fmt.Println("decoded data: ", m)
}

// WeakDecode demo
type P7 struct {
	Name  string
	Age   int
	Color string
}

func WeakDecodeDemo() {
	m := map[string]interface{}{
		"name": "dsdj",
		"age":  "18", // if you changed age to "unknown", then WeakDecode() will return error
		"job":  "programmer",
	}
	p := &P7{
		Color: "gree123n",
	}

	// WeakDecode() will try to convert the value to target type if type is not match, e.g. age field, not not success, then return error
	if err := mapstructure.WeakDecode(m, &p); err != nil {
		fmt.Println("error: ", err.Error())
	}
	fmt.Println("decode data: ", m)

}

// customize decoder demo
type P8 struct {
	Name string
	Age  int
}

func CustomizedDecoder() {
	fmt.Println("--------CustomizedDecoder demo------")
	m := map[string]interface{}{
		"name": 123,
		"age":  "18",
		"job":  "programmer",
	}
	var p P8
	var meta mapstructure.Metadata
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &p,
		Metadata:         &meta,
	})
	if err != nil {
		fmt.Println("error: ", err)
	}
	if err := decoder.Decode(m); err != nil {
		fmt.Println(" Decode error: ", err)
	} else {
		fmt.Println("decode result: ", p)
	}

}

// mapstructure.go
type DecoderConfig struct {
	ErrorUnused      bool // return error if any key in source are unused
	ZeroFields       bool // applied when decode struct to map, if true then empty entire target map, if false, then merge into the target map
	WeaklyTypedInput bool // used to implement WeakDecode/WeakDecodeMetadata
	// Metadata         *Metadata // if not nil then collect meta data
	Result  interface{}
	TagName string // customize tag name, default is mapstructure
}
