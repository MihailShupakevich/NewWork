package domain

import (
	"time"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Courses struct {
	CourseId              uint                `json:"course_id" gorm:"primaryKey"`
	CoursesExecutionTime  time.Time           `json:"courses_execution_time"`
	CoursesName           string              `json:"courses_name"`
	CourseSpecification   CourseSpecification `json:"course_specification" gorm:"foreignKey:CourseSpecificationID;references:CourseSpecificationId"`
	CourseSpecificationID uint                `json:"course_specification_id" gorm:"foreignKey:CourseSpecificationID;references:CourseSpecificationId"`
}

type ListOfSelectedCourses struct {
	ListOfSelectedCoursesId uint      `json:"list_of_selected_courses_id" gorm:"primaryKey"`
	UserID                  uint      `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	User                    User      `json:"user" gorm:"foreignKey:UserID"`
	Courses                 []Courses `json:"courses" gorm:"many2many:selected_courses"`
}

type Sections struct {
	SectionId              uint                 `json:"section_id" gorm:"primaryKey"`
	SectionName            string               `json:"section_name"`
	SectionSpecification   SectionSpecification `json:"section_specification_id" gorm:"foreignKey:SectionSpecificationID"`
	SectionSpecificationID uint                 `json:"section_specification" gorm:"foreignKey:SectionSpecificationID;references:SectionSpecificationId"`
}

type SectionSpecification struct {
	SectionSpecificationId   uint   `json:"section_specification_id" gorm:"primaryKey"`
	SectionSpecificationName string `json:"section_specification_name"`
}

type CourseSpecification struct {
	CourseSpecificationId   uint   `json:"course_specification_id" gorm:"primaryKey"`
	CourseSpecificationName string `json:"course_specification_name"`
}

type ListOfSelectedSections struct {
	ListOfSelectedSectionsId uint       `json:"list_of_selected_sections_id" gorm:"primaryKey"`
	UserID                   uint       `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	Sections                 []Sections `json:"courses" gorm:"many2many:selected_sections"`
	User                     User       `json:"user" gorm:"foreignKey:UserID"`
}

type TopicsForCourses struct {
	TopicsForCoursesId   uint   `json:"topics_for_courses_id" gorm:"primaryKey"`
	User                 User   `json:"user" gorm:"foreignKey:UserID"`
	UserID               uint   `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	TopicsForCoursesName string `json:"topics_for_courses_name"`
}

type AdditionalMaterial struct {
	TopicsForCoursesID      uint             `json:"topics_for_courses_id" gorm:"foreignKey:TopicsForCoursesID;references:TopicsForCoursesId"`
	TopicsForCourses        TopicsForCourses `json:"topics_for_courses" gorm:"foreignKey:TopicsForCoursesID"`
	AdditionalMaterialTitle string           `json:"additional_material_title"`
	AdditionalMaterial      string           `json:"additional_material"`
}

type IndividualWorks struct {
	SectionID uint     `json:"section_id" gorm:"foreignKey:SectionID;references:SectionId"`
	UserID    uint     `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	User      User     `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Section   Sections `json:"section" gorm:"foreignKey:SectionID;references:SectionId"`
	WorkTitle string   `json:"work_title"`
	Work      string   `json:"work"`
}

type CompetitionsUsers struct {
	UserID        uint         `json:"user_id" gorm:"foreignKey:UserID;references:ID"`
	CompetitionID uint         `json:"competition_id" gorm:"foreignKey:CompetitionID;references:CompetitionsId"`
	User          User         `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Competition   Competitions `json:"competitions" gorm:"foreignKey:CompetitionID;references:CompetitionsId"`
}

type Competitions struct {
	CompetitionsId uint      `json:"competitions_id" gorm:"primaryKey"`
	SectionID      uint      `json:"section_id" gorm:"foreignKey:SectionID;references:SectionId"`
	Section        Sections  `json:"section" gorm:"foreignKey:SectionID;references:SectionId"`
	Date           time.Time `json:"date"`
}

type Teachers struct {
	TeachersId  uint     `json:"teachers_id" gorm:"primaryKey"`
	SectionID   uint     `json:"section_id" gorm:"foreignKey:SectionID;references:SectionId"`
	Section     Sections `json:"section" gorm:"foreignKey:SectionID;references:SectionId"`
	TeachersFIO string   `json:"teachers_fio"`
}

type TimetableOfClasses struct {
	TimetableOfClassesId uint      `json:"timetable_of_classes_id" gorm:"primaryKey"`
	SectionID            uint      `json:"section_id" gorm:"foreignKey:SectionID;references:SectionId"`
	Section              Sections  `json:"section" gorm:"foreignKey:SectionID;references:SectionId"`
	Date                 time.Time `json:"date"`
	TeachersID           uint      `json:"teachers_id" gorm:"foreignKey:TeachersID;references:TeachersId"`
	Teacher              Teachers  `json:"teacher" gorm:"foreignKey:TeachersID;references:TeachersId"`
}
