package validation

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Error struct {
	Errors map[string]string
}

func (ve Error) Error() string {
	var msgs []string
	for k, v := range ve.Errors {
		msgs = append(msgs, fmt.Sprintf("%v: %v", k, v))
	}
	return strings.Join(msgs, " | ")
}

func ValidateEmail(email string) error {
	re := regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$")
	if re.Match([]byte(email)) {
		return nil
	}
	return errors.New("email is invalid")
}
