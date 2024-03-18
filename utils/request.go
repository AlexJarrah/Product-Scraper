package utils

import (
	"context"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"

	"github.com/Noooste/azuretls-client"
	"github.com/PuerkitoBio/goquery"
)

// Creates a new TLS session configured for the specified URL and optional proxy
func NewSession(url_, proxy string) (*azuretls.Session, error) {
	// Parse the provided URL
	parsedURL, err := url.Parse(url_)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}

	// Create a new TLS session with context
	session := azuretls.NewSessionWithContext(context.Background())

	// Configure ClientHelloSpec to use the latest Chrome version
	session.GetClientHelloSpec = azuretls.GetLastChromeVersion

	// Set session timeout to 10 seconds
	session.SetTimeout(10 * time.Second)

	// Configure proxy if provided
	if proxy != "" {
		if err := session.SetProxy(proxy); err != nil {
			return nil, fmt.Errorf("invalid proxy: %w", err)
		}
	}

	// Set request headers for the session
	session.OrderedHeaders = azuretls.OrderedHeaders{
		{"accept", "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8"},
		{"accept-language", "en-US,en;q=0.7"},
		{"authority", parsedURL.Host},
		{"cache-control", "max-age=0"},
		{"connection", "keep-alive"},
		{"content-type", "application/json"},
		{"cookie", ""},
		{"correlation-context", "v=1,ms.b.tel.market=en-US"},
		{"if-modified-since", "Fri, 15 Mar 2024 07:49:53 GMT"},
		{"if-none-match", "0x8DC44C4801ABBCF"},
		{"origin", parsedURL.Scheme + "://" + parsedURL.Host},
		{"sec-ch-ua", "\"Chromium\";v=\"122\", \"Not(A:Brand\";v=\"24\", \"Brave\";v=\"122\""},
		{"sec-ch-ua-mobile", "?0"},
		{"sec-ch-ua-model", "\"\""},
		{"sec-ch-ua-platform", "\"Linux\""},
		{"sec-ch-ua-platform-version", "\"6.7.9\""},
		{"sec-fetch-dest", "image"},
		{"sec-fetch-mode", "no-cors"},
		{"sec-fetch-site", "same-origin"},
		{"sec-fetch-user", "?1"},
		{"sec-gpc", "1"},
		{"upgrade-insecure-requests", "1"},
		{"user-agent", getRandomUserAgent()}}

	// Return the configured session
	return session, nil
}

// Returns a random user-agent
func getRandomUserAgent() string {
	return useragents[rand.Intn(len(useragents))]
}

// Parses the provided HTML and returns the content of the specified selector
// or the value of the attribute if specified
func FilterHTML(html string, selector string, attribute ...string) (string, error) {
	// Create a new reader from the HTML
	r := strings.NewReader(html)

	// Parse the HTML using goquery
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	// Find the specified element
	s := doc.Find(selector)

	// Return the contents of the element if no attribute is specified
	if attribute == nil {
		return s.Text(), nil
	}

	// Retrieve the value of the specified attribute
	attr, exists := s.Attr(attribute[0])
	if !exists {
		return "", fmt.Errorf("attribute %s not found", attribute[0])
	}

	// Return the value of the attribute
	return attr, nil
}
