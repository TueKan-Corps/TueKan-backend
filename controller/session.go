package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
)

// SubjectController for Session model
type SessionController struct {
	DB *sql.DB
}

// NewSessionController create  Session controller
func NewSessionController(db *sql.DB) *SessionController {
	return &SessionController{DB: db}
}

// GetAll get all sessions
func (s *SessionController) GetAll(c echo.Context) error {

	queryString := "SELECT * FROM session ORDER BY id"

	rows, err := s.DB.Query(queryString)
	if err != nil {
		return err
	}
	defer rows.Close()

	sessions := make([]*model.Session, 0)
	for rows.Next() {
		session := new(model.Session)

		err := rows.Scan(&session.ID,
			&session.AccountID,
			&session.Token)

		if err != nil {
			return err
		}

		sessions = append(sessions, session)
	}

	return c.JSON(http.StatusOK, sessions)
}
