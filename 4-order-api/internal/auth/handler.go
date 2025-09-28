// auth/handler.go
package auth

import (
	configs "4-order-api/config"
	"4-order-api/pkg/jwt"
	"4-order-api/pkg/req"
	"4-order-api/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	
	router.HandleFunc("POST /auth/initiate", handler.InitiateAuth())
	router.HandleFunc("POST /auth/verify", handler.VerifyCode())
}

// InitiateAuth - первый этап авторизации
func (handler *AuthHandler) InitiateAuth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[InitiateAuthRequest](&w, r)
		if err != nil {
			return
		}

		sessionID, err := handler.AuthService.InitiateAuth(payload.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data := InitiateAuthResponse{
			SessionID: sessionID,
		}
		res.Json(w, data, 200)
	}
}

// VerifyCode - второй этап авторизации
func (handler *AuthHandler) VerifyCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := req.HandleBody[VerifyCodeRequest](&w, r)
		if err != nil {
			return
		}

		phone, err := handler.AuthService.VerifyCode(payload.SessionID, payload.Code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		token, err := jwt.NewJWT(handler.Config.Auth.Secret).Create(jwt.JWTData{Phone: phone})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := VerifyCodeResponse{
			Token: token,
		}
		res.Json(w, data, 200)
	}
}