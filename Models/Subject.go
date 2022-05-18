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

func GetAllChapter(cid int64) ([]Utils.Section, error) {
	template := `Select SectionId,SectionName From Section Where SubjectId = ?`
	rows, err := Utils.MDB().Query(template, cid)
	if err != nil {
		log.Println("err,", err)
		return nil, err
	}
	sections := make([]Utils.Section, 0)
	section := Utils.Section{}
	for rows.Next() {
		rows.Scan(&section.SectionId, &section.SectionName)
		sections = append(sections, section)
	}
	template = `Select * From (Select Count(QuestionId) as Total From QuestionList Where SectionId = ?) A 
    Join (Select Count(QuestionId) as Num From QuestionList C Where QuestionId  In
       (Select QuestionId From QuestionsFinish Where UserId = ?) And SectionId = ?) B`
	for id, item := range sections {
		rows, _ = Utils.MDB().Query(template, item.SectionId, 1, item.SectionId)
		rows.Next()
		rows.Scan(&sections[id].TotalNum, &sections[id].FinishedNum)
	}
	return sections, nil
}
