package requests

type SignRequest struct {
	Email    string `json:"email" example:"email@mail.ru"`
	Password string `json:"password" example:"password"`
}
