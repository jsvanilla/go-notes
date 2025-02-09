// patron de diseño que utiliza interfaces y composicion sobre la herencia
package main

import "fmt"

// SMS Email

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {

}

func (SMSNotification) SendNotification(){
	fmt.Println("Sending Notifications via SMS")
}

func (SMSNotification) GetSender() ISender{
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {

}

func (SMSNotificationSender) GetSenderMethod () string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel () string {
	return "Twilio"
}