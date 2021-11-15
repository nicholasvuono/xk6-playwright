package playwright

import (
	"log"

	"github.com/mxschmitt/playwright-go"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initilization, avaialable to
// import from JS as "k6/x/playwright".
func init() {
	modules.Register("k6/x/playwright", new(Playwright))
}

// Playwright is the k6 extension for a playwright-go client.
type Playwright struct {}

// Starts the playwright client and launches a browser
func (*Playwright) Launch(args []string) (*playwright.Playwright, playwright.Browser) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Args: args,
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	return pw, browser
}

// Opens a new page within the browser
func (*Playwright) NewPage(browser playwright.Browser) playwright.Page {
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	return page
}

// Closes browser instance and stops puppeteer client
func(*Playwright) Kill(pw *playwright.Playwright, browser playwright.Browser) {
	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}