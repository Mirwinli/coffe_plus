package auth_ports_in

import (
	"github.com/Mirwinli/coffe_plus/internal/core/domain"
	"github.com/google/uuid"
)

type PatchUserParams struct {
	UserID uuid.UUID
	Patch  domain.PatchUser
}

func NewPatchUserParams(patch domain.PatchUser, userID uuid.UUID) PatchUserParams {
	return PatchUserParams{
		Patch:  patch,
		UserID: userID,
	}
}

type PatchUserResult struct {
	User domain.User
}

func NewPatchUserResult(user domain.User) PatchUserResult {
	return PatchUserResult{
		User: user,
	}
}
