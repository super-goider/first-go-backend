package users

type User struct {
	ID           int
	Login        string
	Email        string
	PasswordHash string
}

type UserCreateRequest struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// это DTO (Data Transfer Object) - структура, которая существует только для входящих или исходящих данных API

type UserLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserResponse struct { // типа dto для исходящих данных чтобы не отдавать хеш пароля
	ID    int    `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}

func ToUserResponse(u User) UserResponse {
	return UserResponse{
		ID:    u.ID,
		Login: u.Login,
		Email: u.Email,
	}
}
