package entity

import (
	"errors"
	"time"
)

/*
	id				BIGINT	NOT NULL,
  	user_id			BIGINT NOT NULL,
	started_at		timestamp NOT NULL,
	finished_at		timestamp,
*/

var (
	ErrUserIsRequired         = errors.New("user is required")
	ErrStartSessionIsRequired = errors.New("start session is required")
	ErrInvalidFinishedTime    = errors.New("finished time must be greater than start time")
)

type UserSession struct {
	Id         int `json:"id"`
	User       `json:"user"`
	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`
}

func (us *UserSession) Validate() error {
	if us.Id == 0 {
		return ErrIdIsRequired
	}
	if us.User.Id == 0 {
		return ErrUserIsRequired
	}
	if us.StartedAt.IsZero() {
		return ErrStartSessionIsRequired
	}
	// if us.FinishedAt.Compare(us.StartedAt) == -1 {
	// 	return ErrInvalidFinishedTime
	// }

	return nil
}

func NewUserSession(id int, user User, startedAt, finishedAt time.Time) (*UserSession, error) {
	us := &UserSession{
		Id:         id,
		User:       user,
		StartedAt:  startedAt,
		FinishedAt: finishedAt,
	}

	return us, us.Validate()
}
