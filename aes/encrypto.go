package d2

import (
	"bytes"
	"crypto/aes"
	"errors"
)

// ECB模式
func EncryptAES(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()

	// padding
	pad := byte(bs - (len(plaintext) % bs))
	plaintext = append(plaintext, bytes.Repeat([]byte{pad}, int(pad))...)

	ciphertext := make([]byte, len(plaintext))
	ciphertext2 := ciphertext
	for len(plaintext) > 0 {
		block.Encrypt(ciphertext, plaintext)
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}
	return ciphertext2, nil
}

func DecryptAES(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	bs := block.BlockSize()
	plaintext := make([]byte, len(ciphertext))
	plaintext2 := plaintext
	for len(ciphertext) > 0 {
		block.Decrypt(plaintext, ciphertext)
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}

	pad := int(plaintext2[len(plaintext2) - 1])
	if pad > bs {
		return nil, errors.New("AES: 无效的 PKCS5 填充")
	}
	return plaintext2[:len(plaintext2) - pad], nil
}
