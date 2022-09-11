package hasher

import (
	"crypto/sha256"
	"fmt"
)

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hasher.Sum(nil)
	return fmt.Sprintf("%x", hashedPassword)
}
