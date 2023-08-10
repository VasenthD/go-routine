package main

import "fmt"

func myfunc(channel chan string){
	for i:=0;i<1191;i++{
		channel <- "go lang is awesome"
	}
	close(channel)

}
func main(){

	c:=make(chan string,8)
	go myfunc(c)

	counter:=0
	for{

		res, ok:= <-c
		counter++
		if !ok{
			fmt.Println("chanel is clossed ",ok)
			break
		}
		fmt.Println(counter)
		fmt.Println("channel is open and dat is  ",res,ok)
	}
}