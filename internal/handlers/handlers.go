package handlers

import (
	"NewWork/internal/domain"
	"NewWork/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	useCase usecase.UseCase
}

type HandlersI interface {
	AddUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUser(c *gin.Context)
	GetUsers(c *gin.Context)
	UpdateUser(c *gin.Context)
	AddCourse(c *gin.Context)
	UpdateCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
	GetCourse(c *gin.Context)
	AddCourseSpecification(c *gin.Context)
	DeleteCourseSpecification(c *gin.Context)
	GetCourseSpecification(c *gin.Context)
}

func NewHandlers(useCase usecase.UseCase) Handlers {
	return Handlers{useCase: useCase}
}

func (h *Handlers) AddUser(c *gin.Context) {
	user := new(domain.User)
	err := c.BindJSON(&user)
	newUser, err := h.useCase.AddUser(*user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Не удается создать пользователя!"})
	}
	c.JSON(200, gin.H{"newUser": newUser})
}
func (h *Handlers) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	deleteUser, err := h.useCase.DeleteUser(id)
	if err != nil {
		c.JSON(401, gin.H{"Удаление пользователя-": "Не удалось удалить пользователя"})
	}
	c.JSON(200, gin.H{"Пользователь удален": deleteUser})
}
func (h *Handlers) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.useCase.GetUser(id)
	if err != nil {
		c.JSON(401, gin.Error{})
	}
	c.JSON(200, gin.H{"user": user})
}
func (h *Handlers) UpdateUser(c *gin.Context) {
	var user domain.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(401, gin.Error{})
	}
	id := c.Param("id")
	updateUser, err := h.useCase.UpdateUser(id, user)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
	}
	c.JSON(200, gin.H{"updateUser": updateUser})
}
func (h *Handlers) GetUsers(c *gin.Context) {
	allUsers, err := h.useCase.GetUsers()
	if err != nil {
		c.JSON(401, gin.H{"error": err})
	}
	c.JSON(200, gin.H{"users": allUsers})
}
func (h *Handlers) AddCourse(c *gin.Context) {
	course := new(domain.Courses)
	err := c.BindJSON(&course)
	newCourse, err := h.useCase.AddCourse(*course)
	if err != nil {
		c.JSON(400, gin.H{"error": "Не удается создать пользователя!"})
	}
	c.JSON(200, gin.H{"newCourse": newCourse})
}
func (h *Handlers) UpdateCourse(c *gin.Context) {
	var course domain.Courses
	err := c.BindJSON(&course)
	if err != nil {
		c.JSON(401, gin.Error{})
	}
	id := c.Param("id")
	updateCourse, err := h.useCase.UpdateCourse(id, course)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
	}
	c.JSON(200, gin.H{"updateCourse": updateCourse})
}
func (h *Handlers) DeleteCourse(c *gin.Context) {
	id := c.Param("id")
	deleteCourse, err := h.useCase.DeleteCourse(id)
	if err != nil {
		c.JSON(401, gin.H{"Удаление пользователя-": "Не удалось удалить пользователя"})
	}
	c.JSON(200, gin.H{"Пользователь удален": deleteCourse})
}

func (h *Handlers) GetCourse(c *gin.Context) {
	id := c.Param("id")
	course, err := h.useCase.GetCourse(id)
	if err != nil {
		c.JSON(401, gin.Error{})
	}
	c.JSON(200, gin.H{"course": course})
}
func (h *Handlers) AddCourseSpecification(c *gin.Context) {
	courseSpecification := new(domain.CourseSpecification)
	err := c.BindJSON(&courseSpecification)
	newCourseSpecification, err := h.useCase.AddCourseSpecification(*courseSpecification)
	if err != nil {
		c.JSON(400, gin.H{"error": "Не удается создать спецификацию курса!!!"})
	}
	c.JSON(200, gin.H{"newCourseSpecification": newCourseSpecification})

}
func (h *Handlers) DeleteCourseSpecification(c *gin.Context) {
	id := c.Param("id")
	deleteCourseSpecification, err := h.useCase.DeleteCourseSpecification(id)
	if err != nil {
		c.JSON(401, gin.H{"Удаление спецификации -": "Не удалось удалить спецификацию курса"})
	}
	c.JSON(200, gin.H{"Удаление спецификации курса": deleteCourseSpecification})

}
func (h *Handlers) GetCourseSpecification(c *gin.Context) {
	id := c.Param("id")
	courseSpecification, err := h.useCase.GetCourseSpecification(id)
	if err != nil {
		c.JSON(401, gin.Error{})
	}
	c.JSON(200, gin.H{"course": courseSpecification})
}
