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
	mutex sync.Mutex
}

func (teacher *Teacher) workerTasks(workerId int,waitGroup *sync.WaitGroup){
	defer waitGroup.Done()

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
		waitGroup.Add(1)
		go teacher.workerTasks(workerId,&waitGroup)
	}


    waitGroup.Wait()

	time.Sleep(time.Second*5)
    fmt.Println("Average Rating:-",float64(float64(teacher.totalRating)/float64(numberOfStudents)))


}
