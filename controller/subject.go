package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
)

// SubjectController for Post model
type SubjectController struct {
	DB *sql.DB
}

// NewSubjectController create  Post controller
func NewSubjectController(db *sql.DB) *SubjectController {
	return &SubjectController{DB: db}
}

func (s *SubjectController) CreateNewSubject(c echo.Context) error {

	subjectName := c.FormValue("subject_name")

	queryString := "INSERT INTO subject(subject_name) VALUES($1)"
	_, err := s.DB.Exec(queryString, subjectName)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "Subject Created")

}

func (s *SubjectController) GetAllSubject(c echo.Context) error {

	queryString := "SELECT * FROM subject"

	rows, err := s.DB.Query(queryString)
	if err != nil {
		return err
	}

	defer rows.Close()

	subjects := make([]*model.Subject, 0)
	for rows.Next() {
		subject := new(model.Subject)

		err := rows.Scan(&subject.TagID, &subject.SubjectName)
		if err != nil {
			return err
		}

		subjects = append(subjects, subject)
	}

	return c.JSON(http.StatusOK, subjects)
}
