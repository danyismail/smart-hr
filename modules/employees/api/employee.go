package api

import(
	"github.com/gin-gonic/gin"
	"smart-hr/modules/employees/repositories"
	"smart-hr/modules/employees/models"
	_ "strconv"
)

type EmployeeController struct{}

func(u *EmployeeController) GetAll(ctx *gin.Context){
	repo := repositories.EmployeeRepositories{}
	result, err := repo.GetAll()
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"messages" : result,
		})
		return
	}
}

func(u *EmployeeController) Add(ctx *gin.Context){
	repo := repositories.EmployeeRepositories{}
	form := models.Employee{}
 
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "BindJSON failed",
			"messages" : err.Error(),
		})
		return
	}
	result, err := repo.Add(form)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : result,
			"messages" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"messages" : result,
		})
		return
	}
}