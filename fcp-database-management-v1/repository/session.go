package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type SessionsRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailName(name string) error
	SessionAvailToken(token string) (model.Session, error)
}

type sessionsRepoImpl struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionsRepoImpl {
	return &sessionsRepoImpl{db}
}

func (s *sessionsRepoImpl) AddSessions(session model.Session) error {
	data := s.db.Create(&session)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) DeleteSession(token string) error {
	sesion := &model.Session{}
	data := s.db.Debug().Model(sesion).Where("token = ?", token).Delete(sesion)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) UpdateSessions(session model.Session) error {
	sesion := &model.Session{}
	data := s.db.Debug().Model(sesion).Where("username = ?", session.Username).UpdateColumns(
		map[string]interface{}{
			"token":  session.Token,
			"expiry": session.Expiry,
		},
	)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailName(name string) error {
	sesion := &model.Session{}
	data := s.db.Debug().Model(sesion).Where("username = ?", name).First(sesion)
	if data.Error != nil {
		return data.Error
	}
	return nil // TODO: replace this
}

func (s *sessionsRepoImpl) SessionAvailToken(token string) (model.Session, error) {
	sesion := &model.Session{}
	data := s.db.Debug().Model(sesion).Where("token = ?", token).First(sesion)
	if data.Error != nil {
		return model.Session{}, data.Error
	}
	return *sesion, nil // TODO: replace this
}
