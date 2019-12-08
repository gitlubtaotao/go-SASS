package service

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"log"
	"net/smtp"
	"quickstart/models/setting"
	"strconv"
	"strings"
)



func SendMailTest() {
	var (
		from        = "system@youtulink.com"
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
		msg         = []byte("From: system@youtulink.com" + "\r\n" + "To: taotao-it@youtulink.com;xtt691373656@iCloud.com" + "\r\n" + "Subject: 徐涛涛" + "\r\n" + "ContentType:" + contentType + "\r\n\r\n" + "学车")
		recipients  = []string{"taotao-it@youtulink.com", "xtt691373656@iCloud.com"}
	)
	hostname := "smtp.mxhichina.com"
	auth := smtp.PlainAuth("", from, "Youtulink1234", hostname)
	
	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		logs.Error(err)
		log.Fatal(err)
	}
}

/*
SendMail: 发送邮件
from: 发件人邮箱
password: 发件人密码
to: 收件人
body: 发件内容
*/
func SendMail(from string, password string, to []string, body map[string]string,mailType string) (status bool, message string) {
	o := orm.NewOrm()
	emailConfig := setting.EmailConfig{Id: 1}
	err := o.Read(&emailConfig)
	if err != nil {
		return false, "Mail configuration not read"
	}
	if from == "" {
		from = emailConfig.UserName
		password = emailConfig.Password
	}
	toHeader := strings.Join(to, ";")
	subject := body["subject"]
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	address := emailConfig.Address + ":" + strconv.Itoa(int(emailConfig.Port))
	
	auth := smtp.PlainAuth("", from, password, emailConfig.Address)
	logs.Info(from, password, toHeader, to, emailConfig.Address, subject, contentType, address, auth)
	msg := []byte("From: " + from + "\r\n" + "To: " + toHeader + "\r\n" +
		"Subject: "+subject + "\r\n" + "ContentType:" + contentType + "\r\n\r\n" + body["content"])
	err = smtp.SendMail(address, auth, from, to, msg)
	if err != nil{
		logs.Error(err.Error())
		return false,err.Error()
	}
	return true, message
}
