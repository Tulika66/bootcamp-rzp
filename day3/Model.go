package Models

import (
	"bootcamp/api3/Config"
	"fmt"
	"gorm.io/gorm/clause"

	////_ "github.com/go-sql-driver/mysql"
	//"gorm.io/driver/mysql"
)

func GetAllStudents(student *[]Student)(err error){

	if err:=Config.DBStudent.Preload(clause.Associations).Find(student).Error;err!=nil{
		return err
	}

	return nil

}

func CreateStudent(student *Student)(err error){
	if err :=Config.DBStudent.Create(student).Error;err!=nil{
		return err
	}
	return nil
}

func GetStudentById(student *Student, id string)(err error){
	if err :=Config.DBStudent.Preload(clause.Associations).Where("id = ?",id).First(student).Error;err!=nil{
		return err
	}
	return nil
}


func UpdateStudent(student *Student, id string) (err error) {
	fmt.Println(student)
	//Config.DB.Save(student)
	if err =Config.DBStudent.Save(student).Error; err!=nil {
		return err
	}
	return nil
}


func DeleteStudent(student *Student, id string) (err error) {
	Config.DBStudent.Where("id = ?", id).Delete(student)
	var studentMarks *[]Marks
	Config.DBStudentMarks.Where("id=?", id).Find(studentMarks)
	if studentMarks != nil {
		Config.DBStudentMarks.Delete(studentMarks)
	}
	return nil
}
////from database 2

func GetAllStudentMarks(studentMarks *[]Marks)(err error){

	if err:=Config.DBStudentMarks.Find(studentMarks).Error;err!=nil{
		return err
	}

	return nil

}

func CreateStudentMark(studentMark *Marks)(err error){
	if err :=Config.DBStudentMarks.Create(studentMark).Error;err!=nil{
		return err
	}
	return nil
}

func GetStudentMarksById(studentMarks *[]Marks, id string)(err error){
	if err :=Config.DBStudentMarks.Where("id = ?",id).Find(studentMarks).Error;err!=nil{
		return err
	}
	return nil
}

func GetStudentMarkById(studentMark *Marks, id string,subject string)(err error){
	if err :=Config.DBStudentMarks.Where("id = ? AND subject =?",id,subject).Find(studentMark).Error;err!=nil{
		return err
	}
	return nil
}


func UpdateStudentMark(studentMarks *Marks, id string,subject string) (err error) {
	fmt.Println(studentMarks)
	//Config.DB.Save(student)
	if err =Config.DBStudentMarks.Where("id =? AND subject =?",id,subject).Save(studentMarks).Error; err!=nil {
		return err
	}
	return nil
}


func DeleteStudentSubject(studentMarks *Marks, id string,subject string ) (err error) {
	Config.DBStudentMarks.Where("id = ? AND subject =?", id,subject).Delete(studentMarks)
	return nil
}

