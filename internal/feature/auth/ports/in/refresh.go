package auth_ports_in

type RefreshAuthParams struct {
	RefreshToken string
	DeviceName   string
	IpAddress    *string
}

func NewRefreshAuthParams(refreshToken string, deviceName string, ipAddress *string) RefreshAuthParams {
	return RefreshAuthParams{
		RefreshToken: refreshToken,
		DeviceName:   deviceName,
		IpAddress:    ipAddress,
	}
}

type RefreshAuthResult struct {
	RefreshToken string
	AccessToken  string
}

func NewRefreshAuthResult(refreshToken string, accessToken string) RefreshAuthResult {
	return RefreshAuthResult{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}
}
