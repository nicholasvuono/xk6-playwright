package playwright

import (
	"strconv"
	"testing"

	"github.com/mxschmitt/playwright-go"
)

var tests = []func(t *testing.T){
	TestPlaywright,
	TestPlaywrightCustomBrowser,
}

func TestPlaywright(t *testing.T) {
	var pw Playwright
	var opts []string
	var opts2 playwright.PageGotoOptions
	var opts3 playwright.PageWaitForSelectorOptions

	pw.Launch(opts)
	pw.NewPage()
	pw.Goto("https://www.google.com", opts2)
	pw.WaitForSelector("//html/body/div[1]/div[2]", opts3)
	pw.Kill()
}

func TestPlaywrightCustomBrowser(t *testing.T) {
	var pw Playwright
	var opts []string
	var opts2 playwright.PageGotoOptions
	var opts3 playwright.PageWaitForSelectorOptions

	pw.LaunchCustomBrowser(opts, "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe", "")
	pw.NewPage()
	pw.Goto("https://www.google.com", opts2)
	pw.WaitForSelector("//html/body/div[1]/div[2]", opts3)
	pw.Kill()
}

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
