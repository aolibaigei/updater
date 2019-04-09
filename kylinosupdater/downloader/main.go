package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type PassThru struct {
	io.Reader
	curr  int64
	total float64
}

func (pt *PassThru) Read(p []byte) (int, error) {
	n, err := pt.Reader.Read(p)
	pt.curr += int64(n)

	// last read will have EOF err
	if err == nil || (err == io.EOF && n > 0) {
		printProgress(float64(pt.curr), pt.total)
	}

	return n, err
}

func printProgress(curr, total float64) {
	width := 40.0
	output := ""
	threshold := (curr / total) * float64(width)
	for i := 0.0; i < width; i++ {
		if i < threshold {
			output += "="
		} else {
			output += " "
		}
	}

	fmt.Printf("\r[%s] %.1f of %.1fMB", output, curr/bytesToMegaBytes, total/bytesToMegaBytes)
}

var bytesToMegaBytes = 1048576.0
var url = "http://archive.kylinos.cn/yhkylin/dists/juniper/universe/binary-amd64/Packages"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	out, err := os.Create("package.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	src := &PassThru{Reader: resp.Body, total: float64(resp.ContentLength)}
	size, err := io.Copy(out, src)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nFile Transferred. (%.1f MB)\n", float64(size)/bytesToMegaBytes)
}
