package playwright

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/tidwall/gjson"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/playwright".
func init() {
	modules.Register("k6/x/playwright", new(Playwright))
}

// Playwright is the k6 extension for a playwright-go client.
type Playwright struct {
	Self           *playwright.Playwright
	Browser        playwright.Browser
	BrowserContext playwright.BrowserContext
	Page           playwright.Page
}

// Launch starts the playwright client and launches a browser
func (p *Playwright) Launch(args playwright.BrowserTypeLaunchOptions) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(args)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	p.Self = pw
	p.Browser = browser
}

// LaunchPersistent starts the playwright client and launches a browser with a persistent context
func (p *Playwright) LaunchPersistent(dir string, args playwright.BrowserTypeLaunchPersistentContextOptions) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.LaunchPersistentContext(dir, args)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	p.Self = pw
	p.BrowserContext = browser
}

// Connect attaches Playwright to an existing browser instance
func (p *Playwright) Connect(url string, args playwright.BrowserTypeConnectOverCDPOptions) {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.ConnectOverCDP(url, args)
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	context := browser.Contexts()[0]

	p.Self = pw
	p.Browser = browser
	p.Page = context.Pages()[0]
}

// NewPage opens a new page within the browser
func (p *Playwright) NewPage() {
	page, err := p.newPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	p.Page = page
}

// Kill closes browser instance and stops puppeteer client
func (p *Playwright) Kill() {
	if err := p.closeBrowser(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := p.Self.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}
}

//---------------------------------------------------------------------
//                         ACTIONS
//---------------------------------------------------------------------

// Goto wrapper around playwright goto page function that takes in a url and a set of options
func (p *Playwright) Goto(url string, opts playwright.PageGotoOptions) {
	if _, err := p.Page.Goto(url, opts); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}

// WaitForSelector wrapper around playwright waitForSelector page function that takes in a selector and a set of options
func (p *Playwright) WaitForSelector(selector string, opts playwright.PageWaitForSelectorOptions) {
	if _, err := p.Page.WaitForSelector(selector, opts); err != nil {
		log.Fatalf("error with waiting for the selector: %v", err)
	}
}

// Click wrapper around playwright click page function that takes in a selector and a set of options
func (p *Playwright) Click(selector string, opts playwright.PageClickOptions) {
	if err := p.Page.Click(selector, opts); err != nil {
		log.Fatalf("error with clicking: %v", err)
	}
}

// Type wrapper around playwright type page function that takes in a selector, string, and a set of options
func (p *Playwright) Type(selector string, typedString string, opts playwright.PageTypeOptions) {
	if err := p.Page.Type(selector, typedString, opts); err != nil {
		log.Fatalf("error with typing: %v", err)
	}
}

// PressKey wrapper around playwright Press page function that takes in a selector, key, and a set of options
func (p *Playwright) PressKey(selector string, key string, opts playwright.PagePressOptions) {
	if err := p.Page.Press(selector, key, opts); err != nil {
		log.Fatalf("error with pressing the key: %v", err)
	}
}

// Sleep wrapper around playwright waitForTimeout page function that sleeps for the given `timeout` in milliseconds
func (p *Playwright) Sleep(time float64) {
	p.Page.WaitForTimeout(time)
}

// Screenshot wrapper around playwright screenshot page function that attempts to take and save a png image of the current screen.
func (p *Playwright) Screenshot(filename string, perm fs.FileMode, opts playwright.PageScreenshotOptions) {
	image, err := p.Page.Screenshot(opts)
	if err != nil {
		log.Fatalf("error with taking a screenshot: %v", err)
	}
	err = ioutil.WriteFile("Screenshot_"+time.Now().Format("2017-09-07 17:06:06")+".png", image, perm)
	if err != nil {
		log.Fatalf("error with writing the screenshot to the file system: %v", err)
	}
}

// Focus wrapper around playwright focus page function that takes in a selector and a set of options
func (p *Playwright) Focus(selector string, opts playwright.PageFocusOptions) {
	if err := p.Page.Focus(selector); err != nil {
		log.Fatalf("error focusing on the element: %v", err)
	}
}

// Fill wrapper around playwright fill page function that takes in a selector, text, and a set of options
func (p *Playwright) Fill(selector string, filledString string, opts playwright.FrameFillOptions) {
	if err := p.Page.Fill(selector, filledString, opts); err != nil {
		log.Fatalf("error with fill: %v", err)
	}
}

