# Task

Directory that contains images, and that also contains a sub-directory that contains images, create a command-line utility written in Go that reads the EXIF data from the images and writes the following attributes to a CSV file:

 

image file path
GPS position latitude
GPS position longitude
 

Feel free to use go-exif, or other exif-extraction library, to read the EXIF data.

## Usage
### Run application from console:
#### To run the source code:
```
go run cli/main.go

```
after running the above command, exifdata.csv and exifdata.html files will get generated in output directory.

#### Test
```
go test ./...

```

Above command will execute the test files.

## Notes

Extra "comment" field added in the output for those files, which are not processed due to error or no data found.
