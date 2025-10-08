package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://vit.ac.in/"

func main() {
	response,err := http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Println(response)

	defer response.Body.Close() //caller's res to close the connection.

	databytes, err := io.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	} 
	content := string(databytes)
	fmt.Println(content) 
}