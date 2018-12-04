package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	if flag.Arg(0) == "help" {
		fmt.Println("Usage: address <code>")
		os.Exit(0)
	}
	address_code := flag.Arg(0)

	resp, err := http.Get("http://zipcloud.ibsnet.co.jp/api/search?zipcode=" + address_code)
	if err != nil {
		log.Fatal(err)
	}
	status := resp.StatusCode
	if status != http.StatusOK {
		log.Fatal("it's not 200. Return status is " + string(status))
	}

	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := new(Addr)

	if err := json.Unmarshal(res, data); err != nil {
		log.Fatal("JSON Unmarshal error:", err)
	}
	if len(data.Results) == 0 {
		log.Fatal(address_code + " is not exist")
	}
	results := data.Results[0]
	pref := results.Address1
	city := results.Address2
	number := results.Address3

	fmt.Println(pref + city + number)
}

type Addr struct {
	Message interface{} `json:"message"`
	Results []Result    `json:"results"`
	Status  int         `json:"status"`
}

type Result struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	Kana1    string `json:"kana1"`
	Kana2    string `json:"kana2"`
	Kana3    string `json:"kana3"`
	Prefcode string `json:"prefcode"`
	Zipcode  string `json:"zipcode"`
}
