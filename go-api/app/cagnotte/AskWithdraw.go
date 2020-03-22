package cagnotte

type AskWithdraw struct {
	Id              string `json:"id"`
	WithdrawWallets []struct {
		Currency           string `json:"currency"`
		DestinationAddress string `json:"destinationAddress"`
	} `json:"withdrawWallets"`
}
