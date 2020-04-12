package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

func Test_findRecipient(t *testing.T) {
	prepare := func() {
		os.Clearenv()
		_ = os.Setenv("ALLOWED_TO", "mail@example.com,test@example.com")
		_ = os.Setenv("EMAIL_TO", "mail@example.com")
		appConfig, _ = parseConfig()
	}
	t.Run("No recipient specified", func(t *testing.T) {
		prepare()
		values := &FormValues{}
		result := findRecipient(values)
		if result != "mail@example.com" {
			t.Error()
		}
	})
	t.Run("Multiple recipients specified", func(t *testing.T) {
		prepare()
		values := &FormValues{
			"_to": {"abc@example.com", "def@example.com"},
		}
		result := findRecipient(values)
		if result != "mail@example.com" {
			t.Error()
		}
	})
	t.Run("Allowed recipient specified", func(t *testing.T) {
		prepare()
		values := &FormValues{
			"_to": {"test@example.com"},
		}
		result := findRecipient(values)
		if result != "test@example.com" {
			t.Error()
		}
	})
	t.Run("Forbidden recipient specified", func(t *testing.T) {
		prepare()
		values := &FormValues{
			"_to": {"forbidden@example.com"},
		}
		result := findRecipient(values)
		if result != "mail@example.com" {
			t.Error()
		}
	})
}

func Test_findFormName(t *testing.T) {
	t.Run("No form name", func(t *testing.T) {
		if "a form" != findFormName(&FormValues{}) {
			t.Error()
		}
	})
	t.Run("Multiple form names", func(t *testing.T) {
		if "a form" != findFormName(&FormValues{"_formName": {"Test", "ABC"}}) {
			t.Error()
		}
	})
	t.Run("Form name", func(t *testing.T) {
		if "Test" != findFormName(&FormValues{"_formName": {"Test"}}) {
			t.Error()
		}
	})
}

func Test_findReplyTo(t *testing.T) {
	t.Run("No replyTo", func(t *testing.T) {
		if "" != findReplyTo(&FormValues{}) {
			t.Error()
		}
	})
	t.Run("Multiple replyTo", func(t *testing.T) {
		if "" != findReplyTo(&FormValues{"_replyTo": {"test@example.com", "test2@example.com"}}) {
			t.Error()
		}
	})
	t.Run("replyTo", func(t *testing.T) {
		if "test@example.com" != findReplyTo(&FormValues{"_replyTo": {"test@example.com"}}) {
			t.Error()
		}
	})
}

func Test_removeMetaValues(t *testing.T) {
	t.Run("Remove meta values", func(t *testing.T) {
		result := removeMetaValues(&FormValues{
			"_test": {"abc"},
			"test":  {"def"},
		})
		want := FormValues{
			"test": {"def"},
		}
		if !reflect.DeepEqual(*result, want) {
			t.Error()
		}
	})
}

func Test_buildMessage(t *testing.T) {
	t.Run("Test message", func(t *testing.T) {
		os.Clearenv()
		_ = os.Setenv("EMAIL_TO", "mail@example.com")
		_ = os.Setenv("ALLOWED_TO", "mail@example.com,test@example.com")
		_ = os.Setenv("EMAIL_FROM", "forms@example.com")
		appConfig, _ = parseConfig()
		values := &FormValues{
			"_formName":   {"Testform"},
			"_replyTo":    {"reply@example.com"},
			"Testkey":     {"Testvalue"},
			"Another Key": {"Test", "ABC"},
		}
		date := time.Now()
		result := buildMessage("test@example.com", date, values)
		if !strings.Contains(result, "Reply-To: reply@example.com") {
			t.Error()
		}
		if !strings.Contains(result, "Subject: New submission on Testform") {
			t.Error()
		}
		if !strings.Contains(result, "Testkey: Testvalue") {
			t.Error()
		}
		if !strings.Contains(result, "Another Key: Test, ABC") {
			t.Error()
		}
	})
}
