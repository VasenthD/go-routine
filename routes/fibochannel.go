package main

import "fmt"
import "time"


func fibno(n int,c chan int){
	x,y:=0,1
	for i:=0;i<n;i++{
		c<-x
		x,y=y,x+y
	}
	close(c)
}

func main(){
	start:=time.Now()
	c:=make(chan int,100)
	go fibno(cap(c),c)
	for i:=range c{
		fmt.Println(i)
	}
	end:=time.Now().Sub(start)
	fmt.Println(end)
}