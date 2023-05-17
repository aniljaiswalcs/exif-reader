package metawriter_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	fmodel "github.com/aniljaiswalcs/exif-reader/app/model"
	fwrite "github.com/aniljaiswalcs/exif-reader/app/writer"
)

func TestWriteToCSV(t *testing.T) {
	tests := []struct {
		name          string
		csvDatawriter fwrite.CsvDatawriter
		expectedError bool
	}{
		{
			name: "Valid data",
			csvDatawriter: fwrite.CsvDatawriter{
				FilePath: "output.csv",
				CsvExifData: []fmodel.ExifMetaData{
					{
						FilePath:  "image1.jpg",
						Latitude:  40.7128,
						Longitude: -74.0060,
						Comment:   "Beautiful scenery",
					},
					{
						FilePath:  "image2.jpg",
						Latitude:  37.7749,
						Longitude: -122.4194,
						Comment:   "City skyline",
					},
				},
			},
			expectedError: false,
		},
		{
			name: "Empty data",
			csvDatawriter: fwrite.CsvDatawriter{
				FilePath:    "output.csv",
				CsvExifData: []fmodel.ExifMetaData{},
			},
			expectedError: false,
		},
		{
			name: "Invalid file path",
			csvDatawriter: fwrite.CsvDatawriter{
				FilePath:    "/invalid/path/output.csv",
				CsvExifData: []fmodel.ExifMetaData{},
			},
			expectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.csvDatawriter.WriteToCSV()
			if test.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
		os.Remove("output.csv")
	}
}
