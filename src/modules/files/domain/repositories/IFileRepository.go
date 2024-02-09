package repositories

type IFileRepository interface {
	Save() (string, error)
	Delete() error
}
