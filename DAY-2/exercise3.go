package main
import(
	"errors"
	"sync"
	"fmt"
	"time"
)

type Employee struct{
	bankBalance int64
	name string
	mutex sync.Mutex

}

type customError interface{
	WithdrawError()
}

func(employee *Employee) WithdrawError(){
	errors.New("Withdraw Failed,please check balance!")
}

func (employee *Employee)withdrawFromAccount(value int64,waitGroup *sync.WaitGroup){

	waitGroup.Add(1)
	defer waitGroup.Done()


	if value>employee.bankBalance{
		employee.WithdrawError()
		return
	}
	employee.mutex.Lock()

	fmt.Println("Locked.Removing from account:",value)
	employee.bankBalance-=value
	employee.mutex.Unlock()

	fmt.Println("Accountbalance after this withdrawl-:")
	employee.getBalance()

}

func (employee *Employee)AddToAccount(value int64,waitGroup *sync.WaitGroup){

	waitGroup.Add(1)
	waitGroup.Done()

	employee.mutex.Lock()
	fmt.Println("Locked.Adding to account:",value)
	employee.bankBalance+=value
	employee.mutex.Unlock()

	fmt.Println("Accountbalance after this deposit-:")
	employee.getBalance()

}

func (employee *Employee)getBalance(){

	employee.mutex.Lock()
	fmt.Println("current bal:- ",employee.bankBalance)
	employee.mutex.Unlock()
}

func main() {

	waitGroup :=sync.WaitGroup{}
	employee1:=Employee{
		name:"Michael",
		mutex: sync.Mutex{},
		bankBalance: int64(4000),

	}


	go (&employee1).AddToAccount(4000,&waitGroup)
	go (&employee1).AddToAccount(4000,&waitGroup)


	go (&employee1).AddToAccount(4000,&waitGroup)
	go (&employee1).AddToAccount(4000,&waitGroup)
	go (&employee1).withdrawFromAccount(1000,&waitGroup)


	time.Sleep(time.Second*1)

	waitGroup.Wait()



}
