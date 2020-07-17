package main

import (
	"html"
	"net/http"
	"net/url"

	"github.com/microcosm-cc/bluemonday"
)

type FormValues map[string][]string

func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		_, _ = w.Write([]byte("MailyGo works!"))
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte("The HTTP method is not allowed, make a POST request"))
		return
	}
	_ = r.ParseForm()
	sanitizedForm := sanitizeForm(&r.PostForm)
	go func() {
		if checkToken(sanitizedForm) && !findBotSubmit(sanitizedForm) && !isBot(sanitizedForm) {
			sendForm(sanitizedForm)
		}
	}()
	sendResponse(sanitizedForm, w)
	return
}

func sanitizeForm(values *url.Values) *FormValues {
	p := bluemonday.StrictPolicy()
	sanitizedForm := make(FormValues)
	for key, values := range *values {
		var sanitizedValues []string
		for _, value := range values {
			sanitizedValues = append(sanitizedValues, html.UnescapeString(p.Sanitize(value)))
		}
		sanitizedForm[html.UnescapeString(p.Sanitize(key))] = sanitizedValues
	}
	return &sanitizedForm
}

func checkToken(values *FormValues) bool {
	if token := appConfig.Token; token != "" && (*values)["_token"][0] != token {
		return false
	}
	return true
}

func findBotSubmit(values *FormValues) bool {
	if len((*values)["submit"]) == 1 {
		return true
	}
	return false
}

func isBot(values *FormValues) bool {
	for _, honeyPot := range appConfig.HoneyPots {
		if len((*values)[honeyPot]) > 0 {
			for _, value := range (*values)[honeyPot] {
				if value != "" {
					return true
				}
			}
		}
	}
	return checkValues(values)
}

func sendResponse(values *FormValues, w http.ResponseWriter) {
	if len((*values)["_redirectTo"]) == 1 && (*values)["_redirectTo"][0] != "" {
		w.Header().Add("Location", (*values)["_redirectTo"][0])
		w.WriteHeader(http.StatusSeeOther)
		_, _ = w.Write([]byte("Go to " + (*values)["_redirectTo"][0]))
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte("Submitted form"))
		return
	}
}
