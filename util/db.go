package util

import (
	"database/sql"
	"time"
)

type PostEndDate struct {
	id    int
	endAt string
}

func ClearOutdatedPost(db *sql.DB) error {

	queryString := "SELECT id,end_at FROM post"

	rows, err := db.Query(queryString)
	if err != nil {
		return err
	}
	defer rows.Close()

	postEndDates := make([]*PostEndDate, 0)
	for rows.Next() {

		postEndDate := new(PostEndDate)

		err := rows.Scan(&postEndDate.id, &postEndDate.endAt)
		if err != nil {
			return err
		}

		postEndDates = append(postEndDates, postEndDate)
	}

	for i := 0; i < len(postEndDates); i++ {
		t, err := time.Parse("01-02-2006 15:04", postEndDates[i].endAt)
		if err != nil {
			return err
		}

		if t.Before(time.Now()) {
			queryString = "DELETE FROM account WHERE id=$1"
			_, err = db.Exec(queryString, postEndDates[i].id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
