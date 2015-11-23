package d2

import (
	"bytes"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	key := []byte("22g4kg8fljy1s80654732n79kdo2qj7e")
	plaintext := [][]byte{
		[]byte(""),
		[]byte("hello worldaaaaaaaaaaaaaa"),
		[]byte("aaaaaaaa"),
		[]byte("中国人确认s 是@#@……#*@！&（￥）@#￥（安静嗲话多覅喝水费ISA和安全卫士24sdjasiod "),
	}

	for _, expected := range plaintext {
		ciphertext, err := EncryptAES(expected, key)
		if err != nil {
			t.Fatal("call EncryptAES failed")
		}
		got, err := DecryptAES(ciphertext, key)
		if err != nil {
			t.Fatal("call DecryptAES failed")
		}
		if bytes.Compare(got, expected) != 0 {
			t.Errorf("Expected %s, got %s", expected, got)
		}
	}
}
