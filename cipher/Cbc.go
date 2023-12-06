//CBC describes a block cipher mode. In cryptography, a block cipher mode of operation is an algorithm that uses a
//block cipher to provide an information service such as confidentiality or authenticity. A block cipher by itself
//is only suitable for the secure cryptographic transformation (encryption or decryption) of one fixed-length group of
//bits called a block. A mode of operation describes how to repeatedly apply a cipher's single-block operation to
//securely transform amounts of data larger than a block.
//
//This package simplifies the usage of AES-256-CBC.

package cipher

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

// DecryptCbc is a function that decrypts a given cipher text with a provided key and initialization vector(iv).
func DecryptCbc(iv, key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(ciphertext) < blockSize {
		return nil, fmt.Errorf("ciphertext is shorter then block size: %d / %d", len(ciphertext), blockSize)
	}
	if iv == nil {
		iv = ciphertext[:blockSize]
		ciphertext = ciphertext[blockSize:]
	}
	plaintext := make([]byte, len(ciphertext))
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertext, plaintext)
	length := len(plaintext)
	padLen := int(plaintext[length-1])
	if padLen > length {
		return nil, fmt.Errorf("padding is greater then the length: %d / %d", padLen, length)
	}
	return plaintext[:(length - padLen)], nil
}

// EncryptCbc is a function that encrypts plaintext with a given key and an optional initialization vector(iv).
func EncryptCbc(iv, key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(plaintext)%blockSize != 0 {
		return nil, fmt.Errorf("plaintext is not a multiple of the block size: %d / %d", len(plaintext), blockSize)
	}
	padSize := blockSize - len(plaintext)%blockSize
	plaintext = append(plaintext, bytes.Repeat([]byte{byte(padSize)}, padSize)...)
	var ciphertext []byte
	if iv == nil {
		ciphertext = make([]byte, aes.BlockSize+len(plaintext))
		iv := ciphertext[:aes.BlockSize]
		if _, err := rand.Read(iv); err != nil {
			return nil, err
		}
		cbc := cipher.NewCBCEncrypter(block, iv)
		cbc.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	} else {
		ciphertext = make([]byte, len(plaintext))
		cbc := cipher.NewCBCEncrypter(block, iv)
		cbc.CryptBlocks(ciphertext, plaintext)
	}
	return ciphertext, nil
}
