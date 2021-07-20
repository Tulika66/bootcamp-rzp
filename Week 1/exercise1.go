package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct{
	numberOfRows int64
	numberOfColumns int64
	elements [][] int64

}

//Init Constructor to initialise matrix size and the actual matrix elements with that size
//Takes input as dimensions of matrix to be set for the receiver matrix
func (matrix *Matrix) Init(numRows int64,numColumns int64){
	//constructor initialising elements


	//create a temporary variable to define size of elements
	elements :=make([][]int64,matrix.numberOfRows)
	matrix.elements=elements
	for index:=int64(0);index<numRows;index++{
		matrix.elements[index]=make([]int64,numColumns)
	}

}

func (matrix *Matrix) getNumberOfRows()int64{
	return matrix.numberOfRows
}

func (matrix *Matrix) getNumberOfColumns()int64{
	return matrix.numberOfColumns
}

func (matrix *Matrix) setElementAt( row int64, column int64, value int64){
	if int64(matrix.numberOfRows)>row && int64(matrix.numberOfColumns)>column  {
		matrix.elements[row][column]=value
	}
}

//addMatrices adds two existing matrices and stores the result in another third preexisting matrix
func addMatrices(matrix1 *Matrix,matrix2 *Matrix, matrix3 *Matrix) *Matrix{

	for indexRow:=int64(0);indexRow<matrix1.numberOfRows;indexRow++{
		for indexCol:=int64(0);indexCol<matrix1.numberOfColumns;indexCol++{
			matrix3.elements[indexRow][indexCol]=matrix1.elements[indexRow][indexCol]+matrix2.elements[indexRow][indexCol]
		}
	}
    return matrix3

}

//printJsonMatrix converts a matrix to json
func (matrix *Matrix)printJsonMatrix()string {
	jsonMatrix, ok := json.Marshal(matrix.elements)
	if ok!=nil{
		fmt.Println("Json failed")
		return "Failed to convert"
		fmt.Println(ok)
	}
	return string(jsonMatrix)
}
func main() {

matrix:=new (Matrix)
matrix.Init(1,2)
jsonMatrix := matrix.printJsonMatrix()
if jsonMatrix!="Failed to convert"{
	fmt.Println(jsonMatrix)
}else{
	fmt.Println("Converting to JSON of matrix was not possible")
}






}
