package main

import (
	"encoding/json"
	"fmt"
)

type message struct {
	Name string `json:"name_field"` // for ser/deser json field to this structs field
	Val int32
}

func main() {
	jsonStr := `{
	"name_field":"foo",
	"val": 123
}`
	var msg message
	json.Unmarshal([]byte(jsonStr), &msg)
	fmt.Println("json read: ", msg)

	otherMsg := message{Name:"asd", Val:123}
	bytes, err := json.Marshal(otherMsg)
	if err != nil {
		fmt.Println("got error during serialization", err)
		return
	}
	fmt.Println("good serialization", string(bytes))
}