package goo

import (
	"encoding/hex"
	"fmt"
	"gitlab.com/NebulousLabs/fastrand"
	"golang.org/x/crypto/argon2"
)

// return result and salt. Use hexadecimal to save []byte of data.
func Encrypt(password string) (result string, salt string) {
	salt := fastrand.Bytes(32)
	key := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	result, salt = fmt.Sprintf("%x", key), fmt.Sprintf("%x", salt)
	return
}

//Verify that the user's password is equal to the password in the database
func DeEncrypt(password string, resultDB string, saltDB string) (is bool, err error) {
	saltDBByte, err := hex.DecodeString(saltDB)
	if err != nil {
		return false, err
	}
	reslutByte := argon2.IDKey([]byte(password), saltDBByte, 1, 64*1024, 4, 32)
	if resultDB == fmt.Sprintf("%x", reslutByte) {
		return true, nil
	}
	return false, nil
}

