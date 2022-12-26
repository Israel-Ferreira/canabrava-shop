package services

import "testing"

func EmailValidationTestHelper(t *testing.T, email string, expectedValue bool) {
	t.Helper()

	isValid := EmailIsValid(email)

	if isValid != expectedValue {
		t.Errorf("Expected: %v, Result: %v", expectedValue, isValid)
	}

}

func TestEmailIsValid(t *testing.T) {
	t.Run("Deve retornar false ao validar o email @example.com", func(t *testing.T) {
		email := "@example.com"
		EmailValidationTestHelper(t, email, false)
	})

	t.Run("Deve retornar false ao validar o email example.com", func(t *testing.T) {
		email := "example.com"
		EmailValidationTestHelper(t, email, false)
	})

	t.Run("Deve retornar true ao validar o email example@example.com", func(t *testing.T) {
		email := "example@example.com"
		EmailValidationTestHelper(t, email, true)
	})
}
