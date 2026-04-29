package user

func ToProfileResponse(user User) ProfileResponse {
	return ProfileResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}
