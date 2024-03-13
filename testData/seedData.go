package testData

import (
	"gds-onecv-asgt/models"
	"gds-onecv-asgt/utils"
)

func SeedData() {

	db := utils.DB()

	db.Exec("DELETE FROM student_teachers")
	db.Exec("DELETE FROM students")
	db.Exec("DELETE FROM teachers")

	student1 := models.Student{
		Email: "student1@gmail.com",
	}
	student2 := models.Student{
		Email: "student2@gmail.com",
	}
	student3 := models.Student{
		Email: "student3@gmail.com",
	}
	teacher1 := models.Teacher{
		Email:    "teacherW@gmail.com",
		Students: []models.Student{},
	}
	teacher2 := models.Teacher{
		Email:    "teacherX@gmail.com",
		Students: []models.Student{student1},
	}
	teacher3 := models.Teacher{
		Email:    "teacherY@gmail.com",
		Students: []models.Student{student1, student2},
	}

	db.Create(&student1)
	db.Create(&student2)
	db.Create(&student3)
	db.Create(&teacher1)
	db.Create(&teacher2)
	db.Create(&teacher3)
}
