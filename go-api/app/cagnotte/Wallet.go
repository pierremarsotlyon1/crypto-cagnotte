package cagnotte

type Wallet struct {
	Address  string `json:"address" bson:"address"`
	Currency string `json:"currency" bson:"currency"`
}
