package day8

import "testing"

func TestValidateAge(t *testing.T) {
	if err := ValidateAge(-5); err == nil {
		t.Errorf("Expected error of age -5, got nil")
	}
	if err := ValidateAge(25); err != nil {
		t.Errorf("Didnot expect error for age 25")
	}
}
