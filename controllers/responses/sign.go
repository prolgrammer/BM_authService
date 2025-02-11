package responses

type SignResponse struct {
	Id      string  `json:"id" example:"1"`
	Session Session `json:"session"`
}
