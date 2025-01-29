package validation

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type Validator interface {
    ValidateStruct(s interface{}) error
    validateField(field reflect.Value, tag string) error
    CustomValidateStruct(s interface{}) error
    customValidateField(field reflect.Value, tag string) error
}

// Validate struct holds custom validation rules
type Validate struct {
	customRules map[string]func(value string, field reflect.Value) error
}

// NewValidator creates and returns a new Validate instance
func NewValidator() *Validate {
    return &Validate{}
}

// AssignRules assigns custom validation rules to the Validate instance
// Takes in a map of custom rules where the key is the rule name and the value is a function
// that takes a string and a reflect.Value and returns an error if validation fails
func (v *Validate) AssignRules(customRules map[string]func(value string, field reflect.Value) error) {
    v.customRules = customRules
}

// ValidateStruct validates all fields of a struct based on the 'validate' tags.
// Takes in any struct and returns an error if validation fails.
//
// This function uses reflection to iterate over all fields of the struct
// and validate them using the specified tags. If any field fails validation,
// an error is returned immediately.
func (v *Validate) ValidateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("validate")

		if err := v.validateField(field, tag); err != nil {
			return err
		}
	}

	return nil
}

// validateField validates a single field based on the provided tag.
// Takes in the field value and the validation tag, and returns an error if validation fails.
//
// The function parses the validation tag into individual rules and applies
// these rules to the field value. It supports various types of validation
// such as checking for valid username, email, maximum length, and minimum length.
func (v *Validate) validateField(field reflect.Value, tag string) error {
    if tag == "" {
        return nil
    }

    rules := parseTag(tag)

    for key, value := range rules {
        switch key {
        case "symbols":
            if value == "username" && !isValidUsername(field.String()) {
                return errors.New("invalid username")
            }
        case "email":
            if !isValidEmail(field.String()) {
                return errors.New("invalid email")
            }
        case "max_len":
			maxLength := stringToInt(value)
            if field.Len() > maxLength {
                return fmt.Errorf("max length exceeded: field length is %d, max length is %d", field.Len(), maxLength)
            }
		case "min_len":
			minLength := stringToInt(value)
            if field.Len() < minLength {
                return fmt.Errorf("min length not met: field length is %d, min length is %d", field.Len(), minLength)
            }
        }
    }

    return nil
}

// parseTag parses a validation tag and returns a map of rules.
// Takes in a validation tag string, and returns a map where the keys are the rule names
// and the values are the rule values.
//
// This function splits the tag string into key-value pairs based on commas and
// equals signs. If a rule has no value, an empty string is used as the value.
func parseTag(tag string) map[string]string {
    rules := map[string]string{}
    pairs := regexp.MustCompile(`,`).Split(tag, -1)

    for _, pair := range pairs {
        parts := regexp.MustCompile(`=`).Split(pair, 2)
        if len(parts) == 2 {
            rules[parts[0]] = parts[1]
        } else {
            rules[parts[0]] = ""
        }
    }

    return rules
}

func isValidUsername(username string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
    return re.MatchString(username)
}

func isValidEmail(email string) bool {
    re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return re.MatchString(email)
}

func stringToInt(s string) int {
    value, _ := strconv.Atoi(s)
    return value
}


// CustomValidateStruct validates all fields of a struct based on custom validation rules.
// Takes in any struct and returns an error if validation fails.
//
// This function is similar to ValidateStruct but uses custom validation rules
// specified in the 'c_validate' tags. If any field fails validation,
// an error is returned immediately.
func (v *Validate) CustomValidateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag.Get("c_validate")

		if err := v.customValidateField(field, tag); err != nil {
			return err
		}
	}

	return nil
}

// customValidateField validates a single field based on custom rules.
// Takes in the field value and the validation tag, and returns an error if validation fails.
//
// The function parses the validation tag into individual custom rules and applies
// these rules to the field value using custom validation functions defined in
// the Validate struct.
func (v *Validate) customValidateField(field reflect.Value, tag string) error {
	if tag == "" {
        return nil
    }

	rules := parseTag(tag)

    for key, value := range rules {
		if err := v.customRules[key](value, field); err != nil {
			return fmt.Errorf("field %s validation failed: %s", value, err)
		}
	}

	return nil
}
