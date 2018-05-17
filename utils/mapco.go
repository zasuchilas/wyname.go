package utils

import (
	"fmt"
	"wyname/kernel"
)

// Abc Probe
func Abc(s string) {
	fmt.Println("abc use main var", s, kernel.Qwer)
	fmt.Println("Conf", kernel.Conf)
}
