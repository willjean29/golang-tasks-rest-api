package providers

type StoreProvider interface {
	SaveFile(hashedPwd string, plainPwd string) error
	DeleteFile(password string) (string, error)
}
