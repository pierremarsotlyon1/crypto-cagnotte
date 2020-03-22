package coinbase

type CoinbaseAddress struct {
	Data struct {
		Id      string `json:"id"`
		Address string `json:"address"`
	} `json:"data"`
}
