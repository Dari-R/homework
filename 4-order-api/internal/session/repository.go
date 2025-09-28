// internal/session/repository.go
package session

import (
	"4-order-api/pkg/res/db"
	"time"
)

type SessionRepository struct {
	Database *db.Db
}

func NewSessionRepository(database *db.Db) *SessionRepository {
	return &SessionRepository{
		Database: database,
	}
}

func (repo *SessionRepository) Create(session *Session) (*Session, error) {
	res := repo.Database.DB.Create(session)
	if res.Error != nil {
		return nil, res.Error
	}
	return session, nil
}

func (repo *SessionRepository) FindBySessionID(sessionID string) (*Session, error) {
	var session Session
	res := repo.Database.DB.First(&session, "session_id = ?", sessionID)
	if res.Error != nil {
		return nil, res.Error
	}
	return &session, nil
}

func (repo *SessionRepository) MarkAsUsed(sessionID string) error {
	res := repo.Database.DB.Model(&Session{}).Where("session_id = ?", sessionID).Update("used", true)
	return res.Error
}

func (repo *SessionRepository) CleanExpired() error {
	now := time.Now().Unix()
	res := repo.Database.DB.Where("expires_at < ? OR used = ?", now, true).Delete(&Session{})
	return res.Error
}