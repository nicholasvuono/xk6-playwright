package playwright

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/playwright-community/playwright-go"
	"go.k6.io/k6/js/modules"
)

type fileChooserWrapper struct {
	FileChooser playwright.FileChooser
	vu          modules.VU
}

type InputFilePaths struct {
	Name     string
	MimeType string
	Path     string
}

func (f *fileChooserWrapper) SetFiles(files []InputFilePaths, options ...playwright.ElementHandleSetInputFilesOptions) {
	var inputFiles []playwright.InputFile
	log.Printf("I am in our custom setFiles go function")
	cwd, err := os.Getwd()
	if err != nil {
		Throw(f.vu.Runtime(), "Error getting current working directory (cwd)", err)
	}

	for _, inputFilePath := range files {

		var file []byte
		var err error

		if strings.HasPrefix(inputFilePath.Path, "data:") {
			base64File := strings.SplitN(inputFilePath.Path, "base64,", 2)

			file, err = base64.StdEncoding.DecodeString(base64File[1])

		} else {
			file, err = ioutil.ReadFile(filepath.Join(cwd, inputFilePath.Path))
		}

		if err != nil {
			Throw(f.vu.Runtime(), "Error reading file", err)
		}

		inputFiles = append(inputFiles, playwright.InputFile{
			Name:     inputFilePath.Name,
			MimeType: inputFilePath.MimeType,
			Buffer:   file,
		})
	}

	err = f.FileChooser.SetFiles(inputFiles, options...)
	if err != nil {
		Throw(f.vu.Runtime(), "Error setting files", err)
	}
}

func newFileChooserWrapper(fileChooser playwright.FileChooser, vu modules.VU) *fileChooserWrapper {
	return &fileChooserWrapper{
		FileChooser: fileChooser,
		vu:          vu,
	}
}
