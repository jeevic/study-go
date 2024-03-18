package main

import (
	"encoding/json"
	"log"
)

var JSON = `{
 "name": "Gopher",
 "title": "programmer",
 "contact": [{
 "home": "415.333.3333",
 "cell": "415.555.5555"
 },
 {
   "home": "415.333.3333",
   "cell": "415.555.5555"
 }
 ]
 }`

func main() {

	var c map[string]interface{}

	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	var home = c["contact"].([]interface{})[0].(map[string]interface{})["home"].(string)
	_ = home

}
