package converter

import (
	"encoding/csv"
	"log"
	"runtime"
	"strings"
)

//StringSlice2Cvs string slice is converted to a cvs string. A wrapper.
func StringSlice2Cvs(origin []string) (string, error) {
	var stringBuil strings.Builder
	var err error
	w := csv.NewWriter(&stringBuil)
	if runtime.GOOS == "windows" {
		w.UseCRLF = true
	}
	if err = w.Write(origin); err != nil {
		return "", err
	}
	w.Flush()
	return stringBuil.String(), err
}

//Cvs2StringSlice Convert the cvs string to string slices,defalut to use 2 dimension
func Cvs2StringSlice(in string) ([][]string, error) {
	r := csv.NewReader(strings.NewReader(in))
	var err error
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records, err
}
