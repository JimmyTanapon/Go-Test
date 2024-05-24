package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestUserRoute(t *testing.T) {
	app := setup()
	tests := []struct {
		description  string
		requestBody  User
		expectstatus int
	}{
		{
			description:  "valid input",
			requestBody:  User{Email: "Jumonji@gmail.com", Fullname: "JUPIASUW", Age: 22},
			expectstatus: fiber.StatusOK,
		},
		{
			description:  "Invalid email",
			requestBody:  User{Email: "Invalid-Email", Fullname: "JUPIASUW", Age: 22},
			expectstatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid age",
			requestBody:  User{Email: "Jumonji@gmail.com", Fullname: "JUPIASUW", Age: -22},
			expectstatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid Fullname",
			requestBody:  User{Email: "Jumonji@gmail.com", Fullname: "JUPIASUW", Age: -22},
			expectstatus: fiber.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest("POST", "/users", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, tc.expectstatus, resp.StatusCode)
		})
	}
}
