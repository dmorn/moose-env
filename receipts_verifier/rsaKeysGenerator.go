package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Generate RSA Keys - Server
	serverPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	serverPublicKey := &serverPrivateKey.PublicKey

	serverPrivateKeyMarshalled := x509.MarshalPKCS1PrivateKey(serverPrivateKey)
	serverPublicKeyMarshalled, _ := x509.MarshalPKIXPublicKey(serverPublicKey)

	// Generate RSA Keys - Client
	clientPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}

	clientPublicKey := &clientPrivateKey.PublicKey

	clientPrivateKeyMarshalled := x509.MarshalPKCS1PrivateKey(clientPrivateKey)
	clientPublicKeyMarshalled, _ := x509.MarshalPKIXPublicKey(clientPublicKey)

	//save everything

	pwd, err := os.Getwd()
	check(err)

	//server
	path_s_pr := fmt.Sprintf("%s/.moose_s_pr_key", pwd)
	err = ioutil.WriteFile(path_s_pr, serverPrivateKeyMarshalled, 0644)
	check(err)

	path_s_pu := fmt.Sprintf("%s/.moose_s_pu_key", pwd)
	err = ioutil.WriteFile(path_s_pu, serverPublicKeyMarshalled, 0644)
	check(err)

	//client
	path_c_pr := fmt.Sprintf("%s/.moose_c_pr_key", pwd)
	err = ioutil.WriteFile(path_c_pr, clientPrivateKeyMarshalled, 0644)
	check(err)

	path_c_pu := fmt.Sprintf("%s/.moose_c_pu_key", pwd)
	err = ioutil.WriteFile(path_c_pu, clientPublicKeyMarshalled, 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
