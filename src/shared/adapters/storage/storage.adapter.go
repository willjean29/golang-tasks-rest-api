package adapters

type StoreAdapter interface {
	SaveFile(filename string) (string, error)
	DeleteFile(filename string) error
}
