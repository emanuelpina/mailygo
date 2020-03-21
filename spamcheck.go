package main

import (
	"github.com/google/safebrowsing"
	"net/url"
	"strings"
)

// Returns true when it spam
func checkValues(values FormValues) bool {
	var urlsToCheck []string
	for _, value := range values {
		for _, singleValue := range value {
			if strings.Contains(singleValue, "http") {
				parsed, e := url.Parse(singleValue)
				if parsed != nil && e == nil {
					urlsToCheck = append(urlsToCheck, singleValue)
				}
			}
		}
	}
	return checkUrls(urlsToCheck)
}

// Only tests when GOOGLE_API_KEY is set
// Returns true when it spam
func checkUrls(urlsToCheck []string) bool {
	if len(appConfig.GoogleApiKey) < 1 || len(urlsToCheck) == 0 {
		return false
	}
	sb, err := safebrowsing.NewSafeBrowser(safebrowsing.Config{
		APIKey: appConfig.GoogleApiKey,
		ID:     "MailyGo",
	})
	if err != nil {
		return false
	}
	allThreats, err := sb.LookupURLs(urlsToCheck)
	if err != nil {
		return false
	}
	for _, threats := range allThreats {
		if len(threats) > 0 {
			// Unsafe url, mark as spam
			return true
		}
	}
	return false
}
