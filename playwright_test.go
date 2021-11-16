package playwright

import (
	"strconv"
	"testing"

	"github.com/mxschmitt/playwright-go"
)

var tests = []func(t *testing.T){
	TestPlaywright,
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
}

func TestAll(t *testing.T) {
	for i, test := range tests {
		t.Run(strconv.Itoa(i), test)
	}
}
