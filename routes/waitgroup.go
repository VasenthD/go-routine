package main

import(
	"fmt"
	"sync"
	//"time"
)

func routine(i int,wg *sync.WaitGroup){

	fmt.Println("started ",i)
	fmt.Printf("routing %d ended \n",i)
	wg.Done()
}

func main(){
	noofroutines :=3
	var wg sync.WaitGroup
	for i:=0;i<noofroutines;i++{
		wg.Add(1)
		go routine(i,&wg)
	}
	wg.Wait();
	fmt.Println("ALL ROUTES ARE FINISHED")
}