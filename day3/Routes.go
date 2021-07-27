package Routes

import (
	"bootcamp/api3/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {

	r := gin.Default()
	grp1 := r.Group("/school")
	{

		grp1.GET("student",Controllers.GetStudent)
		grp1.POST("student", Controllers.CreateStudent)
		grp1.GET("student/:id", Controllers.GetStudentByID)
		grp1.PUT("student/:id", Controllers.UpdateStudent)
		grp1.DELETE("student/:id", Controllers.DeleteStudent)

		grp1.GET("studentmarks",Controllers.GetStudentMarks)
		grp1.POST("studentmarks", Controllers.CreateStudentMark)
		grp1.GET("studentmarks/:id/:subject", Controllers.GetStudentMarkByID)
		grp1.PUT("studentmarks/:id/:subject", Controllers.UpdateStudentMark)
		//grp1.DELETE("studentmarks/:id/:subject", Controllers.DeleteStudentMark)




	}
	return r
}
