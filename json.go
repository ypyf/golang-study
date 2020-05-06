package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Header struct {
	ID string `json:"id"`
}

type Foo struct {
	Header
	AppName string `json:"app_name"`
	Code    int    `json:"code"`
	Fuck map[string]interface{}
}

func main() {
	foo := Foo{AppName: "aaaaa"}
	err := json.NewEncoder(os.Stdout).Encode(foo)
	if err != nil {
		fmt.Errorf("encode synced PVC list failed: %v", err)
	}

	var x Foo
	json.NewDecoder(strings.NewReader(`{"id":"123", "app_name":"wx", "fuck":{"a":1, "b":1}}`)).Decode(&x)
	fmt.Printf("%v\n", x.Fuck)
	if v, ok := x.Fuck["a"]; ok {
		x.Code = int(v.(float64))
	}
	fmt.Printf("code = %d\n", x.Code)

	b, err := json.Marshal(&Foo{})
	if err == nil {
		fmt.Printf("%s\n", string(b))
	}
}
