package mapper

import (
	"github.com/DenisHoliahaR/go-beautyhub/internal/domain"
	"github.com/DenisHoliahaR/go-beautyhub/internal/transport/http/dto"
)

func UserToGetUserResponse(u *domain.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
		Phone:      u.Phone,
		CreatedAt:  u.CreatedAt,
	}
}

func UserToCreateUserResponse(u *domain.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
		Phone:      u.Phone,
		CreatedAt:  u.CreatedAt,
	}
}

func UserToUpdateUserResponse(u *domain.User) dto.GetUserResponse {
	return dto.GetUserResponse{
		ID:         u.ID,
		FirstName:  u.FirstName,
		SecondName: u.SecondName,
		Email:      u.Email,
		Phone:      u.Phone,
		CreatedAt:  u.CreatedAt,
	}
}

func UsersToGetUserListResponse(users []*domain.User) dto.GetUserListResponse {
	resp := dto.GetUserListResponse{
		Users: make([]dto.GetUserResponse, len(users)),
	}
	for i, u := range users {
		resp.Users[i] = UserToGetUserResponse(u)
	}
	return resp
}