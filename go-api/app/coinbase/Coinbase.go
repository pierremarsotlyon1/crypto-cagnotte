package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const API_ENDPOINT = "https://api.coinbase.com/"
const API_COINBASE_VERSION = "v2"
const API_KEY = "5NfXuQ60G4otrGQj"
const API_SECRET = "2sC5yLeObVry1TZfYsEdkwTPYZvX9G4U"
const USDCAccount = "5be9d0134f06a1033d4073fb"
const DAIAccount = "5c1b334f891a0c0190d1e0a5"

func GetUSDCAddress() *CoinbaseAddress {
	return getAddress("/accounts/" + USDCAccount + "/addresses")
}

func GetDAIAddress() *CoinbaseAddress {
	return getAddress("/accounts/" + DAIAccount + "/addresses")
}

func getAddress(pathUrl string) *CoinbaseAddress {
	req, err := http.NewRequest("POST", API_ENDPOINT+"/"+API_COINBASE_VERSION+pathUrl, nil)

	if err != nil {
		return nil
	}

	addHeaders(req, "", pathUrl, "POST")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	coinbaseAddress := new(CoinbaseAddress)
	json.Unmarshal(body, coinbaseAddress)
	return coinbaseAddress
}

func addHeaders(req *http.Request, bodyJson string, url string, method string) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	message := timestamp + method + ("/" + API_COINBASE_VERSION + url) + bodyJson
	hmac := hmac.New(sha256.New, []byte(API_SECRET))
	hmac.Write([]byte(message))
	sha := hex.EncodeToString(hmac.Sum(nil))

	req.Header.Add("CB-ACCESS-SIGN", sha)
	req.Header.Add("CB-ACCESS-TIMESTAMP", timestamp)
	req.Header.Add("CB-ACCESS-KEY", API_KEY)
	req.Header.Add("Content-Type", "application/json")
}
