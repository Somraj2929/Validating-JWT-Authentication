package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"time"
)

func main() {
    // define the URL and method 
	//on this url check token for signing in 
	url := "http://localhost:8080/login"
	method := "POST"


	//these credentials will be matched with server's map 
	payload := strings.NewReader(`{
    "username": "somraj",
    "password": "password"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))



		// get the token
	token := string(body)
	
	//time.Sleep(time.Minute * 2)
	//books data will be displayed if token found in header
	url2 := "http://localhost:8080/getAllBooks"
	method2 := "GET"

	req2, err := http.NewRequest(method2, url2, nil)

	if err != nil {
		fmt.Println(err)
		return

	}
	//adding token to header
	req2.Header.Add("Token", token)

	res2, err := client.Do(req2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res2.Body.Close()

	body2, err := ioutil.ReadAll(res2.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body2))


	
}