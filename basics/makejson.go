package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//declare a dictionary
	u1 := make(map[string]string)
	var input string
	fmt.Print("Enter name\n")
	fmt.Scan(&input)
	u1["name"] = input

	fmt.Print("Enter address\n")
	fmt.Scan(&input)
	u1["address"] = input

	//convert dictionary to json
	u1_json, _ := json.Marshal(u1)
	fmt.Println(string(u1_json))

}
