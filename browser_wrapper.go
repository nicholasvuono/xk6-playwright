package playwright

import (
	"log"

	"github.com/playwright-community/playwright-go"
	"go.k6.io/k6/js/modules"
)

type browserWrapper struct {
	Browser playwright.Browser
	vu      modules.VU
}

func newBrowserWrapper(browser playwright.Browser, vu modules.VU) *browserWrapper {
	return &browserWrapper{
		Browser: browser,
		vu:      vu,
	}
}

func (b *browserWrapper) Close() {
	if err := b.Browser.Close(); err != nil {
		Throw(b.vu.Runtime(), "Error closing browser", err)
	}
}

func (b *browserWrapper) NewPage() *pageWrapper {
	page, err := b.Browser.NewPage()

	if err != nil {
		Throw(b.vu.Runtime(), "Error creating new page", err)
	}

	return newPageWrapper(page, b.vu)
}

// Cookies wrapper around playwright cookies fetch function
func (b *browserWrapper) Cookies() []*playwright.BrowserContextCookiesResult {
	if b.Browser == nil {
		log.Panicf("It looks like there's no browser attached. please call `launch()` before accessing the cookies.")
	}

	if len(b.Browser.Contexts()) == 0 {
		log.Panicf("It looks like there's no browser contexts are attached. please use `goto()` before accessing the cookies.")
	}

	cookies, err := b.Browser.Contexts()[0].Cookies()
	if err != nil {
		Throw(b.vu.Runtime(), "Error getting cookies from the default context", err)
	}
	return cookies
}
