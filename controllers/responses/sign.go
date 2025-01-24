package responses

type SignResponse struct { //хочу вовзращать access and refresh token
	Id      string  `json:"id"`
	Session Session `json:"session"`
}
