package pkg

import (
	"fmt"
	"strings"
)

type bitgoClient struct {
	Token, URL string
}


func NewBitGoClient(token string, isProd bool) *bitgoClient {
	client := &bitgoClient{Token: token}
	url := "https://%s.bitgo.com/"
	if isProd {
		client.URL = fmt.Sprintf(url, "www")
	} else {
		client.URL = fmt.Sprintf(url, "test")
	}
	return client
}


func (c *bitgoClient) GetUserBy(ID string) (string, error) {
	url := fmt.Sprintf(c.URL + "api/v2/user/%s", ID)
	res, err := query("GET", url, c.Token, nil)
	if err != nil {
		return "", err
	}
	return res, nil
}


func (c *bitgoClient) NewWalletAddress(coin, walletID, payload string)  (string, error){
	if _, err := isJSON(payload); err != nil {
		return "", err
	}

	url := fmt.Sprintf(c.URL + "api/v2/%s/wallet/%s/address", coin, walletID)
	body := strings.NewReader(payload)
	res, err := query("POST", url, c.Token, body)
	if err != nil {
		return "", err
	}
	return res, nil
}


func (c *bitgoClient) CreateEnterprise(payload string) (string, error) {
	if _, err := isJSON(payload); err != nil {
		return "", err
	}

	url := c.URL + "api/v2/enterprise"
	body := strings.NewReader(payload)
	res, err := query("POST", url, c.Token, body)
	if err != nil {
		return "", err
	}

	return res, nil
}
