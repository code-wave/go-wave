package entity

import (
	"log"
	"net/smtp"

	"github.com/code-wave/go-wave/infrastructure/errors"
)

type ValidEmail struct {
	UserID    int64  `json:"user_id"`
	ValidCode int64  `json:"email_code"`
	Email     string `json:"email"`
}

func (v *ValidEmail) SendValidCode() *errors.RestErr {
	auth := smtp.PlainAuth("", "example@live.com", "pwd", "smtp.live.com")

	from := "example@live.com"
	to := []string{"atg950831@gmail.com"} // 복수 수신자 가능

	// validCode := strconv.FormatInt(v.ValidCode, 10)
	// 메시지 작성
	headerSubject := "Subject: 회원가입 인증번호입니다.\r\n"
	headerBlank := "\r\n"

	body := "메일 테스트입니다 \r\n"

	// body := "메일 테스트입니다 \r\n" + "인증번호: " + validCode + "\r\n"
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	err := smtp.SendMail("smtp.live.com:25", auth, from, to, msg)
	if err != nil {
		restErr := errors.NewInternalServerError("sendmail error " + err.Error())
		log.Println("sendmail error " + err.Error())
		return restErr
	}

	return nil
}
