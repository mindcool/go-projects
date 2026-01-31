package intermediate

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"
	// Create a checksum for that password
	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))
	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash256)
	// Get hexadecimal value of our hash
	fmt.Printf("SHA-256 Hash hex value: %x\n", hash256)
	fmt.Println("Hash: ", hash512)
	// Get hexadecimal value of our hash
	fmt.Printf("SHA-512 Hash hex value: %x\n", hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}
	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in the database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt String:", saltStr)
	fmt.Println("Hash this one got password and salt: ", signUpHash)

	// Verify
	// Retrive the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to Decode Salt: ", err)
		return
	}
	loginHash := hashPassword(password, decodedSalt)

	// Compare the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct you are loged in")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash the password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
