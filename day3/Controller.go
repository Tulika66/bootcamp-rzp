package Controllers

import (
	"bootcamp/api3/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


// GetStudent ... Get all users
func GetStudent(c *gin.Context) {
	var student []Models.Student
	err := Models.GetAllStudents(&student)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// CreateStudent Create User
func CreateStudent(c *gin.Context) {
	var student Models.Student
	c.BindJSON(&student)
	err := Models.CreateStudent(&student)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// GetStudentByID Get the user by id
func GetStudentByID(c *gin.Context) {

	id := c.Params.ByName("id")
	var user Models.Student
	err := Models.GetStudentById(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// UpdateStudent Update the user information
func UpdateStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentById(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// DeleteStudent Delete the user
func DeleteStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
//////////////////////////////database 2 queries //////////

// GetStudentMarks
func GetStudentMarks(c *gin.Context) {
	var studentmarks []Models.Marks
	err := Models.GetAllStudentMarks(&studentmarks)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, studentmarks)
	}
}

// CreateStudentMark
func CreateStudentMark(c *gin.Context) {
	var studentMarks Models.Marks
	c.BindJSON(&studentMarks)
	err := Models.CreateStudentMark(&studentMarks)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, studentMarks)
	}
}

// GetStudentMarkByID
func GetStudentMarkByID(c *gin.Context) {

	id := c.Params.ByName("id")
	var marksOfStudent []Models.Marks
	err := Models.GetStudentMarksById(&marksOfStudent, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, marksOfStudent)
	}

}

// UpdateStudentMark
func UpdateStudentMark(c *gin.Context) {
	var studentMark Models.Marks
	id := c.Params.ByName("id")
	subject:=c.Params.ByName("subject")
	err := Models.GetStudentMarkById(&studentMark, id,subject)

	if err != nil {
		c.JSON(http.StatusNotFound, studentMark)
	}
	c.BindJSON(&studentMark)
	err = Models.UpdateStudentMark(&studentMark, id,subject)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, studentMark)
	}
}
//
