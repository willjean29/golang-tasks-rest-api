package adapters

type HashAdapter interface {
	ComparePasswords(hashedPwd string, plainPwd string) error
	HashPassword(password string) (string, error)
}
