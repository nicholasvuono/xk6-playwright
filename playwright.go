package playwright

import (
	"github.com/playwright-community/playwright-go"
	"go.k6.io/k6/js/modules"
)

// Register the extension on module initialization, available to
// import from JS as "k6/x/playwright".
func init() {
	modules.Register("k6/x/playwright", New())
}

type (
	// RootModule is the global module instance that will create module
	// instances for each VU.
	RootModule struct{}

	// ModuleInstance represents an instance of the JS module.
	ModuleInstance struct {
		// vu provides methods for accessing internal k6 objects for a VU
		vu modules.VU
		// playwright is the exported type
		playwright *Playwright
	}
)

// Ensure the interfaces are implemented correctly.
var (
	_ modules.Instance = &ModuleInstance{}
	_ modules.Module   = &RootModule{}
)

// New returns a pointer to a new RootModule instance.
func New() *RootModule {
	return &RootModule{}
}

// NewModuleInstance implements the modules.Module interface returning a new instance for each VU.
func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	var p = new(Playwright)
	pw, err := playwright.Run()

	if err != nil {
		Throw(p.vu.Runtime(), "Unable to start playwright", err)
	}

	p.Self = pw

	return &ModuleInstance{
		vu:         vu,
		playwright: &Playwright{Self: pw, vu: vu},
	}
}

// Playwright is the k6 extension for a playwright-go client.
type Playwright struct {
	Self *playwright.Playwright
	vu   modules.VU // provides methods for accessing internal k6 objects
}

// Launch starts the playwright client and launches a browser
func (p *Playwright) Launch(args playwright.BrowserTypeLaunchOptions) *browserWrapper {
	browser, err := p.Self.Chromium.Launch(args)
	if err != nil {
		Throw(p.vu.Runtime(), "Error launching browser", err)
	}
	return newBrowserWrapper(browser, p.vu)
}

// Kill closes browser instance and stops puppeteer client
func (p *Playwright) Kill() {
	if err := p.Self.Stop(); err != nil {
		Throw(p.vu.Runtime(), "Error stopping Playwright", err)
	}
}

// Exports implements the modules.Instance interface and returns the exported types for the JS module.
func (mi *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{
		Default: mi.playwright,
	}
}
