package Models

import (
	"QuestionSearch/Utils"
	"log"
)

func GetAllSubject(userId int64) ([]Utils.Subject, error) {
	var result = make([]Utils.Subject, 0)
	template := `Select * From (
    	Select SubjectId From UserLesson Where UserId = ?
	) A natural Join Subject `
	rows, err := Utils.MDB().Query(template, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var item Utils.Subject
	for rows.Next() {
		rows.Scan(&item.SubjectId, &item.SubjectName)
		result = append(result, item)
	}
	return result, nil
}
