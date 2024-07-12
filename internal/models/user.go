package models

type UserRegistration struct {
	Username string
	Email    string
	Password string
}

func (u *UserRegistration) Validate() bool {
	return false
}
