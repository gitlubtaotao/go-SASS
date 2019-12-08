package service

import (
	"bytes"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
	"quickstart/models/setting"
	"strconv"
	"strings"
)

/*
SendMail: 发送邮件
from: 发件人邮箱
password: 发件人密码
to: 收件人
body: 发件内容
net/smtp/SendMail 发送风格
*/
func SendMail(from string, password string, to []string, body map[string]string, mailType string) (status bool, message string) {
	emailConfig := getSetting()
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
		"Subject: " + subject + "\r\n" + "ContentType:" + contentType + "\r\n\r\n" + body["content"])
	err := smtp.SendMail(address, auth, from, to, msg)
	if err != nil {
		return false,err.Error()
	}
	return true, message
}


/*
发送附件的邮件形式
将内容以html的格式进行发送
*/
func SendMailAttach(from, password string, to []string, body map[string]string,
	file string, cc []string) (status bool, message string) {
	e := email.NewEmail()
	emailConfig := getSetting()
	if from == "" {
		from = emailConfig.UserName
		password = emailConfig.Password
	}
	
	e.From = from
	e.To = to
	//设置抄送
	if cc != nil {
		e.Cc = cc
	}
	e.Subject = body["subject"]
	t, err := template.ParseFiles("../send_email_template.html")
	if err != nil {
		return false, err.Error()
	}
	result := new(bytes.Buffer)
	_ = t.Execute(result, struct {
		Content string
		Subject string
	}{
		Content: body["content"],
		Subject: body["subject"],
	})
	e.HTML = result.Bytes()
	if file != "" {
		_, err = e.AttachFile(file)
	}
	if err != nil {
		return false, err.Error()
	}
	address := emailConfig.Address + ":" + strconv.Itoa(int(emailConfig.Port))
	auth := smtp.PlainAuth("", from, password, emailConfig.Address)
	err = e.Send(address, auth)
	if err != nil {
		logs.Error(err.Error())
		return false, err.Error()
	}
	return true, ""
}

//获取系统邮件配置
func getSetting() (record setting.EmailConfig) {
	o := orm.NewOrm()
	emailConfig := setting.EmailConfig{Id: 1}
	err := o.Read(&emailConfig)
	if err != nil {
		panic(err)
	} else {
		return emailConfig
	}
}
