package auth_ports_out

import "context"

type AuthRepository interface {
	SaveUser(
		ctx context.Context,
		in SaveUserAuthParams,
	) (SaveUserAuthResult, error)
	SaveRefreshToken(
		ctx context.Context,
		in SaveRefreshTokenAuthParams,
	) error
	LoginUser(
		ctx context.Context,
		in LoginUserAuthParams,
	) (LoginUserAuthResult, error)
	GetAndDeleteRefreshToken(
		ctx context.Context,
		in GetRefreshParams,
	) (GetRefreshResult, error)
	LogoutUser(
		ctx context.Context,
		in LogoutUserAuthParams,
	) error
	GetUser(
		ctx context.Context,
		in GetUserParams,
	) (GetUserResult, error)

	BlackListUser(
		ctx context.Context,
		in BlackListParams,
	) error

	IsUserBlackListed(
		ctx context.Context,
		in IsBlackListedParams,
	) (bool, error)
}
