package datasource

import (
	"app/src/modules/users/domain/entities"
	error "app/src/shared/errors"
	"encoding/json"
	"log"
	"os"
)

var folder = "db"
var collection = "users"
var path = folder + "/" + collection + ".json"

type FileSistemUserDatasource struct{}

func (f *FileSistemUserDatasource) FindAll() (entities.ListUsers, error.Error) {

	file, err := os.Open(path)
	if err != nil {
		return entities.ListUsers{}, *error.New("Error to find users", 500, err)
	}

	defer file.Close()

	var users entities.ListUsers

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&users); err != nil {
		return entities.ListUsers{}, *error.New("Error to find users", 500, err)
	}

	return users, error.Error{}
}

func (f *FileSistemUserDatasource) FindById(id int) (entities.User, error.Error) {

	file, err := os.Open(path)
	if err != nil {
		return entities.User{}, *error.New("Error to find user", 500, err)
	}

	defer file.Close()

	var users entities.ListUsers

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&users); err != nil {
		return entities.User{}, *error.New("Error to find user", 500, err)
	}

	for _, user := range users {
		if user.ID == uint(id) {
			return user, error.Error{}
		}
	}

	return entities.User{}, *error.New("User not found", 404, nil)
}

func (f *FileSistemUserDatasource) FindByEmail(email string) (entities.User, error.Error) {

	file, err := os.Open(path)
	if err != nil {
		return entities.User{}, *error.New("Error to find user", 500, err)
	}

	defer file.Close()

	var users entities.ListUsers

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&users); err != nil {
		return entities.User{}, *error.New("Error to find user", 500, err)
	}

	for _, user := range users {
		if user.Email == email {
			return user, error.Error{}
		}
	}

	return entities.User{}, *error.New("User not found", 404, nil)
}

func (f *FileSistemUserDatasource) Create(createUser entities.CreateUser) (entities.User, error.Error) {
	log.Println("Create user")
	var users entities.ListUsers
	var userId uint
	// Crea el directorio "db" si no existe
	err := os.MkdirAll("db", os.ModePerm)
	if err != nil {
		log.Println("Error al crear el directorio:", err)
		return entities.User{}, *error.New("Error to create user", 500, err)
	}

	// Abre o crea el archivo "users.json" en modo lectura y escritura
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error al abrir o crear el archivo:", err)
		return entities.User{}, *error.New("Error to create user", 500, err)
	}
	defer file.Close()

	// Obtiene información sobre el archivo
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Error al obtener información del archivo:", err)
		return entities.User{}, *error.New("Error to create user", 500, err)
	}

	// Verifica si el archivo está vacío
	if fileInfo.Size() == 0 {
		log.Println("El archivo 'users.json' está vacío. Creando nuevo usuario...")
		// Establece el id
		userId = uint(1)
	} else {
		// Si el archivo contiene datos, decodifica los usuarios existentes
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&users); err != nil {
			log.Println("Error al decodificar los usuarios del archivo:", err)
			return entities.User{}, *error.New("Error to create user", 500, err)
		}
		// Calcula el ID del nuevo usuario
		userId = uint(len(users) + 1)
	}

	// Crea un nuevo usuario y lo guarda en el archivo
	newUser := entities.User{
		ID:       userId,
		Name:     createUser.Name,
		Email:    createUser.Email,
		Password: createUser.Password,
	}

	users = append(users, newUser)

	// Vuelve al inicio del archivo para sobrescribirlo con los usuarios actualizados
	if _, err := file.Seek(0, 0); err != nil {
		log.Println("Error al volver al inicio del archivo:", err)
		return entities.User{}, *error.New("Error to create user", 500, err)
	}

	// Escribe los usuarios actualizados en el archivo
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(users); err != nil {
		log.Println("Error al escribir los usuarios actualizados en el archivo:", err)
		return entities.User{}, *error.New("Error to create user", 500, err)
	}

	log.Println("Nuevo usuario creado exitosamente.")
	return newUser, error.Error{}
}

func (f *FileSistemUserDatasource) Save(user entities.User) error.Error {

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return *error.New("Error to save user", 500, err)
	}

	defer file.Close()

	var users entities.ListUsers

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&users); err != nil {
		return *error.New("Error to save user", 500, err)
	}

	for i, u := range users {
		if u.ID == user.ID {
			users[i] = user
		}
	}

	encoder := json.NewEncoder(file)

	if err := encoder.Encode(users); err != nil {
		return *error.New("Error to save user", 500, err)
	}

	return error.Error{}
}
