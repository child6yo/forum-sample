package handler

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/internal/app/service"
	mock_service "github.com/child6yo/forum-sample/internal/app/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	"go.uber.org/mock/gomock"
)

func TestCreatePost(t *testing.T) {
	type mockBehavior func(s *mock_service.MockPosts, post forum.Posts)

	testCases := []struct {
		name                 string
		inputBody            string
		inputPost            forum.Posts
		auth_flag            bool
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"title":"title", "content":"content"}`,
			inputPost: forum.Posts{
				UserId:  1,
				Title:   "title",
				Content: "content",
				CrTime:  time.Time{},
				Update:  false,
				UpdTime: time.Time{}},
			auth_flag: true,
			mockBehavior: func(s *mock_service.MockPosts, post forum.Posts) {
				s.EXPECT().CreatePost(post).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"status":200,"data":{"post_id":1}}`,
		},
		{
			name:      "Wrong Authorization",
			inputBody: `{"title":"title", "content":"content"}`,
			inputPost: forum.Posts{
				UserId:  1,
				Title:   "title",
				Content: "content",
				CrTime:  time.Time{},
				Update:  false,
				UpdTime: time.Time{}},
			auth_flag:            false,
			mockBehavior:         func(s *mock_service.MockPosts, post forum.Posts) {},
			expectedStatusCode:   403,
			expectedResponseBody: `{"status":403,"data":"unknown user"}`,
		},
		{
			name:      "Wrong Input",
			inputBody: `{"title":"title"}`,
			inputPost: forum.Posts{
				UserId:  1,
				Title:   "title",
				Content: "content",
				CrTime:  time.Time{},
				Update:  false,
				UpdTime: time.Time{}},
			auth_flag:            true,
			mockBehavior:         func(s *mock_service.MockPosts, post forum.Posts) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"status":400,"data":"invalid request body"}`,
		},
		{
			name:      "Internal error",
			inputBody: `{"title":"title", "content":"content"}`,
			inputPost: forum.Posts{
				UserId:  1,
				Title:   "title",
				Content: "content",
				CrTime:  time.Time{},
				Update:  false,
				UpdTime: time.Time{}},
			auth_flag: true,
			mockBehavior: func(s *mock_service.MockPosts, post forum.Posts) {
				s.EXPECT().CreatePost(post).Return(0, errors.New("server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"status":500,"data":"server error"}`,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPosts(c)
			test.mockBehavior(repo, test.inputPost)

			services := &service.Service{Posts: repo}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/posts", func(c *gin.Context) {
				if test.auth_flag {
					c.Set("userId", 1)
				}
			}, handler.createPost)

			// Init Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/posts",
				bytes.NewBufferString(test.inputBody))

			r.ServeHTTP(w, req)

			// Asserts
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}