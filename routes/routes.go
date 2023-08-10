package main

import "fmt"
import "time"

func hello(){
	fmt.Println("hello world")
}

func main(){

	fmt.Println("one ")
	go hello()
	fmt.Println("final")
	time.Sleep(10*time.Second)
	fmt.Println("completed")
}