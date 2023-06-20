package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
	"os"

	"github.com/denisbrodbeck/machineid"
)

const VERSION = 100
const ENCRYPTION_KEY = "4528482B4D6251655468576D5A7134743777217A25432A462D4A404E635266556A586E3272357538782F413F4428472B4B6250645367566B5970337336763979"
const KEY_SIZE = 32
const DISCORD_ENDPOINT = "https://canary.discord.com/api/v9/users/@me"

func IsDigit(str string) bool {
	for _, char := range str {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func GetSystemHWID() (string, error) {
	hwid, err := machineid.ProtectedID(ENCRYPTION_KEY)
	if err != nil {
		return "", err
	}
	return hwid, nil
}

func Encrypt(plaintext *string) (*string, error) {
	def := ""
	key, _ := hex.DecodeString(ENCRYPTION_KEY)
	keyBytes := sha512.Sum512(key)
	block, err := aes.NewCipher(keyBytes[:KEY_SIZE])
	if err != nil {
		return &def, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(*plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return &def, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(*plaintext))

	res := hex.EncodeToString(ciphertext)
	return &res, nil
}

func Decrypt(ciphertext *string) (*string, error) {
	def := ""
	key, _ := hex.DecodeString(ENCRYPTION_KEY)
	keyBytes := sha512.Sum512(key)
	block, err := aes.NewCipher(keyBytes[:KEY_SIZE])
	if err != nil {
		return &def, err
	}

	ciphertextBytes, err := hex.DecodeString(*ciphertext)
	if err != nil {
		return &def, err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return &def, errors.New("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertextBytes, ciphertextBytes)

	res := string(ciphertextBytes)
	return &res, nil
}

func SplitStringListEqual(list *[]string, parts uint8) [][]string {
	var divided [][]string
	l := *list
	var partSize int = len(l) / int(parts)
	for i := 0; i < int(parts); i++ {
		if i == int(parts)-1 {
			divided = append(divided, l[i*partSize:])
		} else {
			divided = append(divided, l[i*partSize:(i+1)*partSize])
		}
	}
	return divided
}

func DeleteFileByPath(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
