package json_prime

type IsPrimeRequest struct {
	Method string `json:"method"`
	Number int    `json:"number"`
}

type IsPrimeResponse struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}
