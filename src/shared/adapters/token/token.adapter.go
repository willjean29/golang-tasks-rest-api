package adapters

type TokenAdapter interface {
	GenerateToken(key string, value string) (string, error)
	ValidateToken(token string) (bool, map[string]interface{}, error)
}
