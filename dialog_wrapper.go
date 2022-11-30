package playwright

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

type dialogWrapper struct {
	Dialog playwright.Dialog
}

func newDialogWrapper(dialog playwright.Dialog) *dialogWrapper {
	return &dialogWrapper{
		Dialog: dialog,
	}
}

func (dw *dialogWrapper) Accept(texts ...string) {
	err := dw.Dialog.Accept(texts...)
	if err != nil {
		log.Fatalf("could not accept dialog: %v", err)
	}
}

func (dw *dialogWrapper) Dismiss() {
	err := dw.Dialog.Dismiss()
	if err != nil {
		log.Fatalf("could not dismiss dialog: %v", err)
	}
}
