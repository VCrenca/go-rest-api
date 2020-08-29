package auth

// Authenticator -
type Authenticator interface {
	CheckEmailAndPassword(email string, password string) (bool, error)
	CheckToken(token string) (bool, error)
}

type authenticator struct {
}

func (a authenticator) CheckEmailAndPassword(email string, password string) (bool, error) {
	return false, nil
}
