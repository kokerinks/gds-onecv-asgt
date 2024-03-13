package main

import (
	"gds-onecv-asgt/routes"
	"gds-onecv-asgt/testData"
	"gds-onecv-asgt/utils"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Test script requires database of the following details:
// Database name: onecv-db
// User: postgres
// Host: localhost
// Port: 5432
// (Password is empty)

func TestRegister(t *testing.T) {

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	utils.ConnectToDB(true)
	testData.SeedData()

	apiGroup := app.Group("/api")
	routes.ApiHandling(apiGroup)

	tests := []struct {
		description string // description of the test case

		route       string // route path to test
		requestBody string // request body

		expectedCode int    // expected HTTP status code
		expectedBody string // expected response body
	}{
		{
			description: "Register a student to a teacher",

			route: "/api/register",
			requestBody: `{"teacher": "teacherW@gmail.com",
											"students": ["student1@gmail.com"]
											}`,

			expectedCode: 204,
			expectedBody: "",
		},

		{
			description: "Register a student to a non-existent teacher",

			route: "/api/register",
			requestBody: `{"teacher": "teacherZ@gmail.com",
											"students": ["student1@gmail.com"]
											}`,

			expectedCode: 404,
			expectedBody: `{"message":"Teacher does not exist in database."}`,
		},

		{
			description: "Register a student to a teacher who is already associated with the student",

			route: "/api/register",
			requestBody: `{"teacher": "teacherX@gmail.com",
											"students": ["student1@gmail.com"]
											}`,

			expectedCode: 204,
			expectedBody: "",
		},

		{
			description: "Register multiple students to a teacher",

			route: "/api/register",
			requestBody: `{"teacher": "teacherW@gmail.com",
											"students": ["student2@gmail.com", "student3@gmail.com"]
											}`,

			expectedCode: 204,
			expectedBody: "",
		},

		{
			description: "Invalid request body (students field misspelled)",

			route: "/api/register",
			requestBody: `{"teacher": "teacherY@gmail.com"
											"student": ["student2@gmail.com", "student3@gmail.com"]
											}`,

			expectedCode: 400,
			expectedBody: `{"message":"Invalid JSON input"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			req := httptest.NewRequest("POST", test.route, strings.NewReader(test.requestBody))
			req.Header.Set("Content-Type", "application/json")

			res, _ := app.Test(req)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		})
	}

}

func TestCommonStudents(t *testing.T) {

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	utils.ConnectToDB(true)
	testData.SeedData()

	apiGroup := app.Group("/api")
	routes.ApiHandling(apiGroup)

	tests := []struct {
		description string // description of the test case

		route string // route path to test

		expectedCode int    // expected HTTP status code
		expectedBody string // expected response body
	}{
		{
			description: "Get 1 common student for a single teacher",

			route: "/api/commonstudents?teacher=teacherX@gmail.com",

			expectedCode: 200,
			expectedBody: `{"students":["student1@gmail.com"]}`,
		},

		{
			description: "Get multiple common students for a single teacher",

			route: "/api/commonstudents?teacher=teacherY@gmail.com",

			expectedCode: 200,
			expectedBody: `{"students":["student2@gmail.com, student3@gmail.com"]}`,
		},

		{
			description: "Get common students for a non-existent teacher",

			route: "/api/commonstudents?teacher=teacherZ@gmail.com",

			expectedCode: 404,
			expectedBody: `{"message":"Teacher(s) [teacherZ@gmail.com] does not exist in database."}`,
		},

		{
			description: "Get common students for multiple teachers",

			route: "/api/commonstudents?teacher=teacherX@gmail.com&teacher=teacherY@gmail.com",

			expectedCode: 200,
			expectedBody: `{"students":["student1@gmail.com"]}`,
		},

		{
			description: "Get common students for multiple teachers, one of which is non-existent",

			route: "/api/commonstudents?teacher=teacherY@gmail.com&teacher=teacherZ@gmail.com",

			expectedCode: 404,
			expectedBody: `{"message":"Teacher(s) [teacherZ@gmail.com] does not exist in database."}`,
		},

		{
			description: "Get common students for multiple teachers, both of which are non-existent",

			route: "/api/commonstudents?teacher=teacherW@gmail.com&teacher=teacherZ@gmail.com",

			expectedCode: 404,
			expectedBody: `{"message":"Teacher(s) [teacherW@gmail.com teacherW@gmail.com] does not exist in database."}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			req := httptest.NewRequest("GET", test.route, nil)

			res, _ := app.Test(req)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		})
	}

}

