package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	serverTime float64
}

func main() {
	fmt.Println("Starting the application...")
	response, err := http.Get("https://api.binance.com//api/v1/time")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var m Message
		err := json.Unmarshal(data, &m)

		if err != nil {
			return
		}

		fmt.Println(string(data))
	}
	/*
	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	*/
	fmt.Println("Terminating the application...")
}