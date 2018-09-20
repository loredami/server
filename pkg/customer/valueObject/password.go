package valueObject

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hashedPassword []byte
}

func NewPasswordFromHashed(hashedPassword []byte) (*Password, error) {
	_, err := bcrypt.Cost(hashedPassword)
	if err != nil {
		return nil, errors.Wrap(err, "The given password is not valid")
	}

	return &Password{
		hashedPassword: hashedPassword,
	}, nil
}

func NewPasswordFromRaw(password string) (*Password, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "Impossible generate new password")
	}

	return &Password{
		hashedPassword: hashedPassword,
	}, nil
}

func (p Password) HashedPassword() []byte {
	return p.hashedPassword
}
