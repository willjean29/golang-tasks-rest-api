package adapters

import "golang.org/x/crypto/bcrypt"

type BcryptAdapter struct{}

func (b *BcryptAdapter) ComparePasswords(hashedPwd string, plainPwd string) error {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err
}

func (b *BcryptAdapter) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
