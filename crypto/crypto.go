package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	apiUrl  = "https://api.binance.com"
	dataUrl = "https://data.binance.com"
	fapiUrl = "https://fapi.binance.com"
)

type Crypto struct {
	apiKey    string
	secretKey string
}

func NewCrypto(api, secret string) *Crypto {
	return &Crypto{
		apiKey: api,
		secretKey: secret,
	}
}

func (t *Crypto) Ping() bool {
	if _, err := http.Get(dataUrl + "/api/v3/ping"); err != nil {
		return false
	}

	return true
}

func (t *Crypto) Price(name ...string) (prices map[string]string) {
	if len(name) == 0 {
		return 
	}
	symbols := fmt.Sprintf(`/api/v3/ticker/price?symbols=["%s"]`, strings.Join(name, `","`))
	res, err := http.NewRequest(http.MethodGet, dataUrl+symbols, nil)
	if err != nil {
		return 
	}

	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	
	price := []priceResp{}
	if err = json.Unmarshal(resBody, &price); err != nil {
		return
	}

	prices = make(map[string]string)
	for _, v := range price {
		prices[v.Symbol] = v.Price
	}

	return 
}

func (t *Crypto) FuturesPrice(name string) (prices string) {
	if len(name) == 0 {
		return 
	}
	symbols := fmt.Sprintf(`/fapi/v1/ticker/price?symbol=%s`, name)
	res, err := http.NewRequest(http.MethodGet, fapiUrl+symbols, nil)
	if err != nil {
		return 
	}

	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	
	price := priceResp{}
	if err = json.Unmarshal(resBody, &price); err != nil {
		return
	}
	

	return price.Price 
}

