package Models

import (
	"QuestionSearch/Utils"
	"log"
	"strconv"
	"time"
)

func GetQuestions(data Utils.GetQuestionForm) ([]Utils.QuestionList, error) {
	var questions = make([]Utils.QuestionList, 0)
	var question = Utils.QuestionList{}
	template := `Select QuestionId,Title,ImageUrl,QuestionType,CorrectAnswer From QuestionInfo A Natural Join (
    Select QuestionId,QuestionType From QuestionList Where SectionId = ?
    And QuestionId not In (
        Select QuestionId From QuestionsFinish Where isCorrect = 1 And UserId = ?
    )
) B Order By Rand() Limit ?`
	rows, err := Utils.MDB().Query(template, data.SectionId, data.UserId, data.Num)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&question.QuestionId, &question.QuestionTitle, &question.QuestionImageURl, &question.QuestionType, &question.CorrectAnswer)
		questions = append(questions, question)
	}
	for id, item := range questions {
		var options = make([]Utils.Options, 0)
		var option Utils.Options
		template = `Select OptionId,OptionTitle From Options Where QuestionId = ?`
		rows, err = Utils.MDB().Query(template, item.QuestionId)
		if err != nil {
			log.Println(err)
			continue
		}
		for rows.Next() {
			rows.Scan(&option.OptionId, &option.Context)
			options = append(options, option)
		}
		questions[id].QuestionOptions = options
		go func(questionId int, correctAnswer string) {
			Utils.RDB().Set(strconv.Itoa(item.QuestionId), item.CorrectAnswer, time.Hour)
		}(item.QuestionId, item.CorrectAnswer)
		questions[id].CorrectAnswer = ""
	}
	return questions, nil
}

func CheckAnswers(answer []Utils.Answer, uid int64) ([]Utils.Answer, int, int, error) {
	var correct int
	affair, _ := Utils.MDB().Begin()
	template := `Replace Into QuestionsFinish Set UserId = ?,QuestionId = ?,LastFinishDate = ?,isCorrect= ?;`
	for id, item := range answer {
		correctAnswer, err := Utils.RDB().Get(strconv.FormatInt(item.QuestionId, 10)).Result()
		if err != nil {
			log.Println(err)
			continue
		}
		tmp := item.Answers == correctAnswer
		if tmp {
			correct++
		}
		affair.Exec(template, uid, item.QuestionId, time.Now().Unix(), tmp)
		answer[id].Answers = correctAnswer
	}
	affair.Commit()
	return answer, correct, len(answer), nil
}

func AddQuestion(list Utils.InsertForm) bool {
	affair, _ := Utils.MDB().Begin()
	template := `Insert Into QuestionList Set QuestionType=?,SectionId = ?`
	result, err := affair.Exec(template, list.QuestionType, list.SectionId)
	if err != nil {
		affair.Rollback()
		return false
	}
	id, _ := result.LastInsertId()
	if list.QuestionType < 2 {
		for key, item := range list.QuestionOptions {
			template = `Insert Into Options Set QuestionId=?,OptionTitle=?`
			thisResult, _ := affair.Exec(template, id, item.Context)
			lastId, err := thisResult.LastInsertId()
			if err != nil {
				affair.Rollback()
				return false
			}
			if key == 0 {
				list.CorrectAnswer = Utils.AddOptionsId(list.CorrectAnswer, lastId)
			}
		}
	}
	template = `Insert Into QuestionInfo Set QuestionId = ?,CreateDate=?,LastChangeDate=?,Title=?,CorrectAnswer=?,ImageUrl=?`
	_, err = affair.Exec(template, id, time.Now().Unix(), time.Now().Unix(), list.QuestionTitle, list.CorrectAnswer, list.QuestionImageURl)
	if err != nil {
		affair.Rollback()
		return false
	}
	affair.Commit()
	return true
}
