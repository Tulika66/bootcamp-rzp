package main

import (
	"bootcamp/api3/Config"
	"bootcamp/api3/Models"
	"bootcamp/api3/Routes"
	"fmt"
	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	//"github.com/jinzhu/gorm"
	"gorm.io/gorm"
)
// Add mutex for transaction
//error structure
//sql injection ? handle it and remove it
//db
//chron job vs channel


func main() {

	var err1 error
	Config.DBStudent, err1  = gorm.Open(mysql.Open(Config.DBUrl(Config.BuildConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})//gorm.Open("mysql", Config.DBUrl(Config.BuildConfig()))
	//Config.DBStudentMarks, err2 = gorm.Open("mysql", Config.DBUrl(Config.BuildConfig()))

	Config.DBStudentMarks=Config.DBStudent
	if err1 != nil   {
		fmt.Println("Error in db connection of Students.Status :", err1)
	}

	//if err2 !=nil {
	//	fmt.Println("Error in db connection of StudentsMarks.Status :", err2)
	//}
	//defer Config.DBStudent.Close()
	//defer Config.DBStudentMarks.Close()


	Config.DBStudent.AutoMigrate(&Models.Student{})
	Config.DBStudent.AutoMigrate(&Models.Marks{})
	r := Routes.SetupRouter()
	//running
	r.Run(":8080")

}

