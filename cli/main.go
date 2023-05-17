package main

import (
	"fmt"
	"log"
	"os"

	fileprocess "github.com/aniljaiswalcs/exif-reader/app/model"
	filepathreader "github.com/aniljaiswalcs/exif-reader/app/reader"
	filewriter "github.com/aniljaiswalcs/exif-reader/app/writer"
)

const imageDirectoryPath = "./imagefile"
const csvfilepath = "./output/exifdata.csv"
const htmlfilepath = "./output/exifdata.html"

func main() {
	fmt.Println("Start Processing images for Exif Data")
	//Read the directory
	filepathAray := filepathreader.NewReader()
	arrayOfFilePath, err := filepathAray.FilePath(imageDirectoryPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	//Calculate the exif data
	calculateExif := fileprocess.NewCalculateExifData()
	var exifdata []fileprocess.ExifMetaData
	for _, elem := range arrayOfFilePath {
		exifMetaData, err := calculateExif.ExtractExif(elem)
		if err != nil {
			exifdata = append(exifdata, exifMetaData)
			continue
		}
		exifdata = append(exifdata, exifMetaData)
	}

	// write the data to csv and html file
	write := filewriter.NewMetaDataWriter()

	csvwrite := filewriter.CsvDatawriter{FilePath: csvfilepath, CsvExifData: exifdata}
	htmlwrite := filewriter.HtmlDatawriter{FilePath: htmlfilepath, HTMLExifData: exifdata}
	err = write.WriteToFile(csvwrite, htmlwrite)
	if err != nil {
		log.Fatal("Error to write exif data in file  ", err)
	}
	fmt.Println("Data Generated, check the exifdata.csv and exifdata.html in output folder")
}
