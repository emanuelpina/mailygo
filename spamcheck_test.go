package main

import (
	"os"
	"testing"
)

func Test_checkSpamlist(t *testing.T) {
	prepare := func() {
		os.Clearenv()
		_ = os.Setenv("SPAMLIST", "test1,test2")
		appConfig, _ = parseConfig()
	}
	t.Run("Allowed values", func(t *testing.T) {
		prepare()
		if checkSpamlist([]string{"Hello", "How are you?"}) == true {
			t.Error()
		}
	})
	t.Run("Forbidden values", func(t *testing.T) {
		prepare()
		if checkSpamlist([]string{"How are you?", "Hello TeSt1"}) == false {
			t.Error()
		}
	})
}
