package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	request "github.com/hunterbdm/hello-requests"
)

func Request(req request.Options) (*request.Response, error) {
	return request.Do(req)
}

func NewRequest(url, proxy string) request.Options {
	return request.Options{
		Method:            "GET",
		URL:               url,
		Body:              "",
		ParseJSONResponse: true,
		FollowRedirects:   true,
		Headers: request.Headers{
			"sec-ch-ua":                 `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36`,
			"sec-ch-ua-mobile":          "?0",
			"upgrade-insecure-requests": "1",
			"user-agent":                "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
			"sec-fetch-site":            "none",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-user":            "?1",
			"sec-fetch-dest":            "document",
			"accept-encoding":           "gzip, deflate, br",
			"accept-language":           "en-US,en;q=0.9",
		},
		HeaderOrder: request.HeaderOrder{
			"sec-ch-ua",
			"sec-ch-ua-mobile",
			"upgrade-insecure-requests",
			"user-agent",
			"accept",
			"sec-fetch-site",
			"sec-fetch-mode",
			"sec-fetch-user",
			"sec-fetch-dest",
			"accept-encoding",
			"accept-language",
			"cookie",
		},
		ClientSettings: &request.ClientSettings{
			IdleTimeoutTime: 10000,
			SkipCertChecks:  false,
			Proxy:           proxy,
		},
	}
}

// Returns JSON data from the HTML script
func FilterHTML(html string, selector string, attribute ...string) (string, error) {
	r := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	s := doc.Find(selector)
	return s.Text(), nil
}
