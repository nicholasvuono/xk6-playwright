package playwright

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/tidwall/gjson"
)

type pageWrapper struct {
	Page playwright.Page
}

func newPageWrapper(page playwright.Page) *pageWrapper {
	return &pageWrapper{
		Page: page,
	}
}

// Goto wrapper around playwright goto page function that takes in a url and a set of options
func (p *pageWrapper) Goto(url string, opts playwright.PageGotoOptions) {
	if _, err := p.Page.Goto(url, opts); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
}

// WaitForSelector wrapper around playwright waitForSelector page function that takes in a selector and a set of options
func (p *pageWrapper) WaitForSelector(selector string, opts playwright.PageWaitForSelectorOptions) {
	if _, err := p.Page.WaitForSelector(selector, opts); err != nil {
		log.Fatalf("error with waiting for the selector: %v", err)
	}
}

func (p *pageWrapper) WaitForURL(url string, opts playwright.FrameWaitForURLOptions) {
	if err := p.Page.WaitForURL(url, opts); err != nil {
		log.Fatalf("error with waiting for the url: %v", err)
	}
}

func (p *pageWrapper) WaitForEvent(event string, predicate interface{}) {
	p.Page.WaitForEvent(event, predicate)
}

func (p *pageWrapper) Locator(selector string, opts playwright.PageLocatorOptions) *locatorWrapper {
	locator, err := p.Page.Locator(selector, opts)
	if err != nil {
		log.Fatalf("error with locator: %v", err)
	}
	return newLocatorWrapper(locator)
}

func (p *pageWrapper) SetExtraHTTPHeaders(headers map[string]string) {
	err := p.Page.SetExtraHTTPHeaders(headers)
	if err != nil {
		log.Fatalf("SetExtraHTTPHeaders function error: %v", err)
	}
}

func (p *pageWrapper) WaitForNavigation(opts playwright.PageWaitForNavigationOptions) {
	_, err := p.Page.WaitForNavigation(opts)
	if err != nil {
		log.Fatalf("WaitForNavigation function error: %v", err)
	}
}

func (p *pageWrapper) ExpectFileChooser(cb func() error) *fileChooserWrapper {
	result, err := p.Page.ExpectFileChooser(cb)

	if err != nil {
		log.Fatalf("error with expecting file chooser: %v", err)
	}

	return newFileChooserWrapper(result)
}

// Click wrapper around playwright click page function that takes in a selector and a set of options
func (p *pageWrapper) Click(selector string, opts playwright.PageClickOptions) {
	if err := p.Page.Click(selector, opts); err != nil {
		log.Fatalf("error with clicking: %v", err)
	}
}

// Type wrapper around playwright type page function that takes in a selector, string, and a set of options
func (p *pageWrapper) Type(selector string, typedString string, opts playwright.PageTypeOptions) {
	if err := p.Page.Type(selector, typedString, opts); err != nil {
		log.Fatalf("error with typing: %v", err)
	}
}

// PressKey wrapper around playwright Press page function that takes in a selector, key, and a set of options
func (p *pageWrapper) PressKey(selector string, key string, opts playwright.PagePressOptions) {
	if err := p.Page.Press(selector, key, opts); err != nil {
		log.Fatalf("error with pressing the key: %v", err)
	}
}

// Sleep wrapper around playwright waitForTimeout page function that sleeps for the given `timeout` in milliseconds
func (p *pageWrapper) Sleep(time float64) {
	p.Page.WaitForTimeout(time)
}

func (p *pageWrapper) WaitForTimeout(time float64) {
	p.Page.WaitForTimeout(time)
}

