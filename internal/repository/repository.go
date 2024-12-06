package repository

import (
	"NewWork/internal/domain"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type RepositoryI interface {
	AddUser(user domain.User) (domain.User, error)
	DeleteUser(IntId int) (string, error)
	UpdateUser(IntId int, updateData domain.User) (domain.User, error)
	GetUser(IntId int) (domain.User, error)
	GetUsers() ([]domain.User, error)
	AddCourse(course domain.Courses) (domain.Courses, error)
	UpdateCourse(IntId int, updateDataCourse domain.Courses) (domain.Courses, error)
	DeleteCourse(IntId int) (string, error)
	GetCourse(IntId int) (domain.Courses, error)
	AddCourseSpecification(courseSpecification domain.CourseSpecification) (domain.CourseSpecification, error)
	DeleteCourseSpecification(IntId int) (string, error)
	GetCourseSpecification(IntId int) (domain.CourseSpecification, error)
}

func NewRepository(DB *gorm.DB) Repository {
	return Repository{db: DB}
}

func (r *Repository) AddUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
func (r *Repository) DeleteUser(IntId int) (string, error) {
	r.db.Where("ID = ?", IntId).Delete(&domain.User{})
	return "Успешно удален", nil
}

func (r *Repository) UpdateUser(IntId int, updateData domain.User) (domain.User, error) {
	r.db.Where("ID = ?", IntId).Updates(&updateData)
	return updateData, nil
}

func (r *Repository) GetUser(IntId int) (domain.User, error) {
	var user domain.User
	r.db.First(&user, IntId)
	return user, nil
}
func (r *Repository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	r.db.Find(&users)
	return users, nil
}
func (r *Repository) AddCourse(course domain.Courses) (domain.Courses, error) {
	err := r.db.Create(&course).Error
	return course, err
}
func (r *Repository) UpdateCourse(IntId int, updateDataCourse domain.Courses) (domain.Courses, error) {
	r.db.Where("course_id = ?", IntId).Updates(&updateDataCourse)
	return updateDataCourse, nil
}
func (r *Repository) DeleteCourse(IntId int) (string, error) {
	r.db.Where("course_id = ?", IntId).Delete(&domain.Courses{})
	return "Курс успешно удален", nil
}
func (r *Repository) GetCourse(IntId int) (domain.Courses, error) {
	var course domain.Courses
	r.db.Preload("CourseSpecification").First(&course, IntId)
	return course, nil
}
func (r *Repository) AddCourseSpecification(courseSpecification domain.CourseSpecification) (domain.CourseSpecification, error) {
	err := r.db.Create(&courseSpecification).Error
	return courseSpecification, err
}
func (r *Repository) DeleteCourseSpecification(IntId int) (string, error) {
	r.db.Where("course_specification_id = ?", IntId).Delete(&domain.CourseSpecification{})
	return "Спецификация курса успешно удалена", nil
}

func (r *Repository) GetCourseSpecification(IntId int) (domain.CourseSpecification, error) {
	var courseSpecification domain.CourseSpecification
	r.db.First(&courseSpecification, IntId)
	return courseSpecification, nil
}
