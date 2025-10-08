package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "this string will be saved inside the txt file"

	file,err := os.Create("./abc.txt")
	if err!=nil{
		panic(err)
	}

	length, err:= io.WriteString(file, content)
	if err!=nil{
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	readddddfile("./abc.txt")

}

func readddddfile(file string){
	databyte,err := ioutil.ReadFile(file)
	if err!=nil{
		panic(err)
	}
	content:=string(databyte)
	
	fmt.Println(content)
}