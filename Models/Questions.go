package Models

import (
	"QuestionSearch/Utils"
	"log"
)

func GetQuestions(data Utils.GetQuestionForm) ([]Utils.QuestionList, error) {
	var questions = make([]Utils.QuestionList, 0)
	var question = Utils.QuestionList{}
	template := `Select QuestionId,Title,TitleHasImage,QuestionType From QuestionInfo A Natural Join (
    Select QuestionId,QuestionType From QuestionList Where CategoryId = ?
    And QuestionId not In (
        Select QuestionId From QuestionsFinish Where isCorrect = 1 And UserId = ?
    )
) B `
	rows, err := Utils.MDB().Query(template, data.Category, data.UserId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&question.QuestionId, &question.QuestionTitle, &question.QuestionImageURl, &question.QuestionType)
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
	}
	return questions, nil
}
