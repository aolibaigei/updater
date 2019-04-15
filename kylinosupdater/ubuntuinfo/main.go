package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Vulnerability struct {
	Name     string
	Severity string
	Package  string
	Version  string
}

func (v *Vulnerability) getInfo() (string, string, string, string) {
	return v.Name, v.Severity, v.Package, v.Version
}

var cvefile = "E:\\github\\ubuntu-cve-tracker\\active"

func main() {
	files, _ := ioutil.ReadDir(cvefile)
	for _, fi := range files {
		if fi.IsDir() == false && strings.Contains(fi.Name(), "CVE-201") {
			path := cvefile + "\\" + fi.Name()

			file, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			buf := bufio.NewReader(file)
			var v Vulnerability
			for {

				line, _, err := buf.ReadLine()
				if err == io.EOF {
					break
				}
				if strings.HasPrefix(string(line), "Candidate:") {
					v.Name = strings.TrimSpace(strings.TrimPrefix(string(line), "Candidate:"))
					// vulnerability.Link = fmt.Sprintf(cveURL, vulnerability.Name)
				}
				if strings.HasPrefix(string(line), "Priority:") {
					priority := strings.TrimSpace(strings.TrimPrefix(string(line), "Priority:"))

					// Handle syntax error: Priority: medium (heap-protector)
					if strings.Contains(priority, " ") {
						priority = priority[:strings.Index(priority, " ")]
					}
					v.Severity = priority
				}

				if strings.HasPrefix(string(line), "upstream_") {

					v.Package = strings.TrimPrefix(string(line)[:strings.Index(string(line), ":")], "upstream_")

					verline := strings.TrimSpace(strings.Replace(strings.TrimPrefix(string(line)[strings.Index(string(line), ":"):], ":"), "released", "", -1))

					// v.Version = verline[1 : len(verline)-1]
					v.Version = strings.Replace(strings.Replace(verline, "(", "", -1), ")", "", -1)

				}

			}
			cve, _, _, _ := v.getInfo()
			// if stats == "needs-triage" {
			// 	fmt.Println(v.getInfo())
			// }

			fmt.Println(cve)

		}

	}

}
