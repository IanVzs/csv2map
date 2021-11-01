package csv2map

import (
	"encoding/csv"
	"io"
)

type Reader struct {
	csv.Reader
	hadColName  bool
	ColumnNames []string
}

func (r *Reader) readColumnNames() error {
	var err error
	if !r.hadColName {
		r.hadColName = true
		r.ColumnNames, err = r.Read()
		return err
	}
	return nil
}
func NewReader(r io.Reader) *Reader {
	reader := &Reader{*csv.NewReader(r), false, []string{}}
	reader.readColumnNames()
	return reader
}

func (r *Reader) Read2Map() (map[string]string, error) {
	record := make(map[string]string)
	if !r.hadColName {
		r.readColumnNames()
	}
	record_str, err := r.Read()
	if err != nil {
		return nil, err
	}
	if len(r.ColumnNames) == len(record_str) {
		for i, v := range r.ColumnNames {
			record[v] = record_str[i]
		}
	}
	return record, err
}
