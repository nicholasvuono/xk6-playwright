package playwright

import (
	"github.com/playwright-community/playwright-go"
	"go.k6.io/k6/js/modules"
)

type dialogWrapper struct {
	Dialog playwright.Dialog
	vu     modules.VU
}

func newDialogWrapper(dialog playwright.Dialog, vu modules.VU) *dialogWrapper {
	return &dialogWrapper{
		Dialog: dialog,
		vu:     vu,
	}
}

func (dw *dialogWrapper) Accept(texts ...string) {
	err := dw.Dialog.Accept(texts...)
	if err != nil {
		Throw(dw.vu.Runtime(), "Error accepting dialog", err)
	}
}

func (dw *dialogWrapper) Dismiss() {
	err := dw.Dialog.Dismiss()
	if err != nil {
		Throw(dw.vu.Runtime(), "Error dismissing dialog", err)
	}
}