// Screenshot wrapper around playwright screenshot page function that attempts to take and save a png image of the current screen.
func (p *pageWrapper) Screenshot(filename string, perm fs.FileMode, opts playwright.PageScreenshotOptions) {
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
func (p *pageWrapper) Focus(selector string, opts playwright.PageFocusOptions) {
	if err := p.Page.Focus(selector); err != nil {
		log.Fatalf("error focusing on the element: %v", err)
	}
}

// Fill wrapper around playwright fill page function that takes in a selector, text, and a set of options
func (p *pageWrapper) Fill(selector string, filledString string, opts playwright.FrameFillOptions) {
	if err := p.Page.Fill(selector, filledString, opts); err != nil {
		log.Fatalf("error with fill: %v", err)
	}
}

// SelectOptions wrapper around playwright selectOptions page function that takes in a selector, values, and a set of options
func (p *pageWrapper) SelectOptions(selector string, values playwright.SelectOptionValues, opts playwright.FrameSelectOptionOptions) {
	_, err := p.Page.SelectOption(selector, values, opts)
	if err != nil {
		log.Fatalf("error selecting the option: %v", err)
	}
}

// Check wrapper around playwright check page function that takes in a selector and a set of options
func (p *pageWrapper) Check(selector string, opts playwright.FrameCheckOptions) {
	if err := p.Page.Check(selector, opts); err != nil {
		log.Fatalf("error with checking the field: %v", err)
	}
}

// Uncheck wrapper around playwright uncheck page function that takes in a selector and a set of options
func (p *pageWrapper) Uncheck(selector string, opts playwright.FrameUncheckOptions) {
	if err := p.Page.Uncheck(selector, opts); err != nil {
		log.Fatalf("error with unchecking the field: %v", err)
	}
}

// DragAndDrop wrapper around playwright draganddrop page function that takes in two selectors(source and target) and a set of options
func (p *pageWrapper) DragAndDrop(sourceSelector string, targetSelector string, opts playwright.FrameDragAndDropOptions) {
	if err := p.Page.DragAndDrop(sourceSelector, targetSelector, opts); err != nil {
		log.Fatalf("error with drag and drop: %v", err)
	}
}

// Evaluate wrapper around playwright evaluate page function that takes in an expresion and a set of options and evaluates the expression/function returning the resulting information.
func (p *pageWrapper) Evaluate(expression string, opts playwright.PageEvaluateOptions) interface{} {
	returnedValue, err := p.Page.Evaluate(expression, opts)
	if err != nil {
		log.Fatalf("error evaluating the expression: %v", err)
	}
	return returnedValue
}

// Reload wrapper around playwright reload page function
func (p *pageWrapper) Reload() {
	if _, err := p.Page.Reload(); err != nil {
		log.Fatalf("error with reloading the page: %v", err)
	}
}

// FirstPaint function that gathers the Real User Monitoring Metrics for First Paint of the current page
func (p *pageWrapper) FirstPaint() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByName('first-paint'))")
	if err != nil {
		log.Fatalf("error with getting the first-paint entries: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// FirstContentfulPaint function that gathers the Real User Monitoring Metrics for First Contentful Paint of the current page
func (p *pageWrapper) FirstContentfulPaint() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByName('first-contentful-paint'))")
	if err != nil {
		log.Fatalf("error with getting the first-contentful-paint entries: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// TimeToMinimallyInteractive function that gathers the Real User Monitoring Metrics for Time to Minimally Interactive of the current page (based on the first input)
func (p *pageWrapper) TimeToMinimallyInteractive() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByType('first-input'))")
	if err != nil {
		log.Fatalf("error with getting the first-input entries time to minimally interactive metrics: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// FirstInputDelay function that gathers the Real User Monitoring Metrics for First Input Delay of the current page
func (p *pageWrapper) FirstInputDelay() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByType('first-input'))")
	if err != nil {
		log.Fatalf("error with getting the first-input entries for first input delay metrics: %v", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.processingStart").Uint() - gjson.Get(entriesToString, "0.startTime").Uint() //https://web.dev/fid/  for calc
}

func (p *pageWrapper) SetViewportSize(width, height int) {
	err := p.Page.SetViewportSize(width, height)
	if err != nil {
		log.Fatalf("Error while setting viewPort size :: %v", err)
	}
}
