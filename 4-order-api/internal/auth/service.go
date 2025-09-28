// auth/service.go
package auth

import (
	"errors"
	"4-order-api/internal/user"
	"4-order-api/internal/session"
	"4-order-api/pkg/sms"
	"time"
)

type AuthService struct {
	UserRepository    *user.UserRepository
	SessionRepository *session.SessionRepository
	SMSService        *sms.SMSService
}

func NewAuthService(userRepository *user.UserRepository, sessionRepository *session.SessionRepository, smsService *sms.SMSService) *AuthService {
	return &AuthService{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
		SMSService:        smsService,
	}
}

func (service *AuthService) InitiateAuth(phone string) (string, error) {
	
	service.SessionRepository.CleanExpired()

	session := session.NewSession(phone)
	_, err := service.SessionRepository.Create(session)
	if err != nil {
		return "", err
	}

	err = service.SMSService.SendCode(phone, session.Code)
	if err != nil {
		return "", err
	}

	return session.SessionID, nil
}

func (service *AuthService) VerifyCode(sessionID, code string) (string, error) {
	// Находим сессию
	session, err := service.SessionRepository.FindBySessionID(sessionID)
	if err != nil {
		return "", errors.New("invalid session")
	}

	if session.Used {
		return "", errors.New("session already used")
	}

	if time.Now().Unix() > session.ExpiresAt {
		return "", errors.New("session expired")
	}

	if session.Code != code {
		return "", errors.New("invalid code")
	}

	usr, err := service.UserRepository.FindByPhone(session.Phone)
	if err != nil {
		newUser := &user.User{Phone: session.Phone}
		usr, err = service.UserRepository.Create(newUser)
		if err != nil {
			return "", err
		}
	}

	service.SessionRepository.MarkAsUsed(sessionID)

	return usr.Phone, nil
}