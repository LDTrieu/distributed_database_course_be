package auth

type DataJWT struct {
	SessionID string `json:"sessionId"`
	UserName  string `json:"userName"`
	FullName  string `json:"fullName"`
	Role      string `json:"role"` // WARNING: role should not be saved in token (RFC 9068)
}
