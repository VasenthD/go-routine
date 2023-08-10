package main

import "fmt"
import "time"

func fibo(n int ){
	a:=0
	b:=1
	fmt.Println(0)
	fmt.Println(1)
	for i:=3;i<n;i++{
		c:=a+b
		fmt.Println(c)
		a=b
		b=c
	}

}

func main(){
	start:=time.Now()
	fibo(100)
	end:=time.Now().Sub(start)
	fmt.Println(end)
}