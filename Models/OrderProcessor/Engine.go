package OrderProcessor

import (
	"bootcamp/commerce/Models/Order"
	"fmt"
	"math"
	"sync"
	"time"
)

// map for each user's last executed order
//currentOrder = contains all orders
//pending orders=>time
// 5 worker

func removeFromQueue(slice []uint, s int) []uint {
	return append(slice[:s], slice[s+1:]...)
}

func AllotQueue(){

	for index,currentOrder:=range(Queue){

       lastExecution,found:=OrderTimeMap[currentOrder]
       hours:=lastExecution.Sub(time.Now()).Hours()
       _,mf:=math.Modf(hours)
       ms := mf * 60
       ms, mf = math.Modf(ms)

       if  !found || math.Abs(float64(ms)) >=2 {
           ExecutionQueue.PushBack(currentOrder)
           Queue=removeFromQueue(Queue,index)
	   }
	}

}

func beginProcessing(){

	for ExecutionQueue!=nil{
		element :=ExecutionQueue.Front()
		time.Sleep(time.Second*1)
		fmt.Println("executing order:",element.Value)
		Order.UpdateOrderStatus(element.Value.(uint))
		Mutex.Lock()
		OrderTimeMap[element.Value.(uint)]=time.Now()
		Mutex.Unlock()
		ExecutionQueue.Remove(element)
	}

}

func StartQueueProcessing(){
	waitGroup :=sync.WaitGroup{}
	for ExecutionQueue.Front()!=nil {
		for index:=0;index<noOfWorkers;index++{
			waitGroup.Add(1)
			defer waitGroup.Done()
			go beginProcessing()
		}
	}
	waitGroup.Wait()
}

func QueueAlloter(){

	for len(Queue)!=0 {
		go AllotQueue()
		time.Sleep(time.Second*10)
	}
}

func StartEngine(){

	QueueAlloter()
	StartQueueProcessing()


}



