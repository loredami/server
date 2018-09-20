package valueObject

import "errors"

type FullName struct {
	firstName string
	lastName  string
}

var (
	InvalidFirstName = errors.New("firstName length must between 1 and 50 chars")
	InvalidLastName  = errors.New("lastName length must between 1 and 50 chars")
)

func NewFullName(firstName, lastName string) (*FullName, error) {
	if len(firstName) == 0 || len(firstName) > 50 {
		return nil, InvalidFirstName
	}
	if len(lastName) == 0 || len(lastName) > 50 {
		return nil, InvalidLastName
	}

	return &FullName{
		firstName: firstName,
		lastName:  lastName,
	}, nil
}

func (f FullName) FirstName() string {
	return f.firstName
}

func (f FullName) LastName() string {
	return f.lastName
}
