package providers

type TokenProvider interface {
	GenerateToken(key string, value string) (string, error)
	ValidateToken(token string) (bool, error)
}
