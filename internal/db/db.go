package db

import (
	"NewWork/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=admin password=admin dbname=newwork port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&domain.User{},
		&domain.Courses{},
		&domain.ListOfSelectedCourses{},
		&domain.Sections{},
		&domain.SectionSpecification{},
		&domain.ListOfSelectedSections{},
		&domain.TopicsForCourses{},
		&domain.AdditionalMaterial{},
		&domain.Teachers{},
		&domain.Competitions{},
		&domain.IndividualWorks{},
		&domain.TimetableOfClasses{},
		&domain.CompetitionsUsers{})
	return db, nil
}
