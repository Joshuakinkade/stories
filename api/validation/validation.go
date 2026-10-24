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

type StringRequirements struct {
	minLength int
	maxLength int
	match     *regexp.Regexp
}

func ValidateString(value any, reqs StringRequirements) (string, error) {
	str, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("%v is not a string", value)
	}

	if len(str) < reqs.minLength {
		return "", fmt.Errorf("%v must be at least %v characters long", value, reqs.minLength)
	}

	if len(str) > reqs.maxLength {
		return "", fmt.Errorf("%v must be no more than %v characters long", value, reqs.maxLength)
	}

	if !reqs.match.Match([]byte(str)) {
		return "", fmt.Errorf("%v doesn't match pattern", value)
	}

	return str, nil
}

func ValidateEmail(value any) (string, error) {
	reqs := StringRequirements{
		match: regexp.MustCompile("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$"),
	}
	email, err := ValidateString(value, reqs)
	if err != nil {
		return "", errors.New("email is invalid")
	}
	return email, nil
}
