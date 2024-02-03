package providers

type StoreProvider interface {
	SaveFile(filename string) (string, error)
	DeleteFile(filename string) error
}
