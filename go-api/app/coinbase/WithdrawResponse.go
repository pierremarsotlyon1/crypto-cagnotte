package coinbase

type WithdrawResponse struct {
	Data struct {
		Id     string `json:"id"`
		Type   string `json:"type"`
		Status string `json:"status"`
		Amount struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"amount"`
		NativeAmount struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"native_amount"`
		Description  string `json:"description"`
		CreatedAt    string `json:"created_at"`
		updatedAd    string `json:"updated_at"`
		Resource     string `json:"transaction"`
		ResourcePath string `json:"resource_path"`
		Network      struct {
			Status string `json:"status"`
			Hash   string `json:"hash"`
			Name   string `json:"name"`
		} `json:"network"`
		To struct {
			Resource string `json:"resource"`
			Address  string `json:"address"`
		} `json:"to"`
		Details struct {
			Title    string `json:"title"`
			Subtitle string `json:"subtitle"`
		} `json:"details"`
	} `json:"data"`
}
