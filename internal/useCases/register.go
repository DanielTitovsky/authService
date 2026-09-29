package usecases

import "authService/internal/domain"

type RegisterTransportDto struct {
	Email       string `json:"email" validate:"required,email" db:"email"`
	Password    string `json:"password" validate:"required,min=8,max=20" db:"password_hash"`
	UserName    string `json:"userName" validate:"required,min=8,max=64" db:"user_name"`
	DisplayName string `json:"displayName" validate:"required,min=8,max=64" db:"display_name"`
}

type RegisterServiceDto struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
}

type RegisterResult struct {
	User         domain.User
	AccessToken  string
	RefreshToken string
}

// {
//   "email": "Test1111@mail.ru",
//   "password": "TestTest121111!",
//   "userName": "test11111111",
//   "displayName": "test11111111"
// }

// {
//     "success": {
//         "id": "b39e50ce-8bc4-4c0c-9ad1-a0fe03d1e70b",
//         "email": "Test1111@mail.ru",
//         "username": "test11111111"
//     }
// }
