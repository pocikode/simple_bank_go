package val

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateString(s string, minLength, maxLength int) error {
	n := len(s)
	if n < minLength || n > maxLength {
		return fmt.Errorf("string length must be between %d and %d", minLength, maxLength)
	}

	return nil
}

func ValidateUsername(s string) error {
	if err := ValidateString(s, 3, 100); err != nil {
		return err
	}

	if !isValidUsername(s) {
		return fmt.Errorf("must contain only letters or spaces")
	}

	return nil
}

func ValidateFullName(s string) error {
	if err := ValidateString(s, 3, 100); err != nil {
		return err
	}

	if !isValidFullName(s) {
		return fmt.Errorf("must contain only alphanumeric, digits, or underscores")
	}

	return nil
}

func ValidatePassword(s string) error {
	return ValidateString(s, 6, 100)
}

func ValidateEmail(s string) error {
	if err := ValidateString(s, 5, 100); err != nil {
		return err
	}

	if _, err := mail.ParseAddress(s); err != nil {
		return fmt.Errorf("invalid email address")
	}

	return nil
}

func ValidateEmailId(val int64) error {
	if val <= 0 {
		return fmt.Errorf("must be positive integer")
	}

	return nil
}

func ValidateSecretCode(s string) error {
	return ValidateString(s, 32, 128)
}
