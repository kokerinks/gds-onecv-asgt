package testing

import (
	"gds-onecv-asgt/models"
	"gds-onecv-asgt/utils"
)

// SeedData is a function that seeds the database with some initial data
func SeedData() {

	db := utils.DB()

	//clean the database
	db.Exec("DELETE FROM student_teachers")
	db.Exec("DELETE FROM students")
	db.Exec("DELETE FROM teachers")
	
	student1 := models.Student{
		Email: "student1@email.com",
	}
	student2 := models.Student{
		Email: "student2@email.com",
	}
	student3 := models.Student{
		Email: "student3@email.com",
	}
	teacher1 := models.Teacher{
		Email: "teacherX@email.com",
		Students: []models.Student{student1, student2},
	}
	teacher2 := models.Teacher{
		Email: "teacherY@email.com",
		Students: []models.Student{student2, student3},
	}

	db.Create(&student1)
	db.Create(&student2)
	db.Create(&student3)
	db.Create(&teacher1)
	db.Create(&teacher2)
}

