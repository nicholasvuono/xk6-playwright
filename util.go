package playwright

import (
	"fmt"

	"github.com/dop251/goja"
)

func Throw(rt *goja.Runtime, message string, err error) {
	panic(rt.ToValue(fmt.Sprintf("%s: %v", message, err)))
}
