package cagnotte

type AddCagnotte struct {
	Cagnotte
	UseUSDCWallet bool `json:"useUsdcWallet"`
	UseDAIWallet  bool `json:"useDaiWallet"`
}
