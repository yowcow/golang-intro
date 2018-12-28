package myprivate

type Secret struct {
	message string
}

func NewSecret(msg string) *Secret {
	return &Secret{message: msg}
}

func (s *Secret) getMessage() string {
	return s.message
}
