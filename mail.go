package main

import (
	"bytes"
	"fmt"
	"net/smtp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func sendForm(values *FormValues) {
	recipient := findRecipient(values)
	sendMail(recipient, buildMessage(recipient, time.Now(), values))
}

func buildMessage(recipient string, date time.Time, values *FormValues) string {
	msgBuffer := &bytes.Buffer{}
	_, _ = fmt.Fprintf(msgBuffer, "From: Forms <%s>", appConfig.Sender)
	_, _ = fmt.Fprintln(msgBuffer)
	_, _ = fmt.Fprintf(msgBuffer, "To: %s", recipient)
	_, _ = fmt.Fprintln(msgBuffer)
	if replyTo := findReplyTo(values); replyTo != "" {
		_, _ = fmt.Fprintf(msgBuffer, "Reply-To: %s", replyTo)
		_, _ = fmt.Fprintln(msgBuffer)
	}
	_, _ = fmt.Fprintf(msgBuffer, "Date: %s", date.Format(time.RFC1123Z))
	_, _ = fmt.Fprintln(msgBuffer)
	_, _ = fmt.Fprintf(msgBuffer, "Subject: New submission on %s", findFormName(values))
	_, _ = fmt.Fprintln(msgBuffer)
	_, _ = fmt.Fprintln(msgBuffer)
	bodyValues := removeMetaValues(values)
	var keys []string
	for key := range *bodyValues {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		_, _ = fmt.Fprint(msgBuffer, key)
		_, _ = fmt.Fprint(msgBuffer, ": ")
		_, _ = fmt.Fprintln(msgBuffer, strings.Join((*bodyValues)[key], ", "))
	}
	return msgBuffer.String()
}

func sendMail(to, message string) {
	auth := smtp.PlainAuth("", appConfig.SmtpUser, appConfig.SmtpPassword, appConfig.SmtpHost)
	err := smtp.SendMail(appConfig.SmtpHost+":"+strconv.Itoa(appConfig.SmtpPort), auth, appConfig.Sender, []string{to}, []byte(message))
	if err != nil {
		fmt.Println("Failed to send mail:", err.Error())
	}
}

func findRecipient(values *FormValues) string {
	if len((*values)["_to"]) == 1 && (*values)["_to"][0] != "" {
		formDefinedRecipient := (*values)["_to"][0]
		for _, allowed := range appConfig.AllowedRecipients {
			if formDefinedRecipient == allowed {
				return formDefinedRecipient
			}
		}
	}
	return appConfig.DefaultRecipient
}

func findFormName(values *FormValues) string {
	if len((*values)["_formName"]) == 1 && (*values)["_formName"][0] != "" {
		return (*values)["_formName"][0]
	}
	return "a form"
}

func findReplyTo(values *FormValues) string {
	if len((*values)["_replyTo"]) == 1 && (*values)["_replyTo"][0] != "" {
		return (*values)["_replyTo"][0]
	}
	return ""
}

func removeMetaValues(values *FormValues) *FormValues {
	cleanedValues := FormValues{}
	for key, value := range *values {
		if !strings.HasPrefix(key, "_") {
			cleanedValues[key] = value
		}
	}
	return &cleanedValues
}
