package playwright

import (
	"log"

	"github.com/playwright-community/playwright-go"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/playwright".
func init() {
	var p = new(Playwright)
	pw, err := playwright.Run()

	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	p.Self = pw

	modules.Register("k6/x/playwright", p)
}

// Playwright is the k6 extension for a playwright-go client.
type Playwright struct {
	Self *playwright.Playwright
}

// Launch starts the playwright client and launches a browser
func (p *Playwright) Launch(args playwright.BrowserTypeLaunchOptions) *browserWrapper {
	browser, err := p.Self.Chromium.Launch(args)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	return newBrowserWrapper(browser)
}

// Kill closes browser instance and stops puppeteer client
func (p *Playwright) Kill() {
	if err := p.Self.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}
