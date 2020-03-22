package cagnotte

type Wallet struct {
	Address         string  `json:"address" bson:"address"`
	Currency        string  `json:"currency" bson:"currency"`
	Amount          float64 `json:"amount" bson:"amount"`
	AvailableAmount float64 `json:"availableAmount" bson:"availableAmount"`
}
