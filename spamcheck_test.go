package main

import (
	"os"
	"testing"
)

func Test_checkBlacklist(t *testing.T) {
	prepare := func() {
		os.Clearenv()
		_ = os.Setenv("BLACKLIST", "test1,test2")
		appConfig, _ = parseConfig()
	}
	t.Run("Allowed values", func(t *testing.T) {
		prepare()
		if checkBlacklist([]string{"Hello", "How are you?"}) == true {
			t.Error()
		}
	})
	t.Run("Forbidden values", func(t *testing.T) {
		prepare()
		if checkBlacklist([]string{"How are you?", "Hello TeSt1"}) == false {
			t.Error()
		}
	})
}
