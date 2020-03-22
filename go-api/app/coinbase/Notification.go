package coinbase

type Notification struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Data struct {
		Id           string `json:"id"`
		Address      string `json:"address"`
		Name         string `json:"name"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		Resource     string `json:"resource"`
		ResourcePath string `json:"resource_path"`
	} `json:"data"`
	User struct {
		Id           string `json:"id"`
		Resource     string `json:"resource"`
		ResourcePath string `json:"resource_path"`
	} `json:"user"`
	Account struct {
		Id           string `json:"id"`
		Resource     string `json:"resource"`
		ResourcePath string `json:"resource_path"`
	} `json:"account"`
	DeliveryAttempts int64  `json:"delivery_attempts"`
	CreatedAt        string `json:"created_at"`
	Resource         string `json:"resource"`
	ResourcePath     string `json:"resource_path"`
	AdditionalData   struct {
		Hash   string `json:"hash"`
		Amount struct {
			Amount   string `json:"amount"`
			Currency string `json:"currency"`
		} `json:"amount"`
		Transaction struct {
			Id           string `json:"id"`
			Resource     string `json:"resource"`
			ResourcePath string `json:"resource_path"`
		} `json:"transaction"`
	} `json:"additional_data"`
}
