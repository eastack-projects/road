package utils

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/argon2"
)

type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func GeneratePassword(c *PasswordConfig, password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, c.time, c.memory, c.threads, c.keyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d,$m=%d,t=%d,p=%d$%s$%s"

	full := fmt.Sprintf(format, argon2.Version, c.memory, c.time, c.threads, b64Salt, b64Hash)

	return full, nil
}

func ComparePassword(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")

	c:=&PasswordConfig{}
	_, err := fmt.Sscanf(part[3], "m=%d,t=%d,p=%d", &c.memory, &c.time, &c.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(part[4])
}
