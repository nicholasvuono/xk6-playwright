package playwright

import (
	"log"
	"os"
	"path/filepath"
	"io/ioutil"

	"github.com/playwright-community/playwright-go"
)

type fileChooserWrapper struct {
	FileChooser playwright.FileChooser
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
		log.Fatalf("could not get cwd: %v", err)
	}

	for _, inputFilePath := range files {
		file, err := ioutil.ReadFile(filepath.Join(cwd, inputFilePath.Path))

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