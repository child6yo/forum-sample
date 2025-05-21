package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/internal/app/service"
	mock_service "github.com/child6yo/forum-sample/internal/app/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
)

func TestSignUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user forum.User)

	testCases := []struct {
		name                 string
		inputBody            string
		inputUser            forum.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"email": "username@gmail.com", "username": "username", "password": "qwerty"}`,
			inputUser: forum.User{
				Email:    "username@gmail.com",
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user forum.User) {
				s.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":200,"data":{"user_id":1}}`,
		},
		{
			name:                 "Wrong Input",
			inputBody:            `{"username": "username"}`,
			inputUser:            forum.User{},
			mockBehavior:         func(s *mock_service.MockAuthorization, user forum.User) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"status":400,"data":"invalid request data"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"email": "username@gmail.com", "username": "username", "password": "qwerty"}`,
			inputUser: forum.User{
				Email:    "username@gmail.com",
				Username: "username",
				Password: "qwerty",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user forum.User) {
				s.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"status":500,"data":"server error"}`,
		},
	}

	for _, test := range testCases {
		// Init Dependencies
		c := gomock.NewController(t)
		defer c.Finish()

		repo := mock_service.NewMockAuthorization(c)
		test.mockBehavior(repo, test.inputUser)

		services := &service.Service{Authorization: repo}
		handler := Handler{services}

		// Init Endpoint
		r := gin.New()
		r.POST("/sign-up", handler.signUp)

		// Create Request
		w := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/sign-up",
			bytes.NewBufferString(test.inputBody))

		// Make Request
		r.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, w.Code, test.expectedStatusCode)
		assert.Equal(t, w.Body.String(), test.expectedResponseBody)
	}
}
