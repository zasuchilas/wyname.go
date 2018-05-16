package utils

import (
	"fmt"
	"wyname/kernel"
)

// Abc Probe
func Abc(s string) error {
	fmt.Println("abc use main var", s, kernel.Qwer)
	return nil
}
