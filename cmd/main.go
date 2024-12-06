package main

import (
	"NewWork/internal/db"
	"NewWork/internal/handlers"
	"NewWork/internal/repository"
	"NewWork/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		panic(err)
	}

	Repo := repository.NewRepository(database)
	UC := usecase.NewUseCase(Repo)
	H := handlers.NewHandlers(UC)
	router := gin.Default()
	router.GET("/")
	router.GET("/users", H.GetUsers)
	router.GET("/users/:id", H.GetUser)
	router.POST("/users", H.AddUser)
	router.PATCH("/users/:id", H.UpdateUser)
	router.DELETE("/users/:id", H.DeleteUser)

	router.GET("/course/:id", H.GetCourse)
	router.POST("/course", H.AddCourse)
	router.PATCH("/course/:id", H.UpdateCourse)
	router.DELETE("/course/:id", H.DeleteCourse)

	router.GET("/coursespecification/:id", H.GetCourseSpecification)
	router.POST("/coursespecification", H.AddCourseSpecification)
	router.DELETE("/coursespecification/:id", H.DeleteCourseSpecification)

	router.Run(":8080")
}
