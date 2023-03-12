package shared

import "golang.org/x/crypto/bcrypt"

func HashPassword(password []byte) string {

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func IsSamePassword(hashedPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)

	return err == nil

}
