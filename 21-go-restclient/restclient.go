package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SimpleHttpGet() {
	httpclient := http.DefaultClient
	url := "https://fakerapi.it/api/v1/persons"
	// url = "http://somedomainthatdoesnotexisttrevor.com"
	resp, err := httpclient.Get(url)
	fmt.Println(err)
	fmt.Println(resp.Body)

	data, err := io.ReadAll(resp.Body)
	fmt.Println(string(data))

	var fakedata FakerPersonResponse = FakerPersonResponse{}
	unmarshalerr := json.Unmarshal(data, &fakedata)
	fmt.Println(unmarshalerr)
	fmt.Printf("%+v\n", fakedata)

	for _, v := range fakedata.Data {
		fmt.Printf("My first name is %v and my gender is %v\n", v.FirstName, v.Gender)
	}
}

func MakeCustomHTTPRequest() {
	method := "get"
	url := "https://fakerapi.it/api/v1/persons"
	bodyString := ""
	bodyReader := strings.NewReader(bodyString)
	req, _ := http.NewRequest(method, url, bodyReader)
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(err)
	respdata, _ := io.ReadAll(resp.Body)

	fmt.Println(string(respdata))
}

func main() {
	// SimpleHttpGet()
	MakeCustomHTTPRequest()
}

type FakerPersonResponse struct {
	Status string        `json:"status"`
	Total  uint16        `json:"total"`
	Code   uint16        `json:"code"`
	Data   []FakerPerson `json:"data"`
}

type FakerPerson struct {
	Id        uint16 `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
}
