package util

import (
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"
)

func Post(url string, data string) (body string, err error) {
	resp, err1 := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	defer resp.Body.Close()
	if err1 != nil {
		err = err1
	}
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
	}
	body = string(b)
	return
}

func Get(url string) (body string, err error) {
	resp, err1 := http.Get(url)
	defer resp.Body.Close()
	if err1 != nil {
		err = err1
	}
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
	}
	body = string(b)
	return
}

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	return smtp.SendMail(host, auth, user, send_to, msg)
}
