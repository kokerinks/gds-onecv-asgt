package controllers

import (
	"gds-onecv-asgt/models"
	"gds-onecv-asgt/utils"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	db := utils.DB()

	type RegisterRequest struct {
		Teacher  string   `json:"teacher"`
		Students []string `json:"students"`
	}

	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid JSON input",
		})
	}

	teacherEmail := req.Teacher
	studentEmails := req.Students

	var teacher models.Teacher
	if err := db.Where("email = ?", teacherEmail).First(&teacher).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Teacher does not exist in database.",
		})
	}

	for _, studentEmail := range studentEmails {
		var student models.Student

		// Check if student exists
		if err := db.Where("email = ?", studentEmail).First(&student).Error; err != nil {
			student = models.Student{
				Email:    studentEmail,
				Teachers: []models.Teacher{teacher},
			}

			db.Create(&student)

			// If student exists, check if teacher is already associated with student
		} else {
			var isFound bool = false

			for _, t := range student.Teachers {
				if t.Email == teacherEmail {
					isFound = true
				}
			}

			if !isFound {
				student.Teachers = append(student.Teachers, teacher)
			}
			db.Save(&student)
		}
	}

	return c.Status(204).JSON(fiber.Map{})
}

func CommonStudents(c *fiber.Ctx) error {

	db := utils.DB()

	query := string(c.Request().URI().QueryString())
	query = strings.Replace(query, "teacher=", "", -1)
	teacherEmails := strings.Split(query, "&")
	log.Println(teacherEmails)

	var students []string
	db.Raw(`
		SELECT studentEmail
		FROM (
			SELECT s.email AS studentEmail, COUNT(DISTINCT t.id) AS cnt
			FROM students AS s
			INNER JOIN student_teachers AS st ON s.id = st.student_id
			INNER JOIN teachers AS t ON t.id = st.teacher_id
			WHERE t.email IN (?)
			GROUP BY s.email
		) AS ret
		WHERE cnt = ?`, teacherEmails, len(teacherEmails)).Scan(&students)

	if len(students) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No common students found",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"students": students,
		})
	}

}

func Suspend(c *fiber.Ctx) error {

	db := utils.DB()

	type SuspendRequest struct {
		Student string `json:"student"`
	}

	req := new(SuspendRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid JSON input",
		})
	}

	studentEmail := req.Student

	var student models.Student
	if err := db.Where("email = ?", studentEmail).First(&student).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Student does not exist in database.",
		})
	}

	student.IsSuspended = true
	db.Save(&student)

	return c.Status(204).JSON(fiber.Map{})
}

func RetrieveForNotifications(c *fiber.Ctx) error {

	db := utils.DB()

	type RetrieveRequest struct {
		Teacher      string `json:"teacher"`
		Notification string `json:"notification"`
	}

	req := new(RetrieveRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid JSON input",
		})
	}

	teacherEmail := req.Teacher
	notification := req.Notification

	// Extract email addresses from notification
	emails := strings.Split(notification, " ")
	var mentionedStudents []string
	for _, email := range emails {
		if strings.HasPrefix(email, "@") {
			mentionedStudents = append(mentionedStudents, strings.Trim(email, "@"))
		}
	}

	var students []string
	db.Raw(`
		SELECT s.email
		FROM students AS s
		INNER JOIN student_teachers AS st ON s.id = st.student_id
		INNER JOIN teachers AS t ON t.id = st.teacher_id
		WHERE t.email = ?
		AND s.is_suspended = FALSE
		UNION
		SELECT s.email
		FROM students AS s
		WHERE s.email in ?
		AND s.is_suspended = FALSE`, teacherEmail, mentionedStudents).Scan(&students)

	if len(students) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No students found",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"recipients": students,
		})
	}
}
