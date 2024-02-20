package plugins

import (
	error "app/src/shared/errors"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	TokenSecretKey string
	DBPort         string
	DBHost         string
	DBName         string
	DBUser         string
	DBPassword     string
}
type EnvLoader interface {
	Load(string) (Env, error.Error)
}

func LoadEnv(env EnvLoader, filename string) (Env, error.Error) {
	envs, err := env.Load(filename)
	if err.StatusCode != 0 {
		return Env{}, err
	}
	return envs, error.Error{}
}

// EnvLoaderGodotenv es una implementación de EnvLoader que carga las variables de entorno usando godotenv.
type EnvLoaderGodotenv struct{}

// Load carga las variables de entorno desde el archivo especificado utilizando godotenv.
func (l *EnvLoaderGodotenv) Load(filename string) (Env, error.Error) {
	err := godotenv.Load(filename)
	if err != nil {
		// Si hubo un error al cargar el archivo .env, devolvemos un error.
		return Env{}, *error.New("Error loading .env file", 500, err)
	}

	// Si no hubo errores, devolvemos las variables de entorno cargadas.
	return Env{
		TokenSecretKey: os.Getenv("TOKEN_SECRET_KEY"),
		DBPort:         os.Getenv("DB_PORT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBName:         os.Getenv("DB_NAME"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
	}, error.Error{}
}

var Envs, _ = LoadEnv(&EnvLoaderGodotenv{}, ".env")
