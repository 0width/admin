package tests

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGetPassword(t *testing.T) {
	r, err := bcrypt.GenerateFromPassword([]byte("123456"), 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
}
