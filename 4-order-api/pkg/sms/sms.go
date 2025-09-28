package sms

import "log"

type SMSService struct {

}

func NewSMSService() *SMSService {
	return &SMSService{}
}

func (s *SMSService) SendCode(phone, code string) error {
	log.Printf("SMS sent to %s: Your verification code is %s", phone, code)
	return nil
}