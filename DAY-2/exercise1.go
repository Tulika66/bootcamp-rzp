package main
import(
	"sync"
	"fmt"

    "time"
	"strings"
)

type stringsManager struct {
	mutex                sync.Mutex
	frequencyOfAlphabets map[string]int
	
	stringSet            []string
}


func (stringManager *stringsManager)frequencyCalculatorForSet(id int, stringVal string){

	

	for indexAtString:=0;indexAtString<len(stringVal);indexAtString++{

		stringManager.mutex.Lock()
		fmt.Println("Current worker holding lock:",id)
		stringManager.frequencyOfAlphabets[string(stringVal[indexAtString])]+=1
		stringManager.mutex.Unlock()
		
		time.Sleep(time.Millisecond)//to avoid temporary starvation

	}
}

func main() {

    var frequencyOfAlphabets=make(map[string]int,26)
	var  waitGroup =sync.WaitGroup{}
    stringset :=[]string{}
    var numOfStrings int=0

    fmt.Println("Enter Number of strings,followed by all strings")
    fmt.Scanln(&numOfStrings)

    for iteration:=0;iteration<numOfStrings;iteration++{
    	var tempString string
    	fmt.Scanln(&tempString)

    	if len(tempString)==0{
    		continue
		}

    	tempString=strings.ToLower(tempString)
    	stringset=append(stringset,tempString)

	}
	//new instance
    stringManager:=stringsManager{
         mutex:sync.Mutex{},
         stringSet:stringset,
         frequencyOfAlphabets: frequencyOfAlphabets,
         
	}


    fmt.Println("current string manager:-",stringManager)
	for id,stringValue:=range stringManager.stringSet{
		func (){
		stringManager.waitGroup.Add(1)
		defer waitGroup.Done()
		go stringManager.frequencyCalculatorForSet(id,stringValue)
		}()

	}

    stringManager.waitGroup.Wait()//dont know how much time each will take,hence waitgroup used
    fmt.Println("Frequency chart of alphabets : ",frequencyOfAlphabets)

}
