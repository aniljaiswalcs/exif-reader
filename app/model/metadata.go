package meta

import (
	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
)

type ExifMetaData struct {
	FilePath  string
	Latitude  float64
	Longitude float64
	Comment   string
}

// interface defines to fetch data .
type MetaData interface {
	ExtractExif(filepath string) (ExifMetaData, error)
}

// exifmeta implements MetaData interface.
type exifmeta struct{}

// NewCalculateExifData returns a new instance of MetaData.
func NewCalculateExifData() MetaData {
	return &exifmeta{}
}

func (e *exifmeta) ExtractExif(filepath string) (ExifMetaData, error) {

	var emd ExifMetaData
	rawExif, err := exif.SearchFileAndExtractExif(filepath)
	if err != nil {
		emd.FilePath = filepath
		emd.Comment = err.Error() + " found in Image"
		return emd, err
	}

	im, err := exifcommon.NewIfdMappingWithStandard()
	if err != nil {
		return emd, err
	}

	ti := exif.NewTagIndex()

	_, index, err := exif.Collect(im, ti, rawExif)
	if err != nil {
		if err.Error() == "unexpected EOF" {
			emd.FilePath = filepath
			emd.Comment = "File corrupt :: " + err.Error()
			return emd, err
		} else {
			emd.FilePath = filepath
			return emd, err
		}
	}

	ifd, err := index.RootIfd.ChildWithIfdPath(exifcommon.IfdGpsInfoStandardIfdIdentity)
	if err != nil {
		return emd, err
	}

	gi, err := ifd.GpsInfo()
	if err != nil {
		return emd, err
	}

	emd.FilePath = filepath
	emd.Latitude = gi.Latitude.Decimal()
	emd.Longitude = gi.Longitude.Decimal()

	return emd, nil
}
