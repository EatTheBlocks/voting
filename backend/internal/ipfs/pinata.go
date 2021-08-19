package ipfs

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type pinata struct {
	apiKey    string
	secretKey string
}

func NewPinataClient(apiKey, secretKey string) *pinata {
	return &pinata{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

func (p pinata) PinJSON(name string, body interface{}) (string, error) {
	var pinataBody struct {
		PinataOptions struct {
			CidVersion int `json:"cid_version"`
		} `json:"pinataOptions"`
		PinataMetadata struct {
			Name string `json:"name"`
		} `json:"pinataMetadata"`
		PinataContent interface{} `json:"pinataContent"`
	}

	pinataBody.PinataOptions.CidVersion = 0
	pinataBody.PinataMetadata.Name = name
	pinataBody.PinataContent = body

	jsonBody, err := json.Marshal(pinataBody)
	if err != nil {
		return "", errors.Wrap(err, "json encode")
	}

	req, err := http.NewRequest("POST", "https://api.pinata.cloud/pinning/pinJSONToIPFS", bytes.NewReader(jsonBody))
	if err != nil {
		return "", errors.Wrap(err, "new request")
	}

	req.Header.Set("pinata_api_key", p.apiKey)
	req.Header.Set("pinata_secret_api_key", p.secretKey)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "send request")
	}

	defer resp.Body.Close()

	var result struct {
		IpfsHash  string `json:"IpfsHash"`
		PinSize   int    `json:"PinSize"`
		Timestamp string `json:"Timestamp"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(&result); err != nil {
		return "", errors.Wrap(err, "json decode")
	}

	return result.IpfsHash, nil
}