// SelectOptions wrapper around playwright selectOptions page function that takes in a selector, values, and a set of options
func (p *Playwright) SelectOptions(selector string, values playwright.SelectOptionValues, opts playwright.FrameSelectOptionOptions) {
	_, err := p.Page.SelectOption(selector, values, opts)
	if err != nil {
		log.Fatalf("error selecting the option: %v", err)
	}
}

// Check wrapper around playwright check page function that takes in a selector and a set of options
func (p *Playwright) Check(selector string, opts playwright.FrameCheckOptions) {
	if err := p.Page.Check(selector, opts); err != nil {
		log.Fatalf("error with checking the field: %v", err)
	}
}

// Uncheck wrapper around playwright uncheck page function that takes in a selector and a set of options
func (p *Playwright) Uncheck(selector string, opts playwright.FrameUncheckOptions) {
	if err := p.Page.Uncheck(selector, opts); err != nil {
		log.Fatalf("error with unchecking the field: %v", err)
	}
}

// DragAndDrop wrapper around playwright draganddrop page function that takes in two selectors(source and target) and a set of options
func (p *Playwright) DragAndDrop(sourceSelector string, targetSelector string, opts playwright.FrameDragAndDropOptions) {
	if err := p.Page.DragAndDrop(sourceSelector, targetSelector, opts); err != nil {
		log.Fatalf("error with drag and drop: %v", err)
	}
}

// Evaluate wrapper around playwright evaluate page function that takes in an expresion and a set of options and evaluates the expression/function returning the resulting information.
func (p *Playwright) Evaluate(expression string, opts playwright.PageEvaluateOptions) interface{} {
	returnedValue, err := p.Page.Evaluate(expression, opts)
	if err != nil {
		log.Fatalf("error evaluating the expression: %v", err)
	}
	return returnedValue
}

// Reload wrapper around playwright reload page function
func (p *Playwright) Reload() {
	if _, err := p.Page.Reload(); err != nil {
		log.Fatalf("error with reloading the page: %v", err)
	}
}

// FirstPaint function that gathers the Real User Monitoring Metrics for First Paint of the current page
func (p *Playwright) FirstPaint() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByName('first-paint'))")
	if err != nil {
		log.Fatalf("error with getting the first-paint entries: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// FirstContentfulPaint function that gathers the Real User Monitoring Metrics for First Contentful Paint of the current page
func (p *Playwright) FirstContentfulPaint() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByName('first-contentful-paint'))")
	if err != nil {
		log.Fatalf("error with getting the first-contentful-paint entries: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// TimeToMinimallyInteractive function that gathers the Real User Monitoring Metrics for Time to Minimally Interactive of the current page (based on the first input)
func (p *Playwright) TimeToMinimallyInteractive() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByType('first-input'))")
	if err != nil {
		log.Fatalf("error with getting the first-input entries time to minimally interactive metrics: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// FirstInputDelay function that gathers the Real User Monitoring Metrics for First Input Delay of the current page
func (p *Playwright) FirstInputDelay() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByType('first-input'))")
	if err != nil {
		log.Fatalf("error with getting the first-input entries for first input delay metrics: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.processingStart").Uint() - gjson.Get(entriesToString, "0.startTime").Uint() //https://web.dev/fid/  for calc
}

// Cookies wrapper around playwright cookies fetch function
func (p *Playwright) Cookies() []*playwright.BrowserContextCookiesResult {
	cookies, err := p.cookies()
	if err != nil {
		log.Fatalf("error with getting the cookies: %v", err)
	}
	return cookies
}

//---------------------------------------------------------------------
//                         Helpers
//---------------------------------------------------------------------

// newPage creates a new page and returns it either with or without a context
func (p *Playwright) newPage() (playwright.Page, error) {
	if p.Browser != nil {
		return p.Browser.NewPage()
	}
	if p.BrowserContext != nil {
		return p.BrowserContext.NewPage()
	}
	return nil, errors.New("no browser or browser context attached")
}

// closeBrowser closes the browser and the browser context
func (p *Playwright) closeBrowser() error {
	if p.Browser != nil {
		return p.Browser.Close()
	}
	if p.BrowserContext != nil {
		return p.BrowserContext.Close()
	}
	return errors.New("no browser or browser context attached")
}

// cookies returns the cookies from the browser context or from browser persistent context
func (p *Playwright) cookies() ([]*playwright.BrowserContextCookiesResult, error) {
	if p.Browser != nil && len(p.Browser.Contexts()) > 0 {
		return p.Browser.Contexts()[0].Cookies()
	}
	if p.BrowserContext != nil {
		return p.BrowserContext.Cookies()
	}
	return nil, errors.New("no browser or browser context attached")
}
