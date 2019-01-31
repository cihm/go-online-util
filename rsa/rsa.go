package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func RSAflow() {

	// generate key
	privatekey, publickey, err := GenerateKey()
	if err != nil {
		log.Fatalf("Cannot generate RSA key\n")
	}

	// // dump private key to file
	// err = DumpPrivateKeyFile(privatekey, "private.pem")
	// if err != nil {
	// 	log.Fatalf("Cannot dump private key file\n")
	// }
	// // dump public key to file
	// err = DumpPublicKeyFile(publickey, "public.pem")
	// if err != nil {
	// 	log.Fatalf("Cannot dump public key file\n")
	// }

	// load private key
	publickey, err = LoadPublicKeyFile("C:/opt/key/public.pem")
	if publickey == nil {
		fmt.Printf("Cannot load publickey key\n")
	}

	// encrypt message use public key
	message := "abcd"
	cipher, err := Encrypt(message, publickey)
	if err != nil {
		log.Fatalf("Cannot encrypt message\n")
	}
	fmt.Printf("encrypt result is (%s)\n", cipher)

	// load private key
	privatekey, err = LoadPrivateKeyFile("C:/opt/key/private.pem")
	if privatekey == nil {
		fmt.Printf("Cannot load private key\n")
	}

	// decrypt use private
	plain, err := Decrypt(cipher, privatekey)
	if err != nil {
		log.Fatalf("Cannot decrypt message\n")
	}
	fmt.Printf("decrypt result is (%s)\n", plain)

}

// decrypt
func Decrypt(ciphertext string, privatekey *rsa.PrivateKey) (string, error) {
	decodedtext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	sha256hash := sha256.New()
	decryptedtext, err := rsa.DecryptOAEP(sha256hash, rand.Reader, privatekey, decodedtext, nil)
	if err != nil {
		return "", fmt.Errorf("RSA decrypt failed, error=%s\n", err.Error())
	}

	return string(decryptedtext), nil
}

// encrypt
func Encrypt(plaintext string, publickey *rsa.PublicKey) (string, error) {
	label := []byte("")
	sha256hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(sha256hash, rand.Reader, publickey, []byte(plaintext), label)

	decodedtext := base64.StdEncoding.EncodeToString(ciphertext)
	return decodedtext, err
}

// Load private key from base64
func LoadPrivateKeyBase64(base64key string) (*rsa.PrivateKey, error) {
	keybytes, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	privatekey, err := x509.ParsePKCS1PrivateKey(keybytes)
	if err != nil {
		return nil, errors.New("parse private key error!")
	}

	return privatekey, nil
}

func LoadPublicKeyBase64(base64key string) (*rsa.PublicKey, error) {
	keybytes, err := base64.StdEncoding.DecodeString(base64key)
	if err != nil {
		return nil, fmt.Errorf("base64 decode failed, error=%s\n", err.Error())
	}

	pubkeyinterface, err := x509.ParsePKIXPublicKey(keybytes)
	if err != nil {
		return nil, err
	}

	publickey := pubkeyinterface.(*rsa.PublicKey)
	return publickey, nil
}

// Load private key from private key file
func LoadPrivateKeyFile(keyfile string) (*rsa.PrivateKey, error) {
	keybuffer, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(keybuffer))
	if block == nil {
		return nil, errors.New("private key error!")
	}

	privatekey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.New("parse private key error!")
	}

	return privatekey, nil
}

func LoadPublicKeyFile(keyfile string) (*rsa.PublicKey, error) {
	keybuffer, err := ioutil.ReadFile(keyfile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keybuffer)
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubkeyinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publickey := pubkeyinterface.(*rsa.PublicKey)
	return publickey, nil
}

// Dump private key to base64 string
// Compared with DumpPrivateKeyBuffer this output:
//  1. Have no header/tailer line
//  2. Key content is merged into one-line format
// The output is:
//  MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2y8mEdCRE8siiI7udpge......2QIDAQAB
func DumpPrivateKeyBase64(privatekey *rsa.PrivateKey) (string, error) {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)

	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}

func DumpPublicKeyBase64(publickey *rsa.PublicKey) (string, error) {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	keybase64 := base64.StdEncoding.EncodeToString(keybytes)
	return keybase64, nil
}

// Dump private key to buffer.
func DumpPrivateKeyBuffer(privatekey *rsa.PrivateKey) (string, error) {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keybytes,
	}

	var keybuffer []byte = pem.EncodeToMemory(block)
	return string(keybuffer), nil
}

func DumpPublicKeyBuffer(publickey *rsa.PublicKey) (string, error) {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return "", err
	}

	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keybytes,
	}

	var keybuffer []byte = pem.EncodeToMemory(block)
	return string(keybuffer), nil
}

// Dump private key into file
// This has same output as DumpPrivateKeyBuffer(), but dump to a file:
//  -----BEGIN RSA PRIVATE KEY-----
//  MIIEoQIBAAKCAQEAuql1lFYgKmKA1x5lQyadktbkeRRO0qrsmAkhvTtiz2p0Y+Ur
//  xWSYqDlmoY6vdkxj0Ex0z4zisoPnI+K89hV69O9v/83Yz7hYkLBHuwGiiSOiPZU7
//  ...
//  PfKnburLQLE50wPkglfnGYfqQxtIiqn1hGTQO1xBxu03g+KM/Q==
//  -----END RSA PRIVATE KEY-----
func DumpPrivateKeyFile(privatekey *rsa.PrivateKey, filename string) error {
	var keybytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keybytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// Dump public key into file
//  -----BEGIN PUBLIC KEY-----
//  MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2y8mEdCRE8siiI7udpge
//  5y1hrlSJzV7Xj0UojL/hi9u7s6TjYQQDA4M++/FezwkO5lBby2C+wK8bY7lgphuP
//  ...
//  OZPrh/jItinhdzhyIXuYn6ohesPlM9i5TMpeBfpBmCwQQTfsAjBnXTTQzT4m4cmo
//  2QIDAQAB
//  -----END PUBLIC KEY-----
func DumpPublicKeyFile(publickey *rsa.PublicKey, filename string) error {
	keybytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		return err
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: keybytes,
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

// Generate RSA private/public key
func GenerateKey() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	publickey := &privatekey.PublicKey
	return privatekey, publickey, nil
}
