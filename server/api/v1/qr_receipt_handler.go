package main

import (
	"encoding/json"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
)

func QRImageFromReceipt(receipt *Receipt) (image.Image, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://api.qrserver.com/v1/create-qr-code/", nil)

	//add params
	receiptString, err := json.Marshal(receipt)

	q := req.URL.Query()
	q.Add("size", "500x500")
	q.Add("data", fmt.Sprintf("%s", receiptString))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	image, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return image, nil
}
