package cagnotte

import (
	"crypto-cagnotte/go-api/app/coinbase"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cagnotte struct {
	ID          primitive.ObjectID          `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string                      `json:"name" bson:"name"`
	Description string                      `json:"description" bson:"description"`
	Wallets     []Wallet                    `json:"wallets" bson:"wallets"`
	Creator     primitive.ObjectID          `json:"creator" bson:"creator"`
	Days        int16                       `json:"days" bson:"days"`
	Status      int8                        `json:"status" bson:"status"`
	Withdraws   []coinbase.WithdrawResponse `json:"withdraws" bson:"withdraws"`
}

func (cagnotte *Cagnotte) IsValid() bool {
	return len(cagnotte.Name) > 0 && len(cagnotte.Description) > 0 && cagnotte.Days > 0
}
