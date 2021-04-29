package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal(err)

	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", data)
	postBody, _ := json.Marshal(map[string]string{
		"name":  "bindu",
		"email": "bin@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	if err != nil {
		log.Fatalf("error occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}
