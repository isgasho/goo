//salt.go This file is all about encrypting and decrypting passwords.
package goo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gitlab.com/NebulousLabs/fastrand"
	"golang.org/x/crypto/argon2"
	"time"
)

// return result and salt. Use hexadecimal to save []byte of data.
func EncryptArgon2(password string) (resultDB , saltDB string) {
	salt := fastrand.Bytes(32)
	key := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
	resultDB, saltDB = fmt.Sprintf("%x", key), fmt.Sprintf("%x", salt)
	return
}

//Verify that the user's password is equal to the password in the database
func DeEncryptArgon2(password , resultDB , saltDB string) (is bool, err error) {
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

// use md5+salt return result and salt.
func EncryptMd5(password string) (resultDB , saltDB string) {
	saltDB = fmt.Sprintf("%x", time.Now().UnixNano())
	resultByte := md5.Sum([]byte(password + saltDB))
	resultDB = fmt.Sprintf("%x", resultByte)
	return
}
//Verify that the user's password is equal to the password in the database
func DeEncryptMd5(password string,resultDB ,saltDB string)(is bool,err error){

	if fmt.Sprintf("%x",md5.Sum([]byte(password+saltDB))) == resultDB {
		return true ,nil
	}
	return false,fmt.Errorf("EdEncryptMd5 error")
}
