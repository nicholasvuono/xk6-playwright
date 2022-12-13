package playwright

import (
	"log"
	"os"
	"path/filepath"
	"io/ioutil"
	"encoding/base64"
	"strings"

	"github.com/playwright-community/playwright-go"
)

type fileChooserWrapper struct {
	FileChooser playwright.FileChooser
}

type InputFilePaths struct {
	Name	 string
	MimeType string
	Path	 string
}

func (f *fileChooserWrapper) SetFiles(files []InputFilePaths, options ...playwright.ElementHandleSetInputFilesOptions) {
	var inputFiles []playwright.InputFile
	log.Printf("I am in our custom setFiles go function")
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get cwd: %v", err)
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
			log.Fatalf("error with reading file: %v", err)
		}

		inputFiles = append(inputFiles, playwright.InputFile{
			Name:     inputFilePath.Name,
			MimeType: inputFilePath.MimeType,
			Buffer:   file,
		})
	}

	err = f.FileChooser.SetFiles(inputFiles, options...)
	if err != nil {
		log.Fatalf("error with setting files: %v", err)
	}
}

func newFileChooserWrapper(fileChooser playwright.FileChooser) *fileChooserWrapper {
	return &fileChooserWrapper {
		FileChooser: fileChooser,
	}
}