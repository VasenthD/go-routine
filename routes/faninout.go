// the purpose of the select is:
//  select statment will allows a go routine to wait on multiple communications operation, it blocks the execution util one of the function on called
//
package main

import(
	"fmt"
	"time"
)

type Item struct{
	ID int
	Name string
	PackingEffort time.Duration
}
func PrepareItems(done<-chan bool)<-chan Item{
	items :=make(chan Item)
	intemsToShip :=[]Item{
		Item{0,"Shirt",1*time.Second},
		Item{1,":egps",1*time.Second},
		Item{2,"Tv",5*time.Second},
		Item{3,"Bananas",2*time.Second},
		Item{4,"Hat",1*time.Second},
		Item{5,"Phone",2*time.Second},
		Item{6,"Plates",3*time.Second},
		Item{7,"Computer",1*time.Second},
		Item{8,"Print Glass",3*time.Second},
		Item{9,"Watch",2*time.Second},
	}
	go func(){
		for _,item:=range intemsToShip{

			select{
			case <-done:
				return
			case items <-item:
			}
			
		}
		close(items)
	}()
	return items
}
func PackItems(done<-chan bool, items<-chan Item)<-chan int{
	packages := make(chan int)
	go func(){
		for item:=range items{
			select{
			case<-done:
				return
			case packages<-item.ID:
				time.Sleep(item.PackingEffort)
			}
		}
		close(packages)

	}()
	return packages

}

func main(){
	done :=make(chan bool)
	defer close(done)
	start :=time.Now()
	packages :=PackItems(done, PrepareItems(done))
	numPackages :=0
	for p:=range packages{
		numPackages++
		fmt.Printf("shipping pakcage no : %d \n",p)

	}
	fmt.Printf("Took %fs to ship %d pakcages \n",time.Since(start).Seconds(),numPackages)
}