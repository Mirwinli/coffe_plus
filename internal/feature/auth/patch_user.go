package auth_service

import (
	"context"
	"fmt"

	auth_ports_in "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/in"
	auth_ports_out "github.com/Mirwinli/coffe_plus/internal/feature/auth/ports/out"
)

func (s *AuthService) PatchUser(
	ctx context.Context,
	in auth_ports_in.PatchUserParams,
) (auth_ports_in.PatchUserResult, error) {
	getParams := auth_ports_out.NewGetUserParams(in.UserID)
	getUserResult, err := s.authRepository.GetUser(ctx, getParams)
	if err != nil {
		return auth_ports_in.PatchUserResult{}, fmt.Errorf(
			"get user from repository: %w", err,
		)
	}

	user := getUserResult.User
	if err = user.ApplyPatch(in.Patch); err != nil {
		return auth_ports_in.PatchUserResult{}, fmt.Errorf(
			"apply patch user: %w", err,
		)
	}

	params := auth_ports_out.NewUpdateUserParams(user)
	patchedUser, err := s.authRepository.UpdateUser(ctx, params)
	if err != nil {
		return auth_ports_in.PatchUserResult{}, fmt.Errorf(
			"update user from repository: %w", err,
		)
	}

	return auth_ports_in.NewPatchUserResult(patchedUser.User), nil
}
