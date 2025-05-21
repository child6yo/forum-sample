package service

import (
	"testing"

	"github.com/child6yo/forum-sample"
)

type MockRepository struct {
	mockCreateUser func(user forum.User) (int, error)
	mockGetUser    func(username, password string) (forum.User, error)
}

func (m *MockRepository) CreateUser(user forum.User) (int, error) {
    return m.mockCreateUser(user)
}

func (m *MockRepository) GetUser(username, password string) (forum.User, error) {
    return m.mockGetUser(username, password)
}


func TestCreateUser(t *testing.T) {
	mockRepo := &MockRepository{
        mockCreateUser: func(user forum.User) (int, error) {
            return 1, nil
        },
        mockGetUser: func(username, password string) (forum.User, error) {
            return forum.User{Id: 1, Username: username, Password: password}, nil
        },
    }
	authService := NewAuthService(mockRepo)

	testCases := []struct {
		name    string
		input   forum.User
		wantErr bool
	}{
		{
			name:    "OK",
			input:   forum.User{Id: 0, Email: "user@gmail.com", Username: "user", Password: "123"},
			wantErr: false,
		},
		{
			name:    "invalid email",
			input:   forum.User{Id: 0, Email: "user", Username: "user", Password: "123"},
			wantErr: true,
		},
		{
			name:    "invalid username",
			input:   forum.User{Id: 0, Email: "user@gmail.com", Username: "supercoolusername", Password: "123"},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		_, err := authService.CreateUser(test.input)
		if err == nil && test.wantErr {
			t.Errorf("Test: %s failed. Unexpected error: %s", test.name, err)
		}
	}
}
