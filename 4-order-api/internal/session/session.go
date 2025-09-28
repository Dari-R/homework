// internal/session/session.go
package session

import (
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	gorm.Model
	SessionID string `gorm:"uniqueIndex;not null"`
	Phone     string `gorm:"not null"`
	Code      string `gorm:"not null"`
	ExpiresAt int64  `gorm:"not null"` // Unix timestamp
	Used      bool   `gorm:"default:false"`
}

func NewSession(phone string) *Session {
	return &Session{
		SessionID: generateSessionID(),
		Phone:     phone,
		Code:      generateCode(),
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(), // 10 минут
		Used:      false,
	}
}

func generateSessionID() string {
	// Генерация уникального ID сессии
	return "sess_" + time.Now().Format("20060102150405") + "_" + randomString(8)
}

func generateCode() string {
	// Генерация 4-значного кода
	return randomNumber(4)
}

func randomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomNumber(n int) string {
	const digits = "0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = digits[rand.Intn(len(digits))]
	}
	return string(b)
}