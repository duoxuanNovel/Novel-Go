package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type HashUtil struct {
	Salt string
}

func NewHashUtil() *HashUtil {
	return &HashUtil{
		Salt: generateSalt(),
	}
}

func (h *HashUtil) EncryptPassword(password string) string {
	// Hash the user password
	hashedPassword := h.md5Hash(h.md5Hash(password) + h.Salt)
	return hashedPassword
}

func (h *HashUtil) CheckPassword(inputPassword, storedPassword string) bool {
	// Hash the input password
	hashedPassword := h.md5Hash(h.md5Hash(inputPassword) + h.Salt)

	// Compare the hashed passwords
	return hashedPassword == storedPassword
}

// md5Hash function generates an MD5 hash of a string
func (h *HashUtil) md5Hash(text string) string {
	hasher := md5.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

// generateSalt function generates a random salt
func generateSalt() string {
	rand.Seed(time.Now().UnixNano())
	uuid := md5.New()
	_, err := uuid.Write([]byte(fmt.Sprint(time.Now().UnixNano(), rand.Int63())))
	if err != nil {
		panic(err)
	}
	salt := hex.EncodeToString(uuid.Sum(nil))
	return salt[len(salt)-16:]
}
