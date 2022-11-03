package playwright

import (
	"log"
	"github.com/playwright-community/playwright-go"
)

type locatorWrapper struct {
	Locator playwright.Locator
}

type MouseButton string
type KeyboardModifier string
type PageClickOptionsPosition struct {
	X *float64 `json:"x"`
	Y *float64 `json:"y"`
}

type PageClickOptions struct {
	// Defaults to `left`.
	Button *MouseButton `json:"button"`
	// defaults to 1. See [UIEvent.detail].
	ClickCount *int `json:"clickCount"`
	// Time to wait between `mousedown` and `mouseup` in milliseconds. Defaults to 0.
	Delay *float64 `json:"delay"`
	// Whether to bypass the [actionability](../actionability.md) checks. Defaults to `false`.
	Force *bool `json:"force"`
	// Modifier keys to press. Ensures that only these modifiers are pressed during the
	// operation, and then restores current modifiers back. If not specified, currently
	// pressed modifiers are used.
	Modifiers []KeyboardModifier `json:"modifiers"`
	// Actions that initiate navigations are waiting for these navigations to happen and
	// for pages to start loading. You can opt out of waiting via setting this flag. You
	// would only need this option in the exceptional cases such as navigating to inaccessible
	// pages. Defaults to `false`.
	NoWaitAfter *bool `json:"noWaitAfter"`
	// A point to use relative to the top-left corner of element padding box. If not specified,
	// uses some visible point of the element.
	Position *PageClickOptionsPosition `json:"position"`
	// When true, the call requires selector to resolve to a single element. If given selector
	// resolves to more then one element, the call throws an exception.
	Strict *bool `json:"strict"`
	// Maximum time in milliseconds, defaults to 30 seconds, pass `0` to disable timeout.
	// The default value can be changed by using the BrowserContext.SetDefaultTimeout()
	// or Page.SetDefaultTimeout() methods.
	Timeout *float64 `json:"timeout"`
	// When set, this method only performs the [actionability](../actionability.md) checks
	// and skips the action. Defaults to `false`. Useful to wait until the element is ready
	// for the action without performing it.
	Trial *bool `json:"trial"`
}

func newLocatorWrapper(locator playwright.Locator) *locatorWrapper {
	return &locatorWrapper {
		Locator: locator,
	}
}

func (loc *locatorWrapper) Click(options ...playwright.PageClickOptions) {
	err := loc.Locator.Click(options...)
	if err != nil {
		log.Fatalf("could not get click work: %v", err)
	}
}

func (loc *locatorWrapper) Count() int {
	num, err := loc.Locator.Count()
	if err != nil {
		log.Fatalf("Locator count function error: %v", err)
	}
	return num
}

func (loc *locatorWrapper) Fill(value string, options ...playwright.FrameFillOptions) {
	err := loc.Locator.Fill(value, options...)
	if err != nil {
		log.Fatalf("fill function error: %v", err)
	}
}

func (loc *locatorWrapper) First() *locatorWrapper {
	firstLocator, err := loc.Locator.First()
	if err != nil {
		log.Fatalf("locator.first function error: %v", err)
	}

	return newLocatorWrapper(firstLocator)
}

func (loc *locatorWrapper) IsDisabled(options ...playwright.FrameIsDisabledOptions) bool {
	isDisabled, err := loc.Locator.IsDisabled(options...)
	if err != nil {
		log.Fatalf("isDisabled function error: %v", err)
	}
	return isDisabled
}

func (loc *locatorWrapper) IsVisible(options ...playwright.FrameIsVisibleOptions) bool {
	isVisible, err := loc.Locator.IsVisible(options...)
	if err != nil {
		log.Fatalf("isVisible function error: %v", err)
	}
	return isVisible
}

func (loc *locatorWrapper) Last() *locatorWrapper {
	lastLocator, err := loc.Locator.Last()
	if err != nil {
		log.Fatalf("locator.last function error: %v", err)
	}
	return newLocatorWrapper(lastLocator)
}

func (loc *locatorWrapper) Type(text string, options ...playwright.PageTypeOptions) {
	err := loc.Locator.Type(text, options...)
	if err != nil {
		log.Fatalf("type function error: %v", err)
	}
}