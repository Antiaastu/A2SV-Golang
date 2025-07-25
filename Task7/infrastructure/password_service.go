package infrastructure

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) (string, error){
	hashed, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePassword(hash, pw string) error{
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}