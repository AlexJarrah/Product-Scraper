package utils

import "net/url"

func IsValidURL(u string) bool {
	// Parse the URL
	parsedURL, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	// Check if the Scheme and Host are not empty
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}

	return true
}

func RemoveURLParameters(u string) (string, error) {
	res, err := url.Parse(u)
	if err != nil {
		return "", err
	}
	res.RawQuery = ""
	return res.String(), nil
}
