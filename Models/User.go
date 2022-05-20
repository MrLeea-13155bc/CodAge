package Models

import (
	"QuestionSearch/Utils"
	"database/sql"
	"errors"
	"log"
	"time"
)

func Register(info Utils.RegisterForm) bool {
	affair, _ := Utils.MDB().Begin()
	template := `Insert Into User Set StudentNum=?,Password = ?`
	result, err := affair.Exec(template, info.StudentNum, info.Password)
	if err != nil {
		log.Println("[Register]", err)
		return false
	}
	uid, _ := result.LastInsertId()
	if uid > 0 {
		template = `Insert Into UserInfo Set UserId = ?,AttendDate = ?`
		result, err = affair.Exec(template, uid, time.Now().Unix())
		if err == nil {
			affair.Commit()
			return true
		}
	}
	affair.Rollback()
	return false
}

func Login(info Utils.RegisterForm) int64 {
	template := `Select UserId,Password From User Where StudentNum = ? Limit 1`
	rows, err := Utils.MDB().Query(template, info.StudentNum)
	if err != nil {
		log.Println("[Login]", err)
		return -1
	}
	defer rows.Close()
	if !rows.Next() {
		return -1
	}
	var correctPassword string
	var userId int64
	rows.Scan(&userId, &correctPassword)
	if correctPassword == info.Password {
		return userId
	}
	return -1
}

func UpdateUserInfo(info Utils.UserInfo) bool {
	template := `Update UserInfo Set RealName = ?,NickName=?,Major=?,Birthday=?,Phone=?,Email=? Where UserId = ? limit 1`
	result, err := Utils.MDB().Exec(template, info.RealName, info.NickName, info.Major, info.Birthday, info.Phone, info.Email, info.UserId)
	if err != nil {
		log.Println("[UpdateUserInfo]", err)
		return false
	}
	num, _ := result.RowsAffected()
	return num == 1
}

func GetUserInfo(uid int64) (result Utils.UserInfo, err error) {
	var rows *sql.Rows
	template := `Select UserId,RealName, NickName, AttendDate, Major, Birthday, Phone, Email From UserInfo Where UserId = ?`
	rows, err = Utils.MDB().Query(template, uid)
	if err != nil {
		return
	}
	defer rows.Close()
	if !rows.Next() {
		err = errors.New("not Fount")
		return
	}
	err = rows.Scan(&result.UserId, &result.RealName, &result.NickName, &result.AttendDate, &result.Major, &result.Birthday, &result.Phone, &result.Email)
	return
}

func GetShowInfo(uid int64) (info Utils.UserShowInfo, err error) {
	template := `
Select * From 
(Select  NickName,Icon  From UserInfo Where UserId = ? ) A  Join 
(Select StudentNum From User Where UserId = ?) B Join
(Select COALESCE(Count(QuestionId),0),COALESCE(Sum(isCorrect),0) From (
    Select QuestionId,isCorrect From QuestionsFinish Where UserId = ?
) D ) C;
`
	var rows *sql.Rows
	rows, err = Utils.MDB().Query(template, uid, uid, uid)
	if err != nil {
		return
	}
	if !rows.Next() {
		err = errors.New("not Found")
	}
	err = rows.Scan(&info.UserName, &info.Icon, &info.UserNum, &info.TotalNum, &info.CorrectNum)
	return
}
