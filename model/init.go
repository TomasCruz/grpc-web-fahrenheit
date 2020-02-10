package model

import (
	"fmt"
)

var (
	// ErrClient signifies client error
	ErrClient error
	// ErrInvalidInput signifies invalid input
	ErrInvalidInput error
)

func init() {
	ErrClient = fmt.Errorf("Client error")
	ErrInvalidInput = fmt.Errorf("Invalid input")
}
