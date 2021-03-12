package encryption

import (
	"github.com/shon-phand/bookstore_users-api/domains/errors"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) ([]byte, *errors.RestErr) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, errors.StatusInternalServerError("error in encrypting password")
	}
	return hashedPassword, nil
}
