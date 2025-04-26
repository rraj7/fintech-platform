package service

import (
	"net/mail"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func (s *UserService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *UserService) ValidatePassword(password string) bool {
	var hasUpper, hasLower, hasNumber, hasSpecial bool

	if len(password) < 8 {
		return false
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		case unicode.IsSpace(char):
			return false // No spaces allowed
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func (s *UserService) ValidateUsername(username string) bool {
	if len(username) < 3 {
		return false
	}
	for _, char := range username {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}

func (s *UserService) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (s *UserService) ValidatePhone(phone string) bool {
	if len(phone) < 10 {
		return false
	}
	for _, char := range phone {
		if !(char >= '0' && char <= '9') {
			return false
		}
	}
	return true
}

func (s *UserService) ValidateAddress(address string) bool {
	if len(address) < 5 {
		return false
	}
	for _, char := range address {
		if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') &&
			!(char >= '0' && char <= '9') && char != ' ' && char != ',' && char != '.' && char != '-' {
			return false
		}
	}
	return true
}
func (s *UserService) ValidateDateOfBirth(dob string) bool {
	// Expecting format: YYYY-MM-DD
	parsedDOB, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return false // Invalid date format
	}

	// Check if DOB is in the future
	if parsedDOB.After(time.Now()) {
		return false
	}

	// Check if user is at least 18 years old
	age := time.Now().Year() - parsedDOB.Year()

	// Adjust if birthday hasn't occurred yet this year
	if time.Now().YearDay() < parsedDOB.YearDay() {
		age--
	}

	return age >= 18
}
