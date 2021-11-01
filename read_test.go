package csv2map

import (
	"io"
	"os"
	"testing"
)

func TestReadCSV(t *testing.T) {
	// t.Fatal("not implemented")
	f, err := os.Open("test.csv")
	if err != nil {
		t.Logf("Path(%s) Error: %s", "test.csv", err.Error())
	}
	reader := NewReader(f)
	for i := 0; true; i++ {
		line, err := reader.Read2Map()
		if err == io.EOF {
			break
		}
		t.Log(line)
	}
}
