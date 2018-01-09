package notice

import(
	"testing"
	"log"
)
func Test_sms(t *testing.T){
	a:=&SmsTemp{}
	a.Daychange="13131"
	a.Hourchange="3131"
	a.Username="frankie"
	a.Nowprice="31313"
	err:=SendSms("18262622659",a)
	if err!= nil{
		log.Println(err)
	}
}