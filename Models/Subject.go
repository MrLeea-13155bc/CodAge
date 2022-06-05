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
	defer rows.Close()
	var item Utils.Subject
	for rows.Next() {
		rows.Scan(&item.SubjectId, &item.SubjectName)
		result = append(result, item)
	}
	return result, nil
}

func GetAllChapter(cid, uid int64) ([]Utils.Section, error) {
	template := `Select SectionId,SectionName From Section Where SubjectId = ?`
	rows, err := Utils.MDB().Query(template, cid)
	if err != nil {
		log.Println("err,", err)
		return nil, err
	}
	defer rows.Close()
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
		rows, _ = Utils.MDB().Query(template, item.SectionId, uid, item.SectionId)
		defer rows.Close()
		rows.Next()
		rows.Scan(&sections[id].TotalNum, &sections[id].FinishedNum)
	}
	return sections, nil
}

func SearchSubject(subjectName string, userId int64) ([]Utils.Subject, error) {
	template := `Select * From Subject Where Match(SubjectName) Against(? In Boolean Mode) And SubjectId Not In (
    Select SubjectId From UserLesson Where UserId = ?
    )
limit 10`
	rows, err := Utils.MDB().Query(template, subjectName, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]Utils.Subject, 0)
	item := Utils.Subject{}
	for rows.Next() {
		rows.Scan(&item.SubjectId, &item.SubjectName)
		result = append(result, item)
	}
	return result, nil
}

func AddSubject(subjectId, userId int64) bool {
	template := `Insert Into UserLesson Set UserId = ?,SubjectId = ?`
	result, err := Utils.MDB().Exec(template, userId, subjectId)
	if err != nil {
		log.Println(err)
		return false
	}
	_, err = result.LastInsertId()
	return err == nil
}

func InsertSubject(subject Utils.Subject) bool {
	template := `Select SubjectId From Subject Where SubjectName = ?`
	rows, err := Utils.MDB().Query(template, subject.SubjectName)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return false
	}
	template = `Insert Into Subject Set SubjectName = ?`
	_, err = Utils.MDB().Exec(template, subject.SubjectName)
	return err == nil
}

func InsertSection(info Utils.SectionInfo) bool {
	template := `Select SectionId From Section Where SectionName = ? And SubjectId = ?`
	rows, err := Utils.MDB().Query(template, info.SectionName, info.SubjectId)
	if err != nil {
		log.Println(err)
		return false
	}
	defer rows.Close()
	if rows.Next() {
		log.Println(err)
		return false
	}
	template = `Insert Into Section Set SubjectId=?,SectionName=?`
	_, err = Utils.MDB().Exec(template, info.SubjectId, info.SectionName)
	return err == nil
}
