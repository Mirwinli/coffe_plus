package core_contextKeys

type contextKey string

const (
	UserIDCtxKey      contextKey = "user_id"
	UserRoleCtxKey    contextKey = "role"
	JWTAccessIDCtxKey contextKey = "jwt_access_id"
)
