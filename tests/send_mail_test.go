package test

import (
	"github.com/stretchr/testify/assert"
	"quickstart/utils"
	"quickstart/utils/service"
	
	_ "github.com/go-sql-driver/mysql"
	_ "quickstart/models/setting"
	"testing"
)

//测试发送邮件功能
/*
go test -v ./tests/send_mail_test.go
*/
func TestSendMail(t *testing.T) {
	to := []string{"taotao-it@youtulink.com", "xtt691373656@iCloud.com"}
	body := map[string]string{"subject": "徐涛涛测试邮件", "content": "测试内容测试内容"}
	status, message := service.SendMail("taotao-it@youtulink.com", "Xutaotao1215.", to, body, "")
	assert.Equal(t, status, true)
	assert.Equal(t, message, "")
}

func TestSendMailAttach(t *testing.T) {
	to := []string{"taotao-it@youtulink.com", "xtt691373656@iCloud.com"}
	body := map[string]string{"subject": "徐涛涛测试邮件", "content": "测试内容测试内容"}
	status, message := service.SendMailAttach("taotao-it@youtulink.com", "Xutaotao1215.", to,
		body, "../tests/send_mail_test.go", "",[]string{})
	assert.Equal(t, status, true)
	assert.Equal(t, message, "")
}

func init() {
	utils.DataBaseConnection("test_quick_start")
}
