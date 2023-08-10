package main

import "fmt"
import "time"

// func p(x int){
// 	for i:=1;i<x;i++{
// 		fmt.Println(i)
// 	}
// }

func main(){

	for i:=0;i<22;i++{
	
		go func(x int){
			fmt.Println(x)
		}(i)
		//time.Sleep(1*time.Second)
		//go p(9)


	}
	time.Sleep(1*time.Second)
}