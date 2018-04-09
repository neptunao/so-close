package data

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/neptunao/so-close/errors"
)

type csvDataIterator struct {
	file   *os.File
	reader *csv.Reader
	errors *errors.AggregateError
}

func (itr csvDataIterator) Next() (interface{}, bool) {
	rec, err := itr.reader.Read()
	if err == io.EOF {
		return nil, false
	}
	if err != nil {
		itr.errors.Add(err)
		return nil, true
	}
	return rec, true
}

func (itr csvDataIterator) Err() error {
	return itr.errors
}

func (itr csvDataIterator) Close() error {
	return itr.file.Close()
}

type csvDataConnector struct {
	filename string
}

func (dc csvDataConnector) Connect() (Iterator, error) {
	f, err := os.Open(dc.filename)
	if err != nil {
		return nil, err
	}
	return csvDataIterator{
		file:   f,
		reader: csv.NewReader(f),
	}, nil
}

// ConnectCSVFile opens CSV file and return Iterator for it
func ConnectCSVFile(filename string) (Iterator, error) {
	c := csvDataConnector{filename}
	return c.Connect()
}
