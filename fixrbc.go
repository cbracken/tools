// CLI tool to rename all RBC statement PDFs in the current directory to
// YYYY-MM-DD.pdf format.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".pdf") {
			// Format: ACCTNO-YYYYmmmDD-YYYYmmmDD.pdf
			date := strings.Split(f.Name(), "-")[2][0:9]
			y := date[0:4]
			m, err := getMonthNumber(date[4:7])
			if err != nil {
				log.Fatal(err)
			}
			d := date[7:9]

			newFname := fmt.Sprintf("%s-%s-%s.pdf", y, m, d)
			os.Rename(f.Name(), newFname)
		}
	}
}

func getMonthNumber(m string) (string, error) {
	switch {
	case m == "Jan":
		return "01", nil
	case m == "Feb":
		return "02", nil
	case m == "Mar":
		return "03", nil
	case m == "Apr":
		return "04", nil
	case m == "May":
		return "05", nil
	case m == "Jun":
		return "06", nil
	case m == "Jul":
		return "07", nil
	case m == "Aug":
		return "08", nil
	case m == "Sep":
		return "09", nil
	case m == "Oct":
		return "10", nil
	case m == "Nov":
		return "11", nil
	case m == "Dec":
		return "12", nil
	}
	return "", fmt.Errorf("Bad month name: %s", m)
}
