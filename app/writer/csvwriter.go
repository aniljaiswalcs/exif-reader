package metawriter

import (
	"encoding/csv"
	"fmt"
	"os"

	fileprocess "github.com/aniljaiswalcs/exif-reader/app/model"
)

type CsvDatawriter struct {
	FilePath    string
	CsvExifData []fileprocess.ExifMetaData
}

func (d *CsvDatawriter) WriteToCSV() error {

	f, err := os.Create(d.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	header := []string{"FilePath", "Latitude", "Longitude", "Comment"}
	err = writer.Write(header)
	if err != nil {
		return err
	}
	var row []string
	for _, data := range d.CsvExifData {
		row = []string{data.FilePath, fmt.Sprintf("%f", data.Latitude), fmt.Sprintf("%f", data.Longitude), data.Comment}
		err = writer.Write(row)
		if err != nil {
			return err
		}
	}
	return nil
}
