package plugins_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"app/src/config/plugins"
	error "app/src/shared/errors"
)

// MockEnvLoader es un mock para la interfaz EnvLoader
type MockEnvLoader struct {
	mock.Mock
}

// Load implementa el método Load de la interfaz EnvLoader para el mock
func (m *MockEnvLoader) Load(filename string) (plugins.Env, error.Error) {
	args := m.Called(filename)

	return args.Get(0).(plugins.Env), args.Get(1).(error.Error)
}

func TestNewEnv(t *testing.T) {
	tests := []struct {
		name          string
		expectedEnv   plugins.Env
		expectedError error.Error
	}{
		{
			name: "Success",
			expectedEnv: plugins.Env{
				TokenSecretKey: "jwt_secret_key_value",
				DBPort:         "db_port_value",
				DBHost:         "db_host_value",
				DBName:         "db_name_value",
				DBUser:         "db_user_value",
				DBPassword:     "db_password_value",
			},
			expectedError: error.Error{},
		},
		{
			name:        "Error",
			expectedEnv: plugins.Env{},
			expectedError: error.Error{
				StatusCode: 500,
				Message:    "Not load the environment",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// Crear un mock de EnvLoader
			mockLoader := new(MockEnvLoader)
			mockLoader.On("Load", ".env").Return(tc.expectedEnv, tc.expectedError)

			// Crear una instancia de Env utilizando el mock
			env, err := plugins.LoadEnv(mockLoader)
			if err.StatusCode != 0 {
				assert.Equal(t, tc.expectedEnv, env)
			}
			assert.Equal(t, tc.expectedEnv.TokenSecretKey, env.TokenSecretKey)
			assert.Equal(t, tc.expectedEnv.DBPort, env.DBPort)
			assert.Equal(t, tc.expectedEnv.DBHost, env.DBHost)
			assert.Equal(t, tc.expectedEnv.DBName, env.DBName)
			assert.Equal(t, tc.expectedEnv.DBUser, env.DBUser)
			assert.Equal(t, tc.expectedEnv.DBPassword, env.DBPassword)

		})
	}
}
