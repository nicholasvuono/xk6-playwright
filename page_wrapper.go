package playwright

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"time"

	"github.com/playwright-community/playwright-go"
	"github.com/tidwall/gjson"
	"go.k6.io/k6/js/modules"
)

type pageWrapper struct {
	Page playwright.Page
	vu   modules.VU
}

func newPageWrapper(page playwright.Page, vu modules.VU) *pageWrapper {
	return &pageWrapper{
		Page: page,
		vu:   vu,
	}
}

// Goto wrapper around playwright goto page function that takes in a url and a set of options
func (p *pageWrapper) Goto(url string, opts playwright.PageGotoOptions) {
	if _, err := p.Page.Goto(url, opts); err != nil {
		Throw(p.vu.Runtime(), "Error attempting to goto", err)
	}
}

// WaitForSelector wrapper around playwright waitForSelector page function that takes in a selector and a set of options
func (p *pageWrapper) WaitForSelector(selector string, opts playwright.PageWaitForSelectorOptions) {
	if _, err := p.Page.WaitForSelector(selector, opts); err != nil {
		Throw(p.vu.Runtime(), "Error waiting for selector", err)
	}
}

func (p *pageWrapper) WaitForURL(url string, opts playwright.FrameWaitForURLOptions) {
	if err := p.Page.WaitForURL(url, opts); err != nil {
		Throw(p.vu.Runtime(), "Error waiting for URL", err)
	}
}

func (p *pageWrapper) WaitForEvent(event string, predicate interface{}) {
	p.Page.WaitForEvent(event, predicate)
}

func (p *pageWrapper) Locator(selector string, opts playwright.PageLocatorOptions) *locatorWrapper {
	locator, err := p.Page.Locator(selector, opts)
	if err != nil {
		Throw(p.vu.Runtime(), "Error locating selector", err)
	}
	return newLocatorWrapper(locator, p.vu)
}

func (p *pageWrapper) SetExtraHTTPHeaders(headers map[string]string) {
	err := p.Page.SetExtraHTTPHeaders(headers)
	if err != nil {
		Throw(p.vu.Runtime(), "Error setting extra HTTP headers", err)
	}
}

func (p *pageWrapper) WaitForNavigation(opts playwright.PageWaitForNavigationOptions) {
	_, err := p.Page.WaitForNavigation(opts)
	if err != nil {
		Throw(p.vu.Runtime(), "Error waiting for navigation", err)
	}
}

func (p *pageWrapper) ExpectFileChooser(cb func() error) *fileChooserWrapper {
	result, err := p.Page.ExpectFileChooser(cb)

	if err != nil {
		Throw(p.vu.Runtime(), "Error expecing file chooser", err)
	}

	return newFileChooserWrapper(result, p.vu)
}

// Click wrapper around playwright click page function that takes in a selector and a set of options
func (p *pageWrapper) Click(selector string, opts playwright.PageClickOptions) {
	if err := p.Page.Click(selector, opts); err != nil {
		Throw(p.vu.Runtime(), "Error clicking", err)
	}
}

// Type wrapper around playwright type page function that takes in a selector, string, and a set of options
func (p *pageWrapper) Type(selector string, typedString string, opts playwright.PageTypeOptions) {
	if err := p.Page.Type(selector, typedString, opts); err != nil {
		Throw(p.vu.Runtime(), "Error typing", err)
	}
}

// PressKey wrapper around playwright Press page function that takes in a selector, key, and a set of options
func (p *pageWrapper) PressKey(selector string, key string, opts playwright.PagePressOptions) {
	if err := p.Page.Press(selector, key, opts); err != nil {
		Throw(p.vu.Runtime(), "Error pressing key", err)
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
		Throw(p.vu.Runtime(), "Error taking screenshot", err)
	}
	err = ioutil.WriteFile("Screenshot_"+time.Now().Format("2017-09-07 17:06:06")+".png", image, perm)
	if err != nil {
		Throw(p.vu.Runtime(), "Error writing screenshot to file system", err)
	}
}

// Focus wrapper around playwright focus page function that takes in a selector and a set of options
func (p *pageWrapper) Focus(selector string, opts playwright.PageFocusOptions) {
	if err := p.Page.Focus(selector); err != nil {
		Throw(p.vu.Runtime(), "Error focusing on element", err)
	}
}

// Fill wrapper around playwright fill page function that takes in a selector, text, and a set of options
func (p *pageWrapper) Fill(selector string, filledString string, opts playwright.FrameFillOptions) {
	if err := p.Page.Fill(selector, filledString, opts); err != nil {
		Throw(p.vu.Runtime(), "Error filling input field", err)
	}
}

// SelectOptions wrapper around playwright selectOptions page function that takes in a selector, values, and a set of options
func (p *pageWrapper) SelectOptions(selector string, values playwright.SelectOptionValues, opts playwright.FrameSelectOptionOptions) {
	_, err := p.Page.SelectOption(selector, values, opts)
	if err != nil {
		Throw(p.vu.Runtime(), "Error selecting input option", err)
	}
}

