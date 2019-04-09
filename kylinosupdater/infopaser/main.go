package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type VersionInfo struct {
	Package  string
	Version  string
	Homepage string
	Bugs     string
	Origin   string
}

func parsinfo(info string) string {
	splitno := strings.Index(info, ":")

	v := info[splitno:]

	return strings.TrimSpace(strings.TrimPrefix(v, ":"))
}

func main() {

	fi, err := os.Open("../package.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	//

	f, err := os.Create("info.txt")
	if err != nil {
		fmt.Printf("create map file error: %v\n", err)
	}
	defer f.Close()

	///

	bug := &VersionInfo{}
	br := bufio.NewReader(fi)
	for {

		a, _, c := br.ReadLine()

		if len(a) != 0 {

			if strings.Contains(string(a), "Package:") {
				bug.Package = parsinfo(string(a))
			}
			if strings.Contains(string(a), "Version:") {

				if strings.HasSuffix(parsinfo(string(a)), "kord1") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord1")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord2") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord2")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord3") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord3")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord4") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord4")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord5") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord5")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord6") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord6")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord7") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord7")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord8") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord8")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord9") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord9")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord10") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord10")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord11") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord11")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord12") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord13")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord14") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord14")
				} else if strings.HasSuffix(parsinfo(string(a)), "kord") {
					bug.Version = strings.TrimSuffix(parsinfo(string(a)), "kord")
				} else {
					bug.Version = parsinfo(string(a))
				}

			}

			if strings.Contains(string(a), "Homepage:") {
				bug.Homepage = parsinfo(string(a))
			}
			if strings.Contains(string(a), "Bugs:") {
				bug.Bugs = parsinfo(string(a))
			}
			if strings.Contains(string(a), "Origin:") {
				bug.Origin = parsinfo(string(a))
			}

		}
		if len(a) == 0 {
			info := fmt.Sprintf("%s %s %s %s %s \n", bug.Package, bug.Version, bug.Bugs, bug.Homepage, bug.Origin)
			_, err1 := io.WriteString(f, info)
			if err1 != nil {
				panic(err1)
			}
		}

		if c == io.EOF {
			break
		}
	}

}
