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
		decoded, _ := base64.StdEncoding.DecodeString(strings.Replace(os.Args[1], "%3D", "=", -1))
		token := string(decoded)
		byteTime, _ := hex.DecodeString(token[:strings.Index(token, "_")])
		fmt.Println("Token expires:", time.Unix(int64(binary.BigEndian.Uint32(byteTime)), 0))
	} else {
		fmt.Println("Missing token.")
	}
}
