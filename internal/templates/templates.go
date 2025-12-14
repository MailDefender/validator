package templates

import (
	_ "embed"
)

var (
	//go:embed expired.html
	expiredHtmlPage string

	//go:embed failure.html
	failureHtmlPage string

	//go:embed success.html
	successHtmlPage string

	//go:embed already_validated.html
	alreadyValidatedPage string
)

func ExpiredPage() string {
	return expiredHtmlPage
}

func FailurePage() string {
	return failureHtmlPage
}

func SuccessPage() string {
	return successHtmlPage
}

func AlreadyValidatedPage() string {
	return alreadyValidatedPage
}
