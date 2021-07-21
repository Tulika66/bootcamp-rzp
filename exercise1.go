package main

import (
	"encoding/json"
	"fmt"
	"errors"
)

type Matrix struct{
	numberOfRows int64
	numberOfColumns int64
	elements [][] int64

}

//Init Constructor to initialise matrix size and the actual matrix with that size
//Takes input as dimensions of matrix to be set for the receiver matrix
func (matrix *Matrix) Init(numRows int64,numColumns int64){
	//default declaration can be used,but
	matrix.numberOfRows=numRows
	matrix.numberOfColumns=numColumns

	matrix.elements=make([][]int64,matrix.numberOfRows)
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


//setElementAt sets element at appropriate index
func (matrix *Matrix) setElementAt( row int64, column int64, value int64) {
	if (row >= matrix.numberOfRows || row < 0 || column < 0 || column >= matrix.numberOfColumns) {
		fmt.Println(errors.New("Index Out of Bounds. Attempt to set value failed."))
	} else {
		matrix.elements[row][column] = value
	}
}


//addMatrices adds two existing matrices and stores the result in another third preexisting matrix
//third matrix is used instead of one of original two matrices,to avoid modifying the original ones
//Assumption : Matrice dimensions are assumed to be compatible for addition
func addMatrices(matrix1 *Matrix,matrix2 *Matrix, matrix3 *Matrix) *Matrix{

	for indexRow:=int64(0);indexRow<matrix1.numberOfRows;indexRow++{
		for indexCol:=int64(0);indexCol<matrix1.numberOfColumns;indexCol++{
			matrix3.setElementAt(indexRow,indexCol,matrix1.elements[indexRow][indexCol]+matrix2.elements[indexRow][indexCol])
		}
	}
    return matrix3

}

//printJsonMatrix converts a matrix in json and prints it
func (matrix *Matrix)printJsonMatrix() {
	jsonMatrix, ok := json.Marshal(matrix.elements)
	if ok!=nil{
		fmt.Println("Json failed")
        fmt.Println(ok)
	}else{
		fmt.Println(string(jsonMatrix))
	}
}


func main() {

	matrix:=new (Matrix)
	matrix.Init(1,2)

	matrix.printJsonMatrix()
	matrix.setElementAt(0,1,2)
	matrix.setElementAt(0,0,1)
	matrix.setElementAt(0,3,23)// attempt to set value at an invalid index
	matrix.printJsonMatrix()



}
