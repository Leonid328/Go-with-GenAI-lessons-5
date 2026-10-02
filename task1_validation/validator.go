// Package validation містить завдання 1 домашньої роботи: валідатор форми
// реєстрації, що повертає власний (custom) тип помилки з переліком УСІХ
// невалідних полів, а не лише першого.
package validation

import "strings"

type RegistrationForm struct {
	Email    string
	Password string
	Age      int
}

type ValidationError struct {
	Fields []string
}

func (e *ValidationError) Error() string {
	return "registration invalid: fields " + strings.Join(e.Fields, ", ")
}

func ValidateRegistration(f RegistrationForm) error {
	var fields []string
	if f.Email == "" {
		fields = append(fields, "email")
	}
	if f.Password == "" || len(f.Password) < 8 {
		fields = append(fields, "password")
	}
	if f.Age < 0 || f.Age > 150 {
		fields = append(fields, "age")
	}
	if len(fields) == 0 {
		return nil
	}
	return &ValidationError{Fields: fields}
}
