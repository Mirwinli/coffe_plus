package auth_ports_in

import "context"

type AuthService interface {
	Register(
		ctx context.Context,
		in RegisterAuthParams,
	) (RegisterAuthResult, error)
	Login(
		ctx context.Context,
		in LoginAuthParams,
	) (LoginAuthResult, error)
	PatchUser(
		ctx context.Context,
		in PatchUserParams,
	) (PatchUserResult, error)
	Refresh(
		ctx context.Context,
		in RefreshAuthParams,
	) (RefreshAuthResult, error)
	GetUser(
		ctx context.Context,
		in GetUserParams,
	) (GetUserResult, error)
	GetAllUsers(
		ctx context.Context,
		in GetAllUsersParams,
	) (GetAllUsersResult, error)
	Logout(
		ctx context.Context,
		in LogoutAuthParams,
	) error
}
