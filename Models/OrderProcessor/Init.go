package OrderProcessor

import (
	"container/list"
	"sync"
	"time"
)

//initialise and start engine

const (
	targetUsers=6
	maxUsersInQueue=3
	noOfWorkers=5
)

var (
	Queue []uint // all orders
	OrderTimeMap map[uint]time.Time
	Mutex sync.Mutex
	ExecutionQueue *list.List
)


// users eligible in a queue
// allot users in
func init (){


	ExecutionQueue=list.New()

	StartEngine()
}
