package validation

import (
	"errors"
	"reflect"
	"testing"
)

func TestValidateStruct(t *testing.T) {
	type User struct {
		name  string `validation:"symbols=username,max_len=10,min_len=2"`
		email string `validation:"email"`
	}

	testCases := []struct {
		name    string
		input   User
		wantErr bool
	}{
		{
			name:    "OK",
			input:   User{name: "username", email: "username@gmail.com"},
			wantErr: false,
		},
		{
			name:    "Invalid symbols",
			input:   User{name: "—u$$Rn@^^e", email: "username@gmail.com"},
			wantErr: true,
		},
		{
			name:    "Invalid max lenght",
			input:   User{name: "coolusername", email: "username@gmail.com"},
			wantErr: true,
		},
		{
			name:    "Invalid min lenght",
			input:   User{name: "u", email: "username@gmail.com"},
			wantErr: true,
		},
		{
			name:    "Invalid email",
			input:   User{name: "username", email: "username"},
			wantErr: true,
		},
	}

	for _, test := range testCases {
		v := Validate{}
		err := v.ValidateStruct(test.input)
		if err == nil && test.wantErr {
			t.Errorf("Test: %s failed. Unexpected error: %s", test.name, err)
		}
	}
}

func TestCustomStruct(t *testing.T) {
	testCases := []struct {
		name        string
		input       interface{}
		customTag   string
		customRules map[string]func(value string, field reflect.Value) error
		wantErr     bool
	}{
		{
			name: "OK",
			input: struct {
				field string `c_validation:"work=true"`
			}{field: "bebe"},
			customRules: map[string]func(value string, field reflect.Value) error{
				"work": func(value string, field reflect.Value) error {
					if value != "true" {
						return errors.New("work ist't true")
					} else if field.String() != "bebe" {
						return errors.New("string ist't bebe")
					}
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid string value",
			input: struct {
				field string `c_validation:"work=false"`
			}{field: "invalid"},
			customRules: map[string]func(value string, field reflect.Value) error{
				"work": func(value string, field reflect.Value) error {
					if value != "true" {
						return errors.New("work ist't true")
					} else if field.String() != "bebe" {
						return errors.New("string ist't bebe")
					}
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "Invalid tag value",
			input: struct {
				field string `c_validation:"work=false"`
			}{field: "bebe"},
			customRules: map[string]func(value string, field reflect.Value) error{
				"work": func(value string, field reflect.Value) error {
					if value != "true" {
						return errors.New("work ist't true")
					} else if field.String() != "bebe" {
						return errors.New("string ist't bebe")
					}
					return nil
				},
			},
			wantErr: true,
		},
		{
			name:  "No custom tag",
			input: struct{ field string }{field: "bebe"},
			customRules: map[string]func(value string, field reflect.Value) error{
				"work": func(value string, field reflect.Value) error {
					if value != "true" {
						return errors.New("work ist't true")
					} else if field.String() != "bebe" {
						return errors.New("string ist't bebe")
					}
					return nil
				},
			},
			wantErr: false,
		},
	}

	for _, test := range testCases {
		v := Validate{customRules: test.customRules}
		err := v.CustomValidateStruct(test.input)
		if err == nil && test.wantErr {
			t.Errorf("Test: %s failed. Unexpected error: %s", test.name, err)
		}
	}
}
