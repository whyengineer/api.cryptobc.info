package notice

import(
	"github.com/denverdino/aliyungo/sms"
	"encoding/json"
	"errors"
)
const(
	ACCESS_KEY="LTAI7VBH4vyx95PL"
	ACCESS_SECRET="67C0xEbUA5pVGzjl0A3ReaQTY0zVnl"
	SMS_FREE_SIGN_NAME="福blockchain"
	SMS_TEMPLATE_CODE="SMS_120405007"
)
type SmsTemp struct{
	Username string `json:"username"`
	Nowprice string `json:"nowprice"`
	Hourchange string `json:"hourchange"`
	Daychange string `json:"daychange"`
}
func SendSms(phoneNum string,para *SmsTemp)error{
	client := sms.NewDYSmsClient(ACCESS_KEY, ACCESS_SECRET)
	parab, err := json.Marshal(para)
	if err != nil{
		return err
	}
	a:=&sms.SendSmsArgs{}
	a.PhoneNumbers=phoneNum
	a.SignName=SMS_FREE_SIGN_NAME
	a.TemplateCode=SMS_TEMPLATE_CODE
	a.TemplateParam=string(parab)
	c,err:=client.SendSms(a)
	if err !=nil{
		return err
	}
	if c.Code!="OK"{
		return errors.New(c.Message)
	}
	return err
}

