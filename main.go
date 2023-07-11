package main

import (
	"github.com/gin-gonic/gin"
	"golang_basic_gin/middlewares"
	"golang_basic_gin/routes"
)

func main() {
	r := gin.Default()
	//r.GET("/", GetHome)
	v1 := r.Group("api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("/register", routes.RegisterUser)
			user.POST("/login", routes.GenereteToken)
		}
		department := v1.Group("/department").Use(middlewares.Auth())
		{
			department.GET("/", routes.GetDepartment)
			department.POST("/", routes.PostDepartment)
			department.PUT("/put/:id", routes.PutDepartment)
			department.DELETE("/delete/:id", routes.DeleteDepartment)
		}

		position := v1.Group("/position")
		{
			position.GET("/", routes.GetPosition)
			position.POST("/", routes.PostPosition)
			position.PUT("/put/:id", routes.PutPosition)
			position.DELETE("/delete/:id", routes.DeletePosition)

		}

		employee := v1.Group("/employee")
		{
			employee.GET("/", routes.GetEmplpoyee)
			employee.POST("/", routes.PostEmployee)
			employee.PUT("/put/:id", routes.PutEmployee)
			employee.DELETE("/delete/:id", routes.DeleteEmployee)

		}
	}

	//r.GET("/positions", routes.GetPosition)

	r.Run(":3000") // terhubung pada port :  0.0.0.0:8080
}
