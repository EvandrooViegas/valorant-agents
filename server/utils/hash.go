package utils 

import (
	"crypto/sha256"
	"fmt"
)

func CreateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashValue := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashValue)
	return hashString
}

func CompareHashes(inputData, storedHash string) bool {
	// Generate hash for the input data
	inputHash := CreateHash(inputData)
	// Compare the generated hash with the stored hash
	return inputHash == storedHash
}