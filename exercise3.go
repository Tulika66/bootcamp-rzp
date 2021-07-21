package main

import (
	"fmt"
	"reflect"
)

type Employee interface {
	salary() float64

}

type FullTime struct{
	basic float64
	noOfMonths int64
}

type Contractor struct{
    basic float64
    noOfMonths int64
}

type FreeLancer struct{
	basic float64
	noOfHours float64
}


// Employee interface is implemented by the FullTime
// Method calculates salary of fulltime
func (fte FullTime) salary()float64{
	return fte.basic * float64(fte.noOfMonths) * float64(28)
}


// Employee interface is implemented by the Contractor
// Method calculates salary of contractor
func (contractor Contractor) salary()float64{
	return contractor.basic * float64(contractor.noOfMonths)*float64(28)
}


// Employee interface is implemented by the Freelancer employee
// Method calculates salary of freelancer
func (freelancer FreeLancer)salary()float64{
	return freelancer.basic * freelancer.noOfHours
}


// CalculateNetLiability :Calculates how much does a company need to pay to its Employees
func CalculateNetLiability(employees[] Employee){
	var TotalLiability float64 =0
	for _,employee := range employees{
		currentSalary :=employee.salary()
		currentType :=  reflect.TypeOf(employee)
		TotalLiability+=currentSalary
		fmt.Printf("Salary of employee type = %s is equal to = %f \n",currentType,currentSalary)
	}
	fmt.Println("Company owes total liability of -: ",TotalLiability)
}


func main() {

	freelancer :=FreeLancer{500,17}
	fte := FullTime{500,10}
	contractor := Contractor{500,23}

	employees:= []Employee{freelancer,fte,contractor}
	CalculateNetLiability(employees)
}
