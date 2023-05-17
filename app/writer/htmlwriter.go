package metawriter

import (
	"os"
	"text/template"

	fileprocess "github.com/aniljaiswalcs/exif-reader/app/model"
)

type HtmlDatawriter struct {
	FilePath     string
	HTMLExifData []fileprocess.ExifMetaData
}

type ExifData struct {
	PageTitle string
	Exifdata  []fileprocess.ExifMetaData
}

func (h *HtmlDatawriter) writeToHtml() error {

	t, err := template.ParseFiles("./app/writer/exifdata.tmpl")
	if err != nil {
		return err
	}
	// Initialze a struct storing page data and data
	dataforhtml := ExifData{
		PageTitle: "Exif Data",
		Exifdata:  h.HTMLExifData,
	}

	var f *os.File
	f, err = os.Create(h.FilePath)
	if err != nil {
		return err
	}
	// Render the data and output using html file

	err = t.Execute(f, dataforhtml)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
