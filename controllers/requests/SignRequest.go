package requests

type SignRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
