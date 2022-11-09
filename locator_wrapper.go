package playwright

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

type locatorWrapper struct {
	Locator playwright.Locator
}

func newLocatorWrapper(locator playwright.Locator) *locatorWrapper {
	return &locatorWrapper{
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
