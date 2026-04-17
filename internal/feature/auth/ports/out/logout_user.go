package auth_ports_out

type LogoutUserAuthParams struct {
	DeviceName   string
	RefreshToken string
}

func NewLogoutUserAuthParams(deviceName string, refreshToken string) LogoutUserAuthParams {
	return LogoutUserAuthParams{
		DeviceName:   deviceName,
		RefreshToken: refreshToken,
	}
}
