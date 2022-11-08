package playwright

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

type browserWrapper struct {
	Browser playwright.Browser
}

func newBrowserWrapper(browser playwright.Browser) *browserWrapper {
	return &browserWrapper{
		Browser: browser,
	}
}

func (b *browserWrapper) Close() {
	if err := b.Browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
}

func (b *browserWrapper) NewPage() *pageWrapper {
	page, err := b.Browser.NewPage()

	if err != nil {
		log.Fatalf("could not create new page: %v", err)
	}

	return newPageWrapper(page)
}

// Cookies wrapper around playwright cookies fetch function
func (b *browserWrapper) Cookies() []*playwright.BrowserContextCookiesResult {
	if b.Browser == nil {
		log.Fatalf("looks like there's no browser attached. please call `launch()` before accessing the cookies.")
	}

	if len(b.Browser.Contexts()) == 0 {
		log.Fatalf("looks like there's no browser contexts are attached. please use `goto()` before accessing the cookies.")
	}

	cookies, err := b.Browser.Contexts()[0].Cookies()
	if err != nil {
		log.Fatalf("error with getting the cookies from the default context: %v", err)
	}
	return cookies
}
