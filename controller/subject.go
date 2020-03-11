package controller

import (
	"TueKan-backend/model"
	"database/sql"
	"fmt"
	"github.com/labstack/echo"
	"net/http"

)

//PostController for Post model
type SubjectController struct {
	DB *sql.DB
}

//NewPostController create  Post controller
func NewSubjectController(db *sql.DB) *SubjectController {
	return &SubjectController{DB: db}
}


func (S *SubjectController)CreateNewSubject( c echo.Context) error{
	subject:=new(model.Subject)

	if err := c.Bind(subject); err != nil{
		return err
	}

	queryString := "INSERT INTO subject(subject_name) VALUES($1)"

	_, err := S.DB.Exec(queryString,subject.SubjectName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusCreated, "Subject Created")

}

func (S *SubjectController)GetAllSubject(c echo.Context)error{

	queryString :="SELECT * FROM subject"

	rows,err := S.DB.Query(queryString)
	if err != nil{
		return err
	}

	defer rows.Close()

	subjects := make([]*model.Subject,0)
	for rows.Next(){
		subject := new(model.Subject)

		err := rows.Scan(&subject.TagID,&subject.SubjectName)
		if err != nil {
			return err
		}

		subjects = append(subjects,subject)
	}

	return c.JSON(http.StatusOK,subjects)
}
