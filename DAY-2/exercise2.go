package main

import(
	"sync"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

const (
	numberOfStudents=10
	sleepTimeLimit=4
	maxRating=10
	minRating=1
)


type Teacher struct{
	totalRating int64
	
}

func (teacher *Teacher) workerTasks(workerId int){
	        

		sleepTime:=rand.Intn(sleepTimeLimit)
		time.Sleep(time.Duration(sleepTime)*time.Second)
		currentStudentRating:=int64(rand.Intn(maxRating-minRating+1)+minRating)
                atomic.AddInt64(&teacher.totalRating,currentStudentRating)
                fmt.Println("Student ID-:",workerId,", Student rating :",currentStudentRating)

	}



func main() {

	waitGroup:=sync.WaitGroup{}

	teacher:=Teacher{
		mutex : sync.Mutex{}	,
		totalRating: int64(0),
	}



	for workerId:=0;workerId<numberOfStudents;workerId++{
		func (){
			waitGroup.Add(1)
			defer waitGroup.Done()
		        go teacher.workerTasks(workerId)
		}()
	}


    waitGroup.Wait()


    fmt.Println("Average Rating:-",float64(float64(teacher.totalRating)/float64(numberOfStudents)))


}
