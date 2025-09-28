// auth/payload.go
package auth

type InitiateAuthRequest struct {
	Phone string `json:"phone" validate:"required,phone"`
}

type InitiateAuthResponse struct {
	SessionID string `json:"sessionId"`
}

type VerifyCodeRequest struct {
	SessionID string `json:"sessionId" validate:"required"`
	Code      string `json:"code" validate:"required,len=4"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}