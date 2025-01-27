package responses

type SignResponse struct {
	Id      string  `json:"id"`
	Session Session `json:"session"`
}
