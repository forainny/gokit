package password

import "testing"

func TestEncrypt(t *testing.T) {
	password := "123456"
	hash := Encrypt(password)
	if hash[:3] != "$2a" {
		t.Fatal("encrypt failed", hash, hash[3])
	}
	t.Log(hash)
}

func TestIsValid(t *testing.T) {
	password := "123456"
	hash := Encrypt(password)
	if !IsValid(hash, password) {
		t.Fatal("password error")
	}
	t.Log("password is valid")
}
