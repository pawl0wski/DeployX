package modules

import "crypto/sha256"

func HashPassword(password string) string {
	hasher := sha256.New()
	hashedPassword := hasher.Sum([]byte(password))
	return string(hashedPassword)
}
