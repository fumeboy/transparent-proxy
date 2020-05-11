package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	resp, err := http.Get("http://1.0.0.1:20080/")
	if err != nil{
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
