package api

import(
	"github.com/gin-gonic/gin"
	"strconv"
	"hr/modules/users/repositories"
	UserModel "hr/modules/users/models"
)

type UserController struct{}




func(u *UserController) GetAll(ctx *gin.Context){
	repo := repositories.UserRepositories{}
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

func(u *UserController) Login(ctx *gin.Context){
	repo := repositories.UserRepositories{}
	form := UserModel.LoginForm{}
	
	err := ctx.ShouldBindJSON(&form)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"data" : err.Error(),
		})
		return
	}
	result, err := repo.Login(form)
	if err != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"data" : err.Error(),
		})
		return
	} else {
		ctx.JSON(200,gin.H{
			"status" : "success",
			"data" : result,
		})
		return
	}
}

func(u *UserController) Logout(ctx *gin.Context){
	repo := repositories.UserRepositories{}
	userID := ctx.Param("id")
	i, e := strconv.Atoi(userID)
	if e != nil {
		ctx.JSON(400,gin.H{
			"status" : "failed",
			"messages" : e,
		})
		return
	}
	result, err := repo.Logout(i)
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

func(u *UserController) Add(ctx *gin.Context){
	repo := repositories.UserRepositories{}
	form := UserModel.User{}
 
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