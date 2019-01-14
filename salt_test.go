package goo

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/argon2"
	"testing"
)

func TestEncryptArgon2(t *testing.T) {
	pass, salt := EncryptArgon2("00oop") // 数据库中的pass和salt

	passReal := "00oop" // 用户的输入
	s, _ := hex.DecodeString(salt)
	passR := argon2.IDKey([]byte(passReal), s, 1, 64*1024, 4, 32)
	if pass != fmt.Sprintf("%x", passR) {
		t.Error("错误的❌")
		fmt.Println("yao-----")
		fmt.Println(pass, fmt.Sprintf("%x", passR))
	}

}

//34577390 ns/op
func BenchmarkEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncryptArgon2("12345678910")
	}
}

//35149770 ns/op
func BenchmarkDeEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeEncryptArgon2("1234567890", "7eca9e914643d82fc76d3dbf0124ab5089284456435d6ddfac88457f3463b06e", "7eca9e914643d82fc76d3dbf0124ab5089284456435d6ddfac88457f3463b06e")
	}
}

func TestEncryptMd5(t *testing.T) {
	pass,salt := EncryptMd5("991182@")
	passReal := "991182@"
	a,b := DeEncryptMd5(passReal,pass,salt)
	if b != nil || !a {
		t.Error("❌")
	}
	fmt.Println(pass,salt)
	//passErr := "123456!@"
	//d,e := DeEncryptMd5(passErr,pass,salt)
	//if e != nil || !d {
	//	t.Error("这个错误 应该被打印出来")
	//}


}

func BenchmarkEncryptMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EncryptMd5("123456!!")
	}
}

func BenchmarkDeEncryptMd5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DeEncryptMd5("12345!!","9e914643","14643d82fc76d3dbf0124ab50892844")
	}
}
