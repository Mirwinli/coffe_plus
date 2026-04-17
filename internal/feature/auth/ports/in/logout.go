package auth_ports_in

type LogoutAuthParams struct {
	DeviceName   string
	RefreshToken string
}

func NewLogoutAuthParams(deviceName string, refreshToken string) LogoutAuthParams {
	return LogoutAuthParams{
		DeviceName:   deviceName,
		RefreshToken: refreshToken,
	}
}
