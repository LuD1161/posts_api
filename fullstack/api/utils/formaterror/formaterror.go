package formaterror

import (
	"errors"
	"strings"
)

// FormatError : Evaluate the error from golang and return
// accordingly user-understandable error
func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("Nickname Already Taken")
	}
	if strings.Contains(err, "emil") {
		return errors.New("Email Already Taken")
	}
	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
