package utils

import (
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"math/rand"
	"strings"
)

func NewPasswordEncoder() PasswordEncoder {
	return PasswordEncoder{
		time: 1,
		memory: 64 * 1024,
		threads: 4,
		keyLen: 32,
	}
}

type PasswordEncoder struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func (encoder PasswordEncoder) Encode(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, encoder.time, encoder.memory, encoder.threads, encoder.keyLen)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d,$m=%d,t=%d,p=%d$%s$%s"

	full := fmt.Sprintf(format, argon2.Version, encoder.memory, encoder.time, encoder.threads, b64Salt, b64Hash)

	return full, nil
}

func (encoder PasswordEncoder) Matches(password, hash string) (bool, error) {
	parts := strings.Split(hash, "$")

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &encoder.memory, &encoder.time, &encoder.threads)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	decodeHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	encoder.keyLen = uint32(len(decodeHash))

	comparisonHash := argon2.IDKey([]byte(password), salt, encoder.time, encoder.memory, encoder.threads, encoder.keyLen)

	return (subtle.ConstantTimeCompare(decodeHash, comparisonHash) == 1), nil
}
