package intermediate

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo Base64 Enconding")
	// Encode to base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("We encounter an error while decoding: ", decoded)
		return
	}
	fmt.Println("Decoded String: ", string(decoded))

	// URL Safe avoid '/' and '+
	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Url Safe Encoded", urlSafeEncoded)
}
