package main

import "fmt"

func main(){
	dic:= newDIContainer(10)
	if err := runHTTPServer(dic,"8000"); err != nil {
		fmt.Println(err)
	}
}
