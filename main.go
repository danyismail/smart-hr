package main

import(
	"os"
	// "fmt"
	"github.com/gin-gonic/gin"
	"smart-hr/modules/users/api"
	Reports "smart-hr/modules/reports/api"
	Employee "smart-hr/modules/employees/api"
	Company "smart-hr/modules/companies/api"
	Component "smart-hr/modules/components/api"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main(){
	// config.Init()
	r := gin.Default()

	
	
	UserController := api.UserController{}
	ReportController := Reports.ReportsController{}
	EmployeeController := Employee.EmployeeController{}
	CompanyController := Company.CompanyController{}
	ComponentController := Component.ComponentController{}

	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers")
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))

	r.GET("/", UserController.Hello)

	r.POST("/auth", UserController.Login)
	r.POST("/logout/:id", UserController.Logout)
	
	r.POST("/users", UserController.Add)
	r.GET("/users", UserController.GetAll)
	
	r.GET("/employees", EmployeeController.GetAll)
	r.POST("/employees/add", EmployeeController.Add)
	
	r.GET("/reports/bpjs", ReportController.ShowBPJSReport)
	r.GET("/reports/bpjs/:id", ReportController.ShowBPJSReportByID)
	
	r.GET("/reports/sallary", ReportController.SallaryReport)
	r.GET("/reports/sallary/:id", ReportController.SallaryReportByID)
	r.GET("/reports/component/:id", ComponentController.GetByID)
	r.GET("/reports/value/:id", ComponentController.GetComponentValue)
	r.POST("/reports/value/edit", ComponentController.EditComponentValue)
	r.POST("/reports/component/add", ComponentController.AddComponentValue)
	
	r.POST("/companies/add", CompanyController.Add)
	r.GET("/companies", CompanyController.GetAll)
	
	godotenv.Load()
	port := os.Getenv("PORT")
	
	//command
	r.Run("localhost" + ":" + port)
}