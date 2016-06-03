package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ReceiptForItem(item *Item) (*Receipt, error) {

	//read everything
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	//server
	path_s_pr := fmt.Sprintf("%s/.moose_s_pr_key", pwd)
	serverPrivateKeyMarshalled, err := ioutil.ReadFile(path_s_pr)
	if err != nil {
		return nil, err
	}

	path_c_pu := fmt.Sprintf("%s/.moose_c_pu_key", pwd)
	clientPublicKeyMarshalled, err := ioutil.ReadFile(path_c_pu)
	if err != nil {
		return nil, err
	}

	//get Actual keys
	serverPrivateKey, err := x509.ParsePKCS1PrivateKey(serverPrivateKeyMarshalled)

	clientPublicKeyI, err := x509.ParsePKIXPublicKey(clientPublicKeyMarshalled)
	var clientPublickey *rsa.PublicKey = clientPublicKeyI.(*rsa.PublicKey)

	//Encrypt  Message
	data := ItemShorter{item.Id, item.Name, item.Coins, item.Quantity, item.StockId, item.ObjectId}

	itemMarshalled, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	label := []byte("")
	hash := sha256.New()

	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, clientPublickey, itemMarshalled, label)
	if err != nil {
		return nil, err
	}

	// Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	PSSmessage := itemMarshalled
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, serverPrivateKey, newhash, hashed, &opts)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("PSS Signature : %x\n", signature)
	return &Receipt{hex.EncodeToString(ciphertext), hex.EncodeToString(signature)}, nil
}

type Receipt struct {
	Data      string `json:"data"`
	Signature string `json:"signature"`
}

type ItemShorter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Coins    int    `json:"coins_payed"`
	Quantity int    `json:"quantity"`
	StockId  int    `json:"stock_id"`
	ObjectId int    `json:"object_id"`
}
