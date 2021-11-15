package playwright

import "go.k6.io/k6/js/modules"

func init() {
	modules.Register("k6/x/playwright", new(Playwright))
}

type Playwright struct {}
