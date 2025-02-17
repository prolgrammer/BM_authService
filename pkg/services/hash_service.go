package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	e "github.com/prolgrammer/BM_package/errors"
	"golang.org/x/crypto/argon2"
	"strings"
)

type hashService struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type HashService interface {
	CreateHash(password string) ([]byte, error)
	CompareStringAndHash(stringToCompare string, hashedString string) (bool, error)
}

func NewHashService() HashService {
	return &hashService{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 1,
		saltLength:  16,
		keyLength:   32,
	}
}

func (h *hashService) CreateHash(password string) ([]byte, error) {
	salt, err := generateRandomBytes(h.saltLength)
	if err != nil {
		return nil, err
	}

	hash := argon2.IDKey([]byte(password), salt, h.iterations, h.memory, h.parallelism, h.keyLength)
	b64salt := base64.RawStdEncoding.EncodeToString(salt)
	b64hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := []byte(fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", VersionKey, h.memory, h.iterations, h.parallelism, b64salt, b64hash))

	return encodedHash, nil
}

func generateRandomBytes(length uint32) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil

}

func (h *hashService) CompareStringAndHash(stringToCompare string, hashedString string) (bool, error) {
	h, salt, hash, err := decodeHash(hashedString)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(stringToCompare), salt, h.iterations, h.memory, h.parallelism, h.keyLength)
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}

	return false, e.ErrPasswordMismatch
}

func decodeHash(encodedHash string) (p *hashService, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, fmt.Errorf("the encoded hash is not in the correct format")
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != VersionKey {
		return nil, nil, nil, fmt.Errorf("incompatible version of argon2")
	}

	p = &hashService{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
