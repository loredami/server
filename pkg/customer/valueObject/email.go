package valueObject

import (
	"regexp"

	"github.com/pkg/errors"
)

var (
	emailRegexp  = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	InvalidEmail = errors.New("The given value is not valid")
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {

	if !emailRegexp.MatchString(value) {
		return nil, InvalidEmail
	}

	return &Email{
		value: value,
	}, nil
}

func (e Email) String() string {
	return e.value
}
