package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
)

func main() {

	resp, err := http.Get("http://cmm-thesis.herokuapp.com")
    if err != nil {
        log.Fatal(err)
    }

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

	var f interface{}
    error := json.Unmarshal(body, &f)
    if error != nil {
        log.Fatal(error)
    }
	fmt.Println(f)
}
