package auth

type AccessDeniedError struct{}

func (m *AccessDeniedError) Error() string {
	return "access denied"
}
