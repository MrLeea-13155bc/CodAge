package Utils

type GetQuestionForm struct {
	UserId      int64 `json:"-"`
	Num         int64 `json:"num"`
	SectionId   int64 `json:"chapterId"`
	RequestType int64 `json:"type"`
}

type QuestionList struct {
	QuestionId       int       `json:"questionId"`
	QuestionTitle    string    `json:"questionTitle"`
	QuestionType     int       `json:"questionType"`
	QuestionImageURl string    `json:"questionImage,omitempty"`
	QuestionOptions  []Options `json:"questionOptions,omitempty"`
	CorrectAnswer    string    `json:"-"`
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

type RegisterForm struct {
	StudentNum string `json:"studentNum"`
	Password   string `json:"password"`
}

type UserInfo struct {
	UserId     int64  `json:"-"`
	RealName   string `json:"realName"`
	AttendDate int    `json:"attendDate"`
	Major      string `json:"major"`
}

type UserShowInfo struct {
	UserName   string `json:"userName"`
	UserNum    int    `json:"userNum"`
	Icon       string `json:"icon"`
	Subjects   []Subject
	TotalNum   int `json:"totalNum"`
	CorrectNum int `json:"correctNum"`
}

type Answer struct {
	QuestionId int64  `json:"id"`
	Answers    string `json:"answer"`
}
