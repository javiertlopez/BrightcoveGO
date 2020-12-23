package main

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		token, err := decode(os.Args[1])
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		fmt.Printf("Token expires: %s", token)
	} else {
		fmt.Println("Missing token.")
	}
}

// decode a Fastly Token. Returns expiration or error.
func decode(token string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(strings.Replace(token, "%3D", "=", -1))
	if err != nil {
		return "", err
	}

	decToken := string(decoded)
	byteTime, err := hex.DecodeString(decToken[:strings.Index(decToken, "_")])
	if err != nil {
		return "", err
	}

	return time.Unix(int64(binary.BigEndian.Uint32(byteTime)), 0).UTC().String(), nil
}
