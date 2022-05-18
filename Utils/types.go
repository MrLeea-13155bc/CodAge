package Utils

type GetQuestionForm struct {
	UserId      int64 `json:"-"`
	Num         int64 `json:"num"`
	SectionId   int64 `json:"sectionId"`
	RequestType int64 `json:"type"`
}

type QuestionList struct {
	QuestionId       int       `json:"questionId"`
	QuestionTitle    string    `json:"questionTitle"`
	QuestionType     int       `json:"questionType"`
	QuestionImageURl string    `json:"questionImage,omitempty"`
	QuestionOptions  []Options `json:"questionOptions"`
}

type Options struct {
	OptionId int    `json:"optionId"`
	Context  string `json:"context"`
}

type Subject struct {
	SubjectId   int64  `json:"subjectId"`
	SubjectName string `json:"subjectName"`
}

type Section struct {
	SectionId   int64  `json:"chapterId"`
	SectionName string `json:"chapterName"`
	FinishedNum int64  `json:"finishedNum"`
	TotalNum    int64  `json:"totalNum"`
}
