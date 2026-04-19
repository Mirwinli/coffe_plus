package auth_ports_out

type IsBlackListedParams struct {
	IDAccess string
}

func NewIsBlackListedParams(idAccess string) IsBlackListedParams {
	return IsBlackListedParams{
		IDAccess: idAccess,
	}

}
