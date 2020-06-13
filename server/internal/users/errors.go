package users

type IncorrectEmailOrPasswordError struct{}

func (m *IncorrectEmailOrPasswordError) Error() string {
	return "incorrect password or email"
}
