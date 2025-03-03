package dto

type UserRegisterDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatar_url"`
}

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInfoDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	PhotoUrl  string `json:"photo_url"`
}

type UserLoginResponseDTO struct {
	User  UserInfoDTO `json:"user"`
	Token string      `json:"token,omitempty"`
}