func TestSuspend(t *testing.T) {

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	utils.ConnectToDB(true)
	testData.SeedData()

	apiGroup := app.Group("/api")
	routes.ApiHandling(apiGroup)

	tests := []struct {
		description string // description of the test case

		route       string // route path to test
		requestBody string // request body

		expectedCode int    // expected HTTP status code
		expectedBody string // expected response body
	}{
		{
			description: "Suspend a student",

			route:       "/api/suspend",
			requestBody: `{"student": "student2@gmail.com"}`,

			expectedCode: 204,
			expectedBody: "",
		},

		{
			description: "Suspend a non-existent student",

			route:       "/api/suspend",
			requestBody: `{"student": "student4@gmail.com"}`,

			expectedCode: 404,
			expectedBody: `{"message":"Student does not exist in database."}`,
		},

		{
			description: "Invalid request body (student field misspelled)",

			route:       "/api/suspend",
			requestBody: `{"teacher": "student2@gmail.com"}`,

			expectedCode: 400,
			expectedBody: `{"message":"Invalid JSON input"}`,
		},

		{
			description: "Suspend a student who is already suspended",

			route:       "/api/suspend",
			requestBody: `{"student": "student2@gmail.com"}`,

			expectedCode: 204,
			expectedBody: "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			req := httptest.NewRequest("POST", test.route, strings.NewReader(test.requestBody))
			req.Header.Set("Content-Type", "application/json")

			res, _ := app.Test(req)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		})
	}

}

func TestRetrieveForNotifications(t *testing.T) {

	if err := godotenv.Load(".env"); err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()
	utils.ConnectToDB(true)
	testData.SeedData()

	apiGroup := app.Group("/api")
	routes.ApiHandling(apiGroup)

	tests := []struct {
		description string // description of the test case

		route       string // route path to test
		requestBody string // request body

		expectedCode int    // expected HTTP status code
		expectedBody string // expected response body
	}{
		{
			description: "Retrieve students for a teacher with no mentioned students",

			route: "/api/retrievefornotifications",
			requestBody: `{"teacher": "teacherX@gmail.com",
											"notification": "Hello students!"	
											}`,

			expectedCode: 200,
			expectedBody: `{"students":["student2@gmail.com"]}`,
		},

		{
			description: "Retrieve students for a teacher",

			route: "/api/retrievefornotifications",
			requestBody: `{"teacher": "teacherX@gmail.com",
											"notification": "Hello students! @student2@gmail.com"
											}`,

			expectedCode: 200,
			expectedBody: `{"students":["student1@gmail.com", "student2@gmail.com"]}`,
		},

		{
			description: "Retrieve students for a non-existent teacher",

			route: "/api/retrievefornotifications",
			requestBody: `{"teacher": "teacherZ@gmail.com",
											"notification": "Hello students! @student2@gmail.com"
											}`,

			expectedCode: 404,
			expectedBody: `{"message":"Teacher does not exist in database."}`,
		},

		{
			description: "Invalid request body (notification field misspelled)",

			route: "/api/retrievefornotifications",
			requestBody: `{"teacher": "teacherX@gmail.com",
											"notice": "Hello students! @@student2@gmail.com
											}`,

			expectedCode: 400,
			expectedBody: `{"message":"Invalid JSON input"}`,
		},

		{
			description: "Retrieve students for a teacher with no registered students and no mentioned students",

			route: "/api/retrievefornotifications",
			requestBody: `{"teacher": "teacherW@gmail.com",
											"notification": "Hello students!"
											}`,

			expectedCode: 404,
			expectedBody: `{"message":"No students found for teacher."}`,
		},

		{
			description: "Retrieve students for a teacher with no registered students and mentioned students",

			route: "/api/retrievefornotifications",
			requestBody: `{"teacher": "teacherW@gmail.com",
											"notification": "Hello students! @student2@gmail.com"
											}`,

			expectedCode: 200,
			expectedBody: `{"students":[student2@gmail.com]}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {

			req := httptest.NewRequest("POST", test.route, strings.NewReader(test.requestBody))
			req.Header.Set("Content-Type", "application/json")

			res, _ := app.Test(req)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
		})
	}
}
