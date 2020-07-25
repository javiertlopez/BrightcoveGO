package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"
)

func main() {
	var out string

	flag.StringVar(&out, "output-dir", "", "Output directory to write files into")
	flag.Parse()

	if out == "" {
		out = "rsa-key_" + strconv.FormatInt(time.Now().Unix(), 10)
	}

	if err := os.MkdirAll(out, os.ModePerm); err != nil {
		panic(err.Error())
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err.Error())
	}

	privBytes := x509.MarshalPKCS1PrivateKey(priv)

	pubBytes, err := x509.MarshalPKIXPublicKey(priv.Public())
	if err != nil {
		panic(err.Error())
	}

	privOut, err := os.OpenFile(path.Join(out, "private.pem"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}

	if err := pem.Encode(privOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: privBytes}); err != nil {
		panic(err.Error())
	}

	pubOut, err := os.OpenFile(path.Join(out, "public.pem"), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err.Error())
	}

	if err := pem.Encode(pubOut, &pem.Block{Type: "PUBLIC KEY", Bytes: pubBytes}); err != nil {
		panic(err.Error())
	}

	var pubEnc = base64.StdEncoding.EncodeToString(pubBytes)

	var pubEncOut = path.Join(out, "public_key.txt")
	if err := ioutil.WriteFile(pubEncOut, []byte(pubEnc+"\n"), 0600); err != nil {
		panic(err.Error())
	}

	fmt.Println("Public key saved in " + pubEncOut)
}
