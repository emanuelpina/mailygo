package main

import (
	"net/url"
	"strings"

	"github.com/google/safebrowsing"
)

// Returns true when it spam
func checkValues(values *FormValues) bool {
	var urlsToCheck []string
	var allValues []string
	for _, value := range *values {
		for _, singleValue := range value {
			allValues = append(allValues, singleValue)
			if strings.Contains(singleValue, "http") {
				parsed, e := url.Parse(singleValue)
				if parsed != nil && e == nil {
					urlsToCheck = append(urlsToCheck, singleValue)
				}
			}
		}
	}
	return checkBlacklist(allValues) || checkUrls(urlsToCheck)
}

func checkBlacklist(values []string) bool {
	for _, value := range values {
		for _, blacklistedString := range appConfig.Blacklist {
			if strings.Contains(strings.ToLower(value), strings.ToLower(blacklistedString)) {
				return true
			}
		}
	}
	return false
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
