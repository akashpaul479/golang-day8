package day8

import "fmt"

func ValidateAge(age int) error {
	if age < 0 || age > 100 {
		return fmt.Errorf("Invalid age %d", age)
	}
	return nil
}
