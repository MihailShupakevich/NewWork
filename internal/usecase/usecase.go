package usecase

import (
	"NewWork/internal/domain"
	"NewWork/internal/repository"
	"strconv"
)

type UseCase struct {
	repo repository.Repository
}

type UsecaseI interface {
	AddUser(user domain.User) (domain.User, error)
	DeleteUser(id string) (string, error)
	GetUser(id string) (domain.User, error)
	UpdateUser(id string, updateData domain.User) (domain.User, error)
	GetUsers() ([]domain.User, error)
	AddCourse(course domain.Courses) (domain.Courses, error)
	UpdateCourse(id string, updateDataCourse domain.Courses) (domain.Courses, error)
	DeleteCourse(id string) (string, error)
	GetCourse(id string) (domain.Courses, error)
	AddCourseSpecification(courseSpecification domain.CourseSpecification) (domain.CourseSpecification, error)
	DeleteCourseSpecification(id string) (string, error)
	GetCourseSpecification(id string) (domain.CourseSpecification, error)
}

func NewUseCase(repo repository.Repository) UseCase {
	return UseCase{repo: repo}
}

func (u *UseCase) AddUser(user domain.User) (domain.User, error) {
	newUser, err := u.repo.AddUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return newUser, err
}

func (u *UseCase) DeleteUser(id string) (string, error) {
	IntId, _ := strconv.Atoi(id)
	deletedUser, err := u.repo.DeleteUser(IntId)
	if err != nil {
		return "Error deleted", err
	}
	return deletedUser, err
}
func (u *UseCase) GetUser(id string) (domain.User, error) {
	IntId, _ := strconv.Atoi(id)
	getUser, err := u.repo.GetUser(IntId)
	if err != nil {
		return domain.User{}, err
	}
	return getUser, nil
}
func (u *UseCase) UpdateUser(id string, updateData domain.User) (domain.User, error) {
	IntId, err := strconv.Atoi(id)
	if err != nil {
		return domain.User{}, err
	}
	updateUser, err := u.repo.UpdateUser(IntId, updateData)
	if err != nil {
		return domain.User{}, err
	}
	return updateUser, nil
}
func (u *UseCase) GetUsers() ([]domain.User, error) {
	getUsers, err := u.repo.GetUsers()
	if err != nil {
		return []domain.User{}, err
	}
	return getUsers, nil
}
func (u *UseCase) AddCourse(course domain.Courses) (domain.Courses, error) {
	newCourse, err := u.repo.AddCourse(course)
	if err != nil {
		return domain.Courses{}, err
	}
	return newCourse, err
}
func (u *UseCase) UpdateCourse(id string, updateDataCourse domain.Courses) (domain.Courses, error) {
	IntId, err := strconv.Atoi(id)
	if err != nil {
		return domain.Courses{}, err
	}
	updateCourse, err := u.repo.UpdateCourse(IntId, updateDataCourse)
	if err != nil {
		return domain.Courses{}, err
	}
	return updateCourse, err
}
func (u *UseCase) DeleteCourse(id string) (string, error) {
	IntId, _ := strconv.Atoi(id)
	deletedCourse, err := u.repo.DeleteCourse(IntId)
	if err != nil {
		return "Error deleted", err
	}
	return deletedCourse, err

}
func (u *UseCase) GetCourse(id string) (domain.Courses, error) {
	IntId, _ := strconv.Atoi(id)
	getCourse, err := u.repo.GetCourse(IntId)
	if err != nil {
		return domain.Courses{}, err
	}
	return getCourse, nil
}
func (u *UseCase) AddCourseSpecification(courseSpecification domain.CourseSpecification) (domain.CourseSpecification, error) {
	newCourseSpecification, err := u.repo.AddCourseSpecification(courseSpecification)
	if err != nil {
		return domain.CourseSpecification{}, err
	}
	return newCourseSpecification, err
}
func (u *UseCase) DeleteCourseSpecification(id string) (string, error) {
	IntId, _ := strconv.Atoi(id)
	deletedCourseSpecification, err := u.repo.DeleteCourseSpecification(IntId)
	if err != nil {
		return "Error deleted", err
	}
	return deletedCourseSpecification, err
}
func (u *UseCase) GetCourseSpecification(id string) (domain.CourseSpecification, error) {
	IntId, _ := strconv.Atoi(id)
	getCourseSpecification, err := u.repo.GetCourseSpecification(IntId)
	if err != nil {
		return domain.CourseSpecification{}, err
	}
	return getCourseSpecification, nil
}
