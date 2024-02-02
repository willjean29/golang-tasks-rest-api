package providers

type HashProvider interface {
	ComparePasswords(hashedPwd string, plainPwd string) error
	HashPassword(password string) (string, error)
}
