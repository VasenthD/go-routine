// main function is also  go routine 
// in the below function the main fuction is sending a data and another go function is an receiver 

package main

import "fmt"

func routine(ch chan int){
	fmt.Println(100+ <-ch)
}


func main(){
	channel_demo :=make (chan int)

	go routine(channel_demo)
	
	// receiver is must if there is sender
	
	channel_demo <- 234

	fmt.Println("valueof channel", channel_demo)
	fmt.Printf("type ofchannel %T",channel_demo)

}