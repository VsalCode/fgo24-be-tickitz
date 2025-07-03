package dto;

type UpdatedUser struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}