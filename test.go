package main

import (
	"./jserver"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ParseResponse(res *http.Response) (string, int) {
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	return string(contents), res.StatusCode
}

func assert(data map[string]interface{}) bool {
	b, _ := json.Marshal(data)
	_, err := http.Post("http://localhost:4000/post", "application/json", bytes.NewReader(b))
	if err != nil {
		fmt.Println("Error: POST FAIL")
		return false
	}

	res, _ := http.Get("http://localhost:4000/get")
	c, _ := ParseResponse(res)
	fmt.Println("post:", string(b))
	fmt.Println(" get:", c)
	if c == string(b) {
		return true
	}
	return false
}

func main() {
	go jserver.Start()

	data1 := map[string]interface{}{
		
		 "created_at":  "2016-05-09T19:45:32Z",
		 "id":  map[int:0 valid:false],
		 "message":  "Hello, World.",
		 
		 "int":  0,
		 "valid":  false,
		 
	}
	fmt.Println("assert 1")
	result1 := assert(data1)


	fmt.Println("result:", result1)
}
