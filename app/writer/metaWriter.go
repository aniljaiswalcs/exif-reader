package metawriter

type MetaDataWriter interface {
	WriteToFile(CsvDatawriter, HtmlDatawriter) error
}

type metadatawriter struct{}

func NewMetaDataWriter() MetaDataWriter {
	return &metadatawriter{}
}

func (m *metadatawriter) WriteToFile(csvwrite CsvDatawriter, htmlwrite HtmlDatawriter) error {
	err := csvwrite.WriteToCSV()
	if err != nil {
		return err
	}

	err = htmlwrite.writeToHtml()
	if err != nil {
		return err
	}

	return nil
}