// Check wrapper around playwright check page function that takes in a selector and a set of options
func (p *pageWrapper) Check(selector string, opts playwright.FrameCheckOptions) {
	if err := p.Page.Check(selector, opts); err != nil {
		Throw(p.vu.Runtime(), "Error checking input field", err)
	}
}

// Uncheck wrapper around playwright uncheck page function that takes in a selector and a set of options
func (p *pageWrapper) Uncheck(selector string, opts playwright.FrameUncheckOptions) {
	if err := p.Page.Uncheck(selector, opts); err != nil {
		Throw(p.vu.Runtime(), "Error unchecking input field", err)
	}
}

// DragAndDrop wrapper around playwright draganddrop page function that takes in two selectors(source and target) and a set of options
func (p *pageWrapper) DragAndDrop(sourceSelector string, targetSelector string, opts playwright.FrameDragAndDropOptions) {
	if err := p.Page.DragAndDrop(sourceSelector, targetSelector, opts); err != nil {
		Throw(p.vu.Runtime(), "Error dragging and dropping source/target", err)
	}
}

// Evaluate wrapper around playwright evaluate page function that takes in an expresion and a set of options and evaluates the expression/function returning the resulting information.
func (p *pageWrapper) Evaluate(expression string, opts playwright.PageEvaluateOptions) interface{} {
	returnedValue, err := p.Page.Evaluate(expression, opts)
	if err != nil {
		Throw(p.vu.Runtime(), "Error evaluating expression", err)
	}
	return returnedValue
}

// Reload wrapper around playwright reload page function
func (p *pageWrapper) Reload(opts playwright.PageReloadOptions) {
	if _, err := p.Page.Reload(opts); err != nil {
		Throw(p.vu.Runtime(), "Error reloading page", err)
	}
}

// FirstPaint function that gathers the Real User Monitoring Metrics for First Paint of the current page
func (p *pageWrapper) FirstPaint() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByName('first-paint'))")
	if err != nil {
		Throw(p.vu.Runtime(), "Error getting the first-paint entries", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// FirstContentfulPaint function that gathers the Real User Monitoring Metrics for First Contentful Paint of the current page
func (p *pageWrapper) FirstContentfulPaint() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByName('first-contentful-paint'))")
	if err != nil {
		Throw(p.vu.Runtime(), "Error getting the first-contentful-paint entries", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// TimeToMinimallyInteractive function that gathers the Real User Monitoring Metrics for Time to Minimally Interactive of the current page (based on the first input)
func (p *pageWrapper) TimeToMinimallyInteractive() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByType('first-input'))")
	if err != nil {
		Throw(p.vu.Runtime(), "Error getting the time to minimally interactive metrics", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.startTime").Uint()
}

// FirstInputDelay function that gathers the Real User Monitoring Metrics for First Input Delay of the current page
func (p *pageWrapper) FirstInputDelay() uint64 {
	entries, err := p.Page.Evaluate("JSON.stringify(performance.getEntriesByType('first-input'))")
	if err != nil {
		Throw(p.vu.Runtime(), "Error getting the first input delay metrics", err)
	}
	entriesToString := fmt.Sprintf("%v", entries)
	return gjson.Get(entriesToString, "0.processingStart").Uint() - gjson.Get(entriesToString, "0.startTime").Uint() //https://web.dev/fid/  for calc
}

func (p *pageWrapper) SetViewportSize(viewPortSize playwright.ViewportSize) {
	err := p.Page.SetViewportSize(viewPortSize.Width, viewPortSize.Height)
	if err != nil {
		Throw(p.vu.Runtime(), "Error setting viewPort size", err)
	}
}

func (p *pageWrapper) ExpectNavigation(cb func() error, options ...playwright.PageWaitForNavigationOptions) {
	navigationOptions := make([]interface{}, 0)
	for _, option := range options {
		navigationOptions = append(navigationOptions, option)
	}
	_, err := newExpectWrapper(p.Page.WaitForNavigation, navigationOptions, cb)
	if err != nil {
		Throw(p.vu.Runtime(), "Error expecting navigation", err)
	}
}

func (p *pageWrapper) Url() string {
	return p.Page.URL()
}

func (p *pageWrapper) Pause() {
	err := p.Page.Pause()
	if err != nil {
		Throw(p.vu.Runtime(), "Error pausing", err)
	}
}

func (p *pageWrapper) SetDefaultNavigationTimeout(timeout float64) {
	p.Page.SetDefaultNavigationTimeout(timeout)
}

func (p *pageWrapper) SetDefaultTimeout(timeout float64) {
	p.Page.SetDefaultTimeout(timeout)
}

func (p *pageWrapper) ExpectedDialog(cb func() error) *dialogWrapper {
	result, err := p.Page.ExpectedDialog(cb)

	if err != nil {
		Throw(p.vu.Runtime(), "Error expecting dialog", err)
	}
	return newDialogWrapper(result, p.vu)
}

func (p *pageWrapper) AcceptDialog() {
	p.Page.On("dialog", func(dialog playwright.Dialog) {
		fmt.Print("accepting the dialog")
		dialog.Accept()
	})
}
