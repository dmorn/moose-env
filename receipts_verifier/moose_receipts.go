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

func main() {

	//read everything
	pwd, err := os.Getwd()
	check(err)

	path_s_pu := fmt.Sprintf("%s/moose_s_pu_key", pwd)
	serverPublicKeyMarshalled, err := ioutil.ReadFile(path_s_pu)
	check(err)

	//client
	path_c_pr := fmt.Sprintf("%s/moose_c_pr_key", pwd)
	clientPrivateKeyMarshalled, err := ioutil.ReadFile(path_c_pr)
	check(err)

	//get Actual keys
	serverPublicKeyI, err := x509.ParsePKIXPublicKey(serverPublicKeyMarshalled)
	var serverPublickey *rsa.PublicKey = serverPublicKeyI.(*rsa.PublicKey)

	clientPrivateKey, err := x509.ParsePKCS1PrivateKey(clientPrivateKeyMarshalled)

	label := []byte("")
	hash := sha256.New()

	// Decrypt Message
	// then config file settings
	source, err := os.Open("source.json")
	check(err)

	var receipt Receipt
	jsonParser := json.NewDecoder(source)
	err = jsonParser.Decode(&receipt)
	check(err)

	//fmt.Println(string(receipt.Data))
	//data is json -> hex
	var jsonFromHexData []byte

	jsonFromHexData, err = hex.DecodeString(string(receipt.Data))
	check(err)

	//decode jsonFromHexData
	plainTextJsonData, err := rsa.DecryptOAEP(hash, rand.Reader, clientPrivateKey, jsonFromHexData, label)
	check(err)

	//decode json into ItemShorter that is a json as well
	var item ItemShorter
	err = json.Unmarshal(plainTextJsonData, &item)
	check(err)

	//fmt.Printf("OAEP decrypted %s to \n%s\n", receipt.Data, item)

	fmt.
		Printf("Receipt: \n\n \nItemID:\t\t\t%d \nName:\t\t\t%s \nCoins:\t\t\t%d \nQuantity:\t\t%d \nStockID:\t\t%d \nObjectID:\t\t%d\n\n",
		item.ID, item.Name, item.Coins, item.Quantity, item.StockId, item.ObjectId)

	//Verify Signature
	//Message - Signature
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto
	PSSmessage := plainTextJsonData
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	//signature from hex
	signature, err := hex.DecodeString(receipt.Signature)
	check(err)

	err = rsa.VerifyPSS(serverPublickey, newhash, hashed, signature, &opts)

	if err != nil {
		fmt.Println("Who are U? Verify Signature failed")
		os.Exit(1)
	} else {
		fmt.Println("Verify Signature successful")
	}

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
